package xsd

import (
	"fmt"
	"strings"

	util "github.com/metaleap/go-util"
	ustr "github.com/metaleap/go-util/str"

	xsdt "github.com/metaleap/go-xsd/types"
)

var (
	anonCount uint64
	PkgGen = &pkgGen {
		BaseCodePath: util.BaseCodePath("metaleap", "go-xsd-pkg"),
		BasePath: "github.com/metaleap/go-xsd-pkg",
	}
)

type pkgGen struct {
	BaseCodePath, BasePath string
}

type beforeAfterMake interface {
	afterMakePkg (*PkgBag)
	beforeMakePkg (*PkgBag)
}

type pkgStack []interface{}

	func (me *pkgStack) Pop () (el interface{}) { sl := *me; el = sl[0]; *me = sl[1 :]; return }

	func (me *pkgStack) Push (el interface{}) { nu := []interface{} { el }; *me = append(nu, *me...) }

type pkgStacks struct {
	Name, SimpleType pkgStack
}

	func (me *pkgStacks) CurName () (r xsdt.NCName) {
		if len(me.Name) > 0 { r = me.Name[0].(xsdt.NCName) }; return
	}

	func (me *pkgStacks) CurSimpleType () (r *SimpleType) {
		if len(me.SimpleType) > 0 { r = me.SimpleType[0].(*SimpleType) }; return
	}

type PkgBag struct {
	Schema *Schema
	Stacks pkgStacks

	lines []string
	impName string
	imports map[string]string
	impsUsed map[string]bool
	now int64
	snow string
}

	func (me *PkgBag) AnonName () string {
		anonCount++; return fmt.Sprintf("XsdAnon%v", anonCount)
	}

	func (me *PkgBag) append (lines ... string) {
		me.lines = append(me.lines, lines ...)
	}

	func (me *PkgBag) appendFmt (addLineAfter bool, format string, fmtArgs ... interface{}) {
		me.append(fmt.Sprintf(format, fmtArgs ...))
		if addLineAfter { me.append("") }
	}

	func (me *PkgBag) insertFmt (index int, format string, fmtArgs ... interface{}) {
		me.lines = append(me.lines[: index], append([]string { fmt.Sprintf(format, fmtArgs ...) }, me.lines[index : ] ...) ...)
	}

	func (me *PkgBag) reinit () {
		me.impName = "xsdt"
		me.imports, me.impsUsed, me.lines = map[string]string {}, map[string]bool {}, []string { "package gopkg_" + me.safeName(me.Schema.RootSchema().loadUri), "" }
	}

	func (me *PkgBag) resolveTypeRef (typeRef string) string {
		var ns = me.Schema.XMLNamespaces[""]
		var impName = ""
		if len(typeRef) == 0 { return "" }
		if pos := strings.Index(typeRef, ":"); pos > 0 {
			impName, ns = typeRef[: pos], me.Schema.XMLNamespaces[typeRef[: pos]]
			typeRef = typeRef[pos + 1 :]
		}
		if ns == xsdNamespaceUri { impName = me.impName }
		if ns == me.Schema.TargetNamespace.String() { impName = "" }
		me.impsUsed[impName] = true
		return ustr.PrefixWithSep(impName, ".", me.safeName(typeRef))
	}

	func (me *PkgBag) safeName (name string) string {
		return ustr.SafeIdentifier(name)
	}
