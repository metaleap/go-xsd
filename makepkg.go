package xsd

import (
	"fmt"
	"strings"
	"time"

	util "github.com/metaleap/go-util"
	ustr "github.com/metaleap/go-util/str"
)

var (
	PkgGen = &goPkgSrcMaker {
		BaseTypes: map[string]string {
			"decimal": "complex128",
			"float": "float32",
			"double": "float64",
			"duration": "time.Duration",
			"dateTime": "time.Time",
			"date": "time.Time",
			"hexBinary": "[]byte",
			"base64Binary": "[]byte",
			"integer": "int",
			"nonPositiveInteger": "int",
			"negativeInteger": "int",
			"long": "int64",
			"int": "int32",
			"short": "int16",
			"byte": "byte",
			"nonNegativeInteger": "uint",
			"unsignedLong": "uint64",
			"unsignedInt": "uint32",
			"unsignedShort": "uint16",
			"unsignedByte": "uint8",
			"positiveInteger": "uint",
		},
		Name: "goxsdpkg",
		BaseCodePath: util.BaseCodePath("metaleap", "go-xsd-pkg"),
		BasePath: "github.com/metaleap/go-xsd-pkg",
	}
)

type goPkgSrcMaker struct {
	BaseTypes map[string]string
	BaseCodePath, BasePath, Name string

	impName string
	lines []string
	imports map[string]string
	impsUsed map[string]bool
}

	func (me *goPkgSrcMaker) append (lines ... string) {
		me.lines = append(me.lines, lines ...)
	}

	func (me *goPkgSrcMaker) appendFmt (addLineAfter bool, format string, fmtArgs ... interface{}) {
		me.append(fmt.Sprintf(format, fmtArgs ...))
		if addLineAfter { me.append("") }
	}

	func (me *goPkgSrcMaker) insertFmt (index int, format string, fmtArgs ... interface{}) {
		me.lines = append(me.lines[: index], append([]string { fmt.Sprintf(format, fmtArgs ...) }, me.lines[index : ] ...) ...)
	}

	func (me *goPkgSrcMaker) safeName (name string) string {
		return ustr.SafeIdentifier(name, "")
	}

	func (me *goPkgSrcMaker) reinit () {
		me.impName = "xsdt"
		me.imports, me.impsUsed, me.lines = map[string]string {}, map[string]bool {}, []string { "package " + me.Name, "" }
	}

type makerBag struct {
	Schema *Schema

	curSimpleType *SimpleType
	now int64
	snow string
}

	func (me *makerBag) resolveTypeRef (typeRef string) string {
		var ns = me.Schema.XMLNamespaces[""]
		var impName = ""
		if len(typeRef) == 0 { return "" }
		if pos := strings.Index(typeRef, ":"); pos > 0 {
			impName, ns = typeRef[: pos], me.Schema.XMLNamespaces[typeRef[: pos]]
			typeRef = typeRef[pos + 1 :]
		}
		if ns == xsdNamespaceUri { impName = PkgGen.impName }
		if ns == me.Schema.TargetNamespace.String() { impName = "" }
		PkgGen.impsUsed[impName] = true
		return ustr.PrefixWithSep(impName, ".", PkgGen.safeName(typeRef))
	}

func (me *Schema) makePkg (bag *makerBag) {
	var impPos int
	PkgGen.reinit(); bag.now = time.Now().UnixNano(); bag.snow = fmt.Sprintf("%v", bag.now)
	for k, _ := range me.XMLNamespaces { if k == PkgGen.impName { PkgGen.impName += bag.snow } }
	PkgGen.imports[PkgGen.impName] = "github.com/metaleap/go-xsd/types"
	for _, im := range []string { "strings" } { PkgGen.imports[im] = "" }
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsNotation.makePkg(bag)
	me.hasElemsImport.makePkg(bag)
	impPos = len(PkgGen.lines) + 1
	PkgGen.append("import (", ")", "")
	me.hasElemsSimpleType.makePkg(bag)
	for impName, impPath := range PkgGen.imports {
		if PkgGen.impsUsed[impName] {
			if len(impPath) > 0 {
				PkgGen.insertFmt(impPos, "\t%v \"%v\"", impName, impPath)
			} else {
				PkgGen.insertFmt(impPos, "\t\"%v\"", impName)
			}
		}
	}
}
