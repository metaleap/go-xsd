package xsd

import (
	"fmt"
	"path"
	"strings"
	"time"

	ustr "github.com/metaleap/go-util/str"

	xsdt "github.com/metaleap/go-xsd/types"
)

func (me *All) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Annotation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAppInfo.makePkg(bag)
	me.hasElemsDocumentation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Any) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AnyAttribute) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AppInfo) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Attribute) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AttributeGroup) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Choice) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemExtensionComplexContent.makePkg(bag)
	me.hasElemRestrictionComplexContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexType) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemComplexContent.makePkg(bag)
	me.hasElemSimpleContent.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Documentation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	if len(me.CDATA) > 0 { var s, ln string;
		for _, ln = range ustr.Split(me.CDATA, "\n") { if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 { bag.appendFmt(false, "//\t%s", s) } }
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Element) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemUnique.makePkg(bag)
	me.hasElemsKey.makePkg(bag)
	me.hasElemComplexType.makePkg(bag)
	me.hasElemKeyRef.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionSimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Field) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Group) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Import) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
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
		bag.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Key) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *KeyRef) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *List) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	var rtr = bag.resolveTypeRef(me.ItemType.String())
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	if len(rtr) == 0 {
		rtr = me.SimpleTypes[0].Name.String()
		if len(rtr) == 0 {
			rtr = bag.safeName(bag.Stacks.CurSimpleType().Name.String() + "SubList")
			me.SimpleTypes[0].Name = xsdt.NCName(rtr)
		}
	}
	bag.impsUsed["xsdt"] = true
	bag.appendFmt(false, `	func (me %v) Values () (list []%v) {
		var btv = new(%v)
		var svals = xsdt.ListValues(string(me))
		list = make([]%v, len(svals))
		for i, s := range svals { btv.SetFromString(s); list[i] = *btv }
		return
	}
	`, bag.safeName(bag.Stacks.CurSimpleType().Name.String()), rtr, rtr, rtr)
	me.elemBase.afterMakePkg(bag)
}

func (me *Notation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(false, "Notations.Add(%#v, %#v, %#v, %#v)", me.Id, me.Name, me.Public, me.System)
	me.elemBase.afterMakePkg(bag)
}

func (me *Redefine) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemLength.makePkg(bag)
	me.hasElemPattern.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsEnumeration.makePkg(bag)
	me.hasElemFractionDigits.makePkg(bag)
	me.hasElemMaxExclusive.makePkg(bag)
	me.hasElemMaxInclusive.makePkg(bag)
	me.hasElemMaxLength.makePkg(bag)
	me.hasElemMinExclusive.makePkg(bag)
	me.hasElemMinInclusive.makePkg(bag)
	me.hasElemMinLength.makePkg(bag)
	me.hasElemTotalDigits.makePkg(bag)
	me.hasElemWhiteSpace.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleEnumeration) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	bag.appendFmt(true, "\tfunc (me %v) Is%v () bool { return me == %#v }", bag.safeName(bag.Stacks.CurSimpleType().Name.String()), bag.safeName(me.Value), me.Value)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleFractionDigits) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxExclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxInclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinExclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinInclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimplePattern) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleTotalDigits) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleType) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemLength.makePkg(bag)
	me.hasElemPattern.makePkg(bag)
	me.hasElemsEnumeration.makePkg(bag)
	me.hasElemFractionDigits.makePkg(bag)
	me.hasElemMaxExclusive.makePkg(bag)
	me.hasElemMaxInclusive.makePkg(bag)
	me.hasElemMaxLength.makePkg(bag)
	me.hasElemMinExclusive.makePkg(bag)
	me.hasElemMinInclusive.makePkg(bag)
	me.hasElemMinLength.makePkg(bag)
	me.hasElemTotalDigits.makePkg(bag)
	me.hasElemWhiteSpace.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleWhiteSpace) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Schema) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	var impPos int
	bag.reinit(); bag.now = time.Now().UnixNano(); bag.snow = fmt.Sprintf("%v", bag.now)
	for k, _ := range me.XMLNamespaces { if k == bag.impName { bag.impName += bag.snow } }
	bag.imports[bag.impName] = "github.com/metaleap/go-xsd/types"
	// for _, im := range []string { "strings" } { bag.imports[im] = "" }
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsImport.makePkg(bag)
	impPos = len(bag.lines) + 1
	bag.append("import (", ")", "")
	if me.XSDParentSchema == nil { anonCount = 0; bag.append("type hasCdata struct { CDATA string `xml:\",chardata\"` }", "") }
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsNotation.makePkg(bag)
	me.hasElemsRedefine.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	for impName, impPath := range bag.imports {
		if bag.impsUsed[impName] {
			if len(impPath) > 0 {
				bag.insertFmt(impPos, "\t%v \"%v\"", impName, impPath)
			} else {
				bag.insertFmt(impPos, "\t\"%v\"", impName)
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Selector) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Sequence) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemExtensionSimpleContent.makePkg(bag)
	me.hasElemRestrictionSimpleContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleType) makePkg (bag *PkgBag) {
	var typeName = me.Name
	var baseType, safeName = "", ""
	if len(typeName) == 0 { typeName = xsdt.NCName(bag.AnonName()); me.hasAttrName.Name = typeName }
	me.elemBase.beforeMakePkg(bag)
	bag.Stacks.SimpleType.Push(me)
	safeName = bag.safeName(typeName.String())
	me.hasElemAnnotation.makePkg(bag)
	if me.RestrictionSimpleType != nil { baseType = me.RestrictionSimpleType.Base.String() }
	if len(baseType) == 0 { baseType = ustr.PrefixWithSep(bag.Schema.XSDNamespace, ":", "string") }
	baseType = bag.resolveTypeRef(baseType)
	bag.appendFmt(false, `type %s %s

	func (me *%s) SetFromString (s string) { (*%v)(me).SetFromString(s) }

	func (me %s) String () string { return %v(me).String() }
	`, safeName, baseType, safeName, baseType, safeName, baseType)
	me.hasElemRestrictionSimpleType.makePkg(bag)
	me.hasElemList.makePkg(bag)
	me.hasElemUnion.makePkg(bag)
	bag.Stacks.SimpleType.Pop()
	me.elemBase.afterMakePkg(bag)
}

func (me *Union) makePkg (bag *PkgBag) {
	var memberTypes []string
	var rtn string
	var isParseType = false
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(me.MemberTypes, " ")
	for _, st := range me.SimpleTypes { memberTypes = append(memberTypes, st.Name.String()) }
	for _, mt := range memberTypes {
		isParseType, rtn = false, bag.resolveTypeRef(mt)
		for _, pt := range xsdt.ParseTypes {
			if isParseType = (strings.HasSuffix(rtn, pt) || strings.HasSuffix(strings.ToLower(rtn), strings.ToLower(pt + "type"))); isParseType { break }
		}
		if isParseType {
			bag.appendFmt(true, "func (me %v) To%v () %v { var x = new(%v); x.SetFromString(me.String()); return *x }", bag.safeName(bag.Stacks.CurSimpleType().Name.String()), bag.safeName(rtn), rtn, rtn)
		} else {
			bag.appendFmt(true, "func (me %v) To%v () %v { return %v(me) }", bag.safeName(bag.Stacks.CurSimpleType().Name.String()), bag.safeName(rtn), rtn, rtn)
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Unique) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}
