package xsd

import (
	"path"
	"strings"
)

func (me *Annotation) makePkg (bag *makerBag) {
	me.hasElemsDocumentation.makePkg(bag)
}

func (me *Documentation) makePkg (bag *makerBag) {
	if len(me.CDATA) > 0 {
		var ln, s string
		for _, ln = range strings.Split(me.CDATA, "\n") {
			if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 { PkgGen.appendFmt("//\t%s", s) }
		}
	}
}

func (me *Import) makePkg (bag *makerBag) {
	var impName, impPath string
	var pos int
	me.hasElemAnnotation.makePkg(bag)
	for k, v := range bag.Schema.XMLNamespaces { if v == me.Namespace { impName = k; break } }
	if len(impName) > 0 {
		if pos, impPath = strings.Index(me.SchemaLocation, protSep), me.SchemaLocation; pos > 0 {
			impPath = impPath[pos + len(protSep) :]
		} else {
			impPath = path.Join(path.Dir(bag.Schema.loadUri), impPath)
		}
		PkgGen.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
}

func (me *SimpleType) makePkg (bag *makerBag) {
	me.hasElemAnnotation.makePkg(bag)
	if me.RestrictionSimpleType != nil {
		PkgGen.appendFmt("type %s%s %s", PkgGen.TypePrefix, PkgGen.pascalCase(me.Name), me.RestrictionSimpleType.Base)
	}
}
