package xsd

import (
	"path"
	"strings"

	ustr "github.com/metaleap/go-util/str"

	xsdt "github.com/metaleap/go-xsd/types"
)

func (me *Annotation) makePkg (bag *makerBag) {
	me.hasElemsDocumentation.makePkg(bag)
}

func (me *Documentation) makePkg (bag *makerBag) {
	if len(me.CDATA) > 0 {
		var ln, s string
		for _, ln = range strings.Split(me.CDATA, "\n") {
			if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 { PkgGen.appendFmt(false, "//\t%s", s) }
		}
	}
}

func (me *RestrictionSimpleEnumeration) makePkg (bag *makerBag) {
	PkgGen.appendFmt(true, "\tfunc (me %v) Is%v () bool { return me == %#v }", PkgGen.safeName(bag.curSimpleType.Name.String()), PkgGen.safeName(me.Value), me.Value)
}

func (me *Import) makePkg (bag *makerBag) {
	var impName, impPath string
	var pos int
	me.hasElemAnnotation.makePkg(bag)
	for k, v := range bag.Schema.XMLNamespaces { if v == me.Namespace { impName = k; break } }
	if len(impName) > 0 {
		if pos, impPath = strings.Index(me.SchemaLocation.String(), protSep), me.SchemaLocation.String(); pos > 0 {
			impPath = impPath[pos + len(protSep) :]
		} else {
			impPath = path.Join(path.Dir(bag.Schema.loadUri), impPath)
		}
		impPath = path.Join(path.Dir(impPath), goPkgPrefix + path.Base(impPath) + goPkgSuffix)
		PkgGen.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
}

func (me *List) makePkg (bag *makerBag) {
	var rtr = bag.resolveTypeRef(me.ItemType.String())
	if len(rtr) == 0 {
		var cst = bag.curSimpleType
		rtr = me.SimpleTypes[0].Name.String()
		if len(rtr) == 0 { rtr = PkgGen.safeName(cst.Name.String() + "SubList"); me.SimpleTypes[0].Name = xsdt.NCName(rtr) }
		me.SimpleTypes[0].makePkg(bag)
		bag.curSimpleType = cst
	}
	me.hasElemAnnotation.makePkg(bag)
	PkgGen.impsUsed["strings"] = true
	PkgGen.appendFmt(false, `	func (me %v) Values () (list []%v) {
		var btv = new(%v)
		var spl = strings.Split(string(me), " ")
		list = make([]%v, len(spl))
		for i, s := range spl { btv.SetFromString(s); list[i] = *btv }
		return
	}
	`, PkgGen.safeName(bag.curSimpleType.Name.String()), rtr, rtr, rtr)
//	me.hasElemsSimpleType.makePkg(bag)
}

func (me *Notation) makePkg (bag *makerBag) {
	me.hasElemAnnotation.makePkg(bag)
	PkgGen.appendFmt(false, "Notations.Add(%#v, %#v, %#v, %#v)", me.Id, me.Name, me.Public, me.System)
}

func (me *SimpleType) makePkg (bag *makerBag) {
	var baseType, safeName = "", PkgGen.safeName(me.Name.String())
	bag.curSimpleType = me
	me.hasElemAnnotation.makePkg(bag)
	if me.RestrictionSimpleType != nil { baseType = me.RestrictionSimpleType.Base.String() }
	if len(baseType) == 0 { baseType = ustr.PrefixWithSep(bag.Schema.XSDNamespace, ":", "string") }
	baseType = bag.resolveTypeRef(baseType)
	PkgGen.appendFmt(false, `type %s %s

	func (me *%s) SetFromString (s string) { (*%v)(me).SetFromString(s) }

	func (me %s) String () string { return %v(me).String() }
	`, safeName, baseType, safeName, baseType, safeName, baseType)
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.hasElemAnnotation.makePkg(bag)
		me.RestrictionSimpleType.hasElemsEnumeration.makePkg(bag)
	}
	me.hasElemList.makePkg(bag)
}
