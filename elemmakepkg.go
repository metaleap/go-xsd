package xsd

import (
	"fmt"
	"path"
	"strings"
	"time"

	util "github.com/metaleap/go-util"
	ustr "github.com/metaleap/go-util/str"

	xsdt "github.com/metaleap/go-xsd/types"
)

const (
	idPrefix = "XsdGoPkg"
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
	var safeName, typeName, tmp, key, defVal, impName string
	var defName = "Default"
	me.elemBase.beforeMakePkg(bag)
	if len(me.Form) == 0 { me.Form = bag.Schema.AttributeFormDefault }
	me.hasElemsSimpleType.makePkg(bag)
	if len(me.Ref) > 0 {
		key = bag.resolveQnameRef(me.Ref.String(), "", &impName)
		tmp = ustr.PrefixWithSep(impName, ".", idPrefix + "HasAttr_" + bag.safeName(me.Ref.String()[(strings.Index(me.Ref.String(), ":") + 1) :]))
		if perPkgState.attRefImps[me], perPkgState.attsKeys[me] = impName, key; len(perPkgState.attsCache[key]) == 0 { perPkgState.attsCache[key] = tmp }
	} else {
		safeName = bag.safeName(me.Name.String())
		if typeName = me.Type.String(); (len(typeName) == 0) && (len(me.SimpleTypes) > 0) {
			typeName = me.SimpleTypes[0].Name.String()
		} else {
			if len(typeName) == 0 { typeName = bag.xsdStringTypeRef() }
			typeName = bag.resolveQnameRef(typeName, "T", &impName)
		}
		if defVal = me.Default; len(defVal) == 0 { defName, defVal = "Fixed", me.Fixed }
		if me.Parent() == bag.Schema { key = safeName } else { key = safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal) }
		if len(perPkgState.attsCache[key]) == 0 {
			bag.impsUsed[impName] = true
			tmp = idPrefix + "HasAttr_" + key
			perPkgState.attsKeys[me] = key
			perPkgState.attsCache[key] = tmp
			me.hasElemAnnotation.makePkg(bag)
			bag.appendFmt(false, "type %v struct {", tmp)
			me.hasElemAnnotation.makePkg(bag)
			bag.appendFmt(false, "\t%v %v `xml:\"%v,attr\"`", safeName, typeName, util.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String() + " ", "") + me.Name.String())
			bag.appendFmt(true, "}")
			if isPt := bag.isParseType(typeName); len(defVal) > 0 {
				bag.appendFmt(false, "//\tReturns the %v value for %v -- " + util.Ifs(isPt, "%v", "%#v"), defName, safeName, defVal)
				if isPt {
					bag.appendFmt(true, "func (me *%v) %v%v () %v { return %v(%v) }", tmp, safeName, defName, typeName, typeName, defVal)
				} else {
					bag.appendFmt(true, "func (me *%v) %v%v () %v { return %v(%#v) }", tmp, safeName, defName, typeName, typeName, defVal)
				}
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *AttributeGroup) makePkg (bag *PkgBag) {
	var refName, refImp string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	if len(me.Ref) > 0 {
		if len(perPkgState.attGroups[me]) == 0 {
			refName = bag.resolveQnameRef(me.Ref.String(), "", &refImp)
			perPkgState.attGroups[me] = idPrefix + "HasAtts_" + refName
			perPkgState.attGroupRefImps[me] = refImp
		}
	} else {
		me.hasElemAnnotation.makePkg(bag)
		safeName := bag.safeName(me.Name.String())
		tmp := idPrefix + "HasAtts_" + safeName
		bag.appendFmt(false, "type %v struct {", tmp)
		perPkgState.attGroups[me] = tmp
		for _, ag := range me.AttributeGroups {
			if len(ag.Ref) == 0 { ag.Ref.SetFromString(ag.Name.String()) }
			ag.hasElemAnnotation.makePkg(bag)
			if refName = bag.resolveQnameRef(ag.Ref.String(), "", &refImp); len(refImp) > 0 {
				bag.impsUsed[refImp] = true
				bag.appendFmt(true, "\t%v.%vHasAtts_%v", refImp, idPrefix, refName[(len(refImp) + 1) :])
			} else {
				bag.appendFmt(true, "\t%vHasAtts_%v", idPrefix, refName)
			}
		}
		for _, att := range me.Attributes {
			if len(att.Ref) > 0 { att.hasElemAnnotation.makePkg(bag) }
			if imp := perPkgState.attRefImps[att]; len(imp) > 0 { bag.impsUsed[imp] = true }
			if key := perPkgState.attsKeys[att]; len(key) > 0 { bag.appendFmt(false, "\t%v", perPkgState.attsCache[key]) }
		}
		bag.appendFmt(true, "}")
	}
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
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemComplexContent.makePkg(bag)
	me.hasElemSimpleContent.makePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	if len(me.Name) == 0 { me.Name = bag.AnonName(me.longSafeName(bag)) }
	bag.appendFmt(false, "type %v struct {", bag.safeName(ustr.PrependIf(me.Name.String(), "T")))
	if me.Mixed { bag.appendFmt(true, "\t%vHasCdata", idPrefix) }
	for _, ag := range me.AttributeGroups {
		ag.hasElemAnnotation.makePkg(bag)
		bag.appendFmt(true, "\t%v", ustr.PrefixWithSep(perPkgState.attGroupRefImps[ag], ".", perPkgState.attGroups[ag][(strings.Index(perPkgState.attGroups[ag], ".") + 1) :]))
		bag.impsUsed[perPkgState.attGroupRefImps[ag]] = true
	}
	for _, at := range me.Attributes {
		if key := perPkgState.attsKeys[at]; len(key) > 0 {
			at.hasElemAnnotation.makePkg(bag)
			bag.appendFmt(true, "\t%v", ustr.PrefixWithSep(perPkgState.attRefImps[at], ".", perPkgState.attsCache[key][(strings.Index(perPkgState.attsCache[key], ".") + 1) :]))
			bag.impsUsed[perPkgState.attRefImps[at]] = true
		}
	}
	bag.appendFmt(true, "}")
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
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemUnique.makePkg(bag)
	me.hasElemsKey.makePkg(bag)
	me.hasElemComplexType.makePkg(bag)
	me.hasElemKeyRef.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
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
	var rtr = bag.resolveQnameRef(me.ItemType.String(), "T", nil)
	var safeName string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	if len(rtr) == 0 {
		rtr = me.SimpleTypes[0].Name.String()
		if len(rtr) == 0 {
			rtr = bag.safeName(bag.Stacks.CurSimpleType().Name.String() + "SubList")
			me.SimpleTypes[0].Name = xsdt.NCName(rtr)
		}
	}
	bag.impsUsed[bag.impName] = true
	safeName = bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T"))
	bag.appendFmt(false, "//\t%v defines a String containing a whitespace-separated list of %v values. This Values() method creates and returns a slice of all elements in that list.", safeName, rtr)
	if bag.isParseType(rtr) {
		bag.appendFmt(false, `func (me %v) Values () (list []%v) {
	var btv = new(%v)
	var svals = xsdt.ListValues(string(me))
	list = make([]%v, len(svals))
	for i, s := range svals { btv.SetFromString(s); list[i] = *btv }
	return
}
		`, safeName, rtr, rtr, rtr)
	} else {
		bag.appendFmt(false, `func (me %v) Values () (list []%v) {
	var svals = xsdt.ListValues(string(me))
	list = make([]%v, len(svals))
	for i, s := range svals { list[i] = %v(s) }
	return
}
		`, safeName, rtr, rtr, rtr)
	}
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
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
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
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleEnumeration) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	safeName := bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T"))
	bag.appendFmt(false, "//\tReturns true if the value of this enumerated %v is %#v.", safeName, me.Value)
	bag.appendFmt(true, "func (me %v) Is%v () bool { return me == %#v }", safeName, bag.safeName(me.Value), me.Value)
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
	me.hasElemsSimpleType.makePkg(bag)
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
	me.hasElemAnnotation.makePkg(bag)
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
	bag.ParseTypes = []string {}
	for _, pt := range []string { "Boolean", "Byte", "Double", "Float", "Int", "Integer", "Long", "NegativeInteger", "NonNegativeInteger", "NonPositiveInteger", "PositiveInteger", "Short", "UnsignedByte", "UnsignedInt", "UnsignedLong", "UnsignedShort" } {
		bag.ParseTypes = append(bag.ParseTypes, bag.impName + "." + pt)
	}
	// for _, im := range []string { "strings" } { bag.imports[im] = "" }
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsImport.makePkg(bag)
	impPos = len(bag.lines) + 1
	bag.append("import (", ")", "")
	if me.XSDParentSchema == nil {
		perPkgState.anonCounts = map[string]uint64 {}
		perPkgState.attGroups = map[*AttributeGroup]string {}
		perPkgState.attGroupRefImps = map[*AttributeGroup]string {}
		perPkgState.attsCache = map[string]string {}
		perPkgState.attsKeys = map[*Attribute]string {}
		perPkgState.attRefImps = map[*Attribute]string {}
		bag.appendFmt(true, "type %vHasCdata struct { CDATA string `xml:\",chardata\"` }", idPrefix)
	}
	me.hasElemsNotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsRedefine.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
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
	var resolve = true
	var isPt bool
	if len(typeName) == 0 { typeName = bag.AnonName(me.longSafeName(bag)); me.Name = typeName } else { me.Name = typeName }
	typeName = xsdt.NCName(ustr.PrependIf(typeName.String(), "T"))
	me.elemBase.beforeMakePkg(bag)
	bag.Stacks.SimpleType.Push(me)
	safeName = bag.safeName(typeName.String())
	me.hasElemAnnotation.makePkg(bag)
	if me.RestrictionSimpleType != nil {
		if baseType = me.RestrictionSimpleType.Base.String(); (len(baseType) == 0) && (len(me.RestrictionSimpleType.SimpleTypes) > 0) {
			resolve, baseType = false, me.RestrictionSimpleType.SimpleTypes[0].Name.String()
		}
	}
	if len(baseType) == 0 { baseType = bag.xsdStringTypeRef() }
	if resolve { baseType = bag.resolveQnameRef(baseType, "T", nil) }
	if isPt = bag.isParseType(baseType); isPt { bag.ParseTypes = append(bag.ParseTypes, safeName) }
	bag.appendFmt(true, "type %s %s", safeName, baseType)
	if isPt {
		bag.appendFmt(false, "//\tSince %v is a non-string scalar type (either boolean or numeric), sets the current value obtained from parsing the specified string.", safeName)
	} else {
		bag.appendFmt(false, "//\tSince %v is just a simple String type, this merely sets the current value from the specified string.", safeName)
	}
	bag.appendFmt(true, "func (me *%s) SetFromString (s string) { (*%v)(me).SetFromString(s) }", safeName, baseType)
	if isPt {
		bag.appendFmt(false, "//\tReturns a string representation of this %v's current non-string scalar value.", safeName)
	} else {
		bag.appendFmt(false, "//\tSince %v is just a simple String type, this merely returns the current string value.", safeName)
	}
	bag.appendFmt(true, "func (me %s) String () string { return %v(me).String() }", safeName, baseType)
	bag.appendFmt(false, "//\tThis convenience method just performs a simple type conversion to %v's alias type %v", safeName, baseType)
	bag.appendFmt(true, "func (me %s) To%v () %v { return %v(me) }", safeName, bag.safeName(baseType), baseType, baseType)
	me.hasElemRestrictionSimpleType.makePkg(bag)
	me.hasElemList.makePkg(bag)
	me.hasElemUnion.makePkg(bag)
	bag.Stacks.SimpleType.Pop()
	me.elemBase.afterMakePkg(bag)
}

func (me *Union) makePkg (bag *PkgBag) {
	var memberTypes []string
	var rtn, rtnSafeName, safeName string
	var isParseType = false
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(me.MemberTypes, " ")
	for _, st := range me.SimpleTypes { memberTypes = append(memberTypes, st.Name.String()) }
	for _, mt := range memberTypes {
		rtn = bag.resolveQnameRef(mt, "T", nil)
		safeName, rtnSafeName = bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T")), bag.safeName(rtn)
		bag.appendFmt(false, "//\t%v is an XSD union type of several types. This is a simple type conversion to %v, but keep in mind the actual value may or may not be a valid %v value.", safeName, rtnSafeName, rtnSafeName)
		if isParseType = bag.isParseType(rtn); isParseType {
			bag.appendFmt(true, "func (me %v) To%v () %v { var x = new(%v); x.SetFromString(me.String()); return *x }", safeName, rtnSafeName, rtn, rtn)
		} else {
			bag.appendFmt(true, "func (me %v) To%v () %v { return %v(me) }", safeName, rtnSafeName, rtn, rtn)
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
