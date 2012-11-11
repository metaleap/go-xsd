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
	me.elemBase.afterMakePkg(bag)
}

func (me *AnyAttribute) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
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
					if PkgGen.ForceParseForDefaults {
						bag.appendFmt(true, "func (me *%v) %v%v () %v { var x = new(%v); x.SetFromString(%#v); return *x }", tmp, safeName, defName, typeName, typeName, defVal)
					} else {
						bag.appendFmt(true, "func (me *%v) %v%v () %v { return %v(%v) }", tmp, safeName, defName, typeName, typeName, defVal)
					}
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
			if imp := perPkgState.attRefImps[att]; len(imp) > 0 { bag.impsUsed[imp] = true }
			if key := perPkgState.attsKeys[att]; len(key) > 0 {
				att.hasElemAnnotation.makePkg(bag)
				bag.appendFmt(false, "\t%v", perPkgState.attsCache[key])
			}
		}
		bag.appendFmt(true, "}")
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Choice) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemExtensionComplexContent.makePkg(bag)
	me.hasElemRestrictionComplexContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexType) makePkg (bag *PkgBag) {
	var att *Attribute
	var attGroup *AttributeGroup
	var ctBaseType, ctValueType, typeSafeName string
	var allAtts = map[*Attribute]bool {}
	var allAttGroups = map[*AttributeGroup]bool {}
	var allElems = map[*Element]bool {}
	var allElemGroups = map[*Group]bool {}
	var elsDone, grsDone = map[string]bool {}, map[string]bool {}
	var allChoices, tmpChoices = []*Choice {}, []*Choice { me.Choice }
	var allSeqs, tmpSeqs = []*Sequence {}, []*Sequence { me.Sequence }
	var elCache = perPkgState.elemsCacheOnce
	var el *Element
	var elGr *Group
	var mixed = false
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemGroup.makePkg(bag)
	me.hasElemSequence.makePkg(bag)
	me.hasElemComplexContent.makePkg(bag)
	me.hasElemSimpleContent.makePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	if len(me.Name) == 0 { me.Name = bag.AnonName(me.longSafeName(bag)) }
	typeSafeName = bag.safeName(ustr.PrependIf(me.Name.String(), "T"))
	bag.appendFmt(false, "type %v struct {", typeSafeName)
	for _, att = range me.Attributes { allAtts[att] = true }
	for _, attGroup = range me.AttributeGroups { allAttGroups[attGroup] = true }
	allChoices, allSeqs = Flattened(tmpChoices, tmpSeqs)
	if me.All != nil { for _, el = range me.All.Elements { allElems[el] = true } }
	if me.Group != nil { allElemGroups[me.Group] = true }
	if mixed = me.Mixed; me.ComplexContent != nil {
		mixed = mixed || me.ComplexContent.Mixed
		me.ComplexContent.hasElemAnnotation.makePkg(bag)
		if me.ComplexContent.ExtensionComplexContent != nil {
			me.ComplexContent.ExtensionComplexContent.hasElemAnnotation.makePkg(bag)
			if me.ComplexContent.ExtensionComplexContent.All != nil { for _, el = range me.ComplexContent.ExtensionComplexContent.All.Elements { allElems[el] = true } }
			for _, elGr = range me.ComplexContent.ExtensionComplexContent.Groups { allElemGroups[elGr] = true }
			tmpChoices, tmpSeqs = Flattened(me.ComplexContent.ExtensionComplexContent.Choices, me.ComplexContent.ExtensionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices ...), append(allSeqs, tmpSeqs ...)
			for _, att = range me.ComplexContent.ExtensionComplexContent.Attributes { allAtts[att] = true }
			for _, attGroup = range me.ComplexContent.ExtensionComplexContent.AttributeGroups { allAttGroups[attGroup] = true }
			if len(me.ComplexContent.ExtensionComplexContent.Base) > 0 { ctBaseType = me.ComplexContent.ExtensionComplexContent.Base.String() }
		}
		if me.ComplexContent.RestrictionComplexContent != nil {
			me.ComplexContent.RestrictionComplexContent.hasElemAnnotation.makePkg(bag)
			if me.ComplexContent.RestrictionComplexContent.All != nil { for _, el = range me.ComplexContent.RestrictionComplexContent.All.Elements { allElems[el] = true } }
			tmpChoices, tmpSeqs = Flattened(me.ComplexContent.RestrictionComplexContent.Choices, me.ComplexContent.RestrictionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices ...), append(allSeqs, tmpSeqs ...)
			for _, att = range me.ComplexContent.RestrictionComplexContent.Attributes { allAtts[att] = true }
			for _, attGroup = range me.ComplexContent.RestrictionComplexContent.AttributeGroups { allAttGroups[attGroup] = true }
			if len(me.ComplexContent.RestrictionComplexContent.Base) > 0 { ctBaseType = me.ComplexContent.RestrictionComplexContent.Base.String() }
		}
	}
	if me.SimpleContent != nil {
		me.SimpleContent.hasElemAnnotation.makePkg(bag)
		if me.SimpleContent.ExtensionSimpleContent != nil {
			me.SimpleContent.ExtensionSimpleContent.hasElemAnnotation.makePkg(bag)
			for _, att = range me.SimpleContent.ExtensionSimpleContent.Attributes { allAtts[att] = true }
			for _, attGroup = range me.SimpleContent.ExtensionSimpleContent.AttributeGroups { allAttGroups[attGroup] = true }
			if (len(ctValueType) == 0) && (len(me.SimpleContent.ExtensionSimpleContent.Base) > 0) { ctValueType = me.SimpleContent.ExtensionSimpleContent.Base.String() }
		}
		if me.SimpleContent.RestrictionSimpleContent != nil {
			me.SimpleContent.RestrictionSimpleContent.hasElemAnnotation.makePkg(bag)
			for _, att = range me.SimpleContent.RestrictionSimpleContent.Attributes { allAtts[att] = true }
			for _, attGroup = range me.SimpleContent.RestrictionSimpleContent.AttributeGroups { allAttGroups[attGroup] = true }
			if (len(ctValueType) == 0) && (len(me.SimpleContent.RestrictionSimpleContent.Base) > 0) { ctValueType = me.SimpleContent.RestrictionSimpleContent.Base.String() }
			if (len(ctValueType) == 0) && (len(me.SimpleContent.RestrictionSimpleContent.SimpleTypes) > 0) { ctValueType = me.SimpleContent.RestrictionSimpleContent.SimpleTypes[0].Name.String() }
			for _, enum := range me.SimpleContent.RestrictionSimpleContent.Enumerations {
				println("ENUMTODO!! " + enum.selfName().String())
			}
		}
	}
	if ctBaseType = bag.resolveQnameRef(ctBaseType, "T", nil); len(ctBaseType) > 0 {
		bag.appendFmt(true, "\t%v", bag.safeName(ctBaseType))
	}
	if ctValueType = bag.resolveQnameRef(ctValueType, "T", nil); len(ctValueType) > 0 {
		perPkgState.simpleContentValueTypes[typeSafeName] = ctValueType
		bag.appendFmt(true, "\t%vValue %v `xml:\",chardata\"`", idPrefix, ctValueType)
	} else if mixed {
		bag.appendFmt(true, "\t%vHasCdata", idPrefix)
	}
	for elGr, _ = range allElemGroups { subMakeElemGroup(bag, elGr, grsDone) }
	for el, _ = range allElems { subMakeElem(bag, el, elsDone, elCache) }
	for _, ch := range allChoices {
		if ch.MaxOccurs == 1 { elCache = perPkgState.elemsCacheOnce } else { elCache = perPkgState.elemsCacheMult }
		ch.hasElemAnnotation.makePkg(bag); for _, el = range ch.Elements { subMakeElem(bag, el, elsDone, elCache) }
		for _, elGr = range ch.Groups { subMakeElemGroup(bag, elGr, grsDone) }
	}
	for _, seq := range allSeqs {
		if seq.MaxOccurs == 1 { elCache = perPkgState.elemsCacheOnce } else { elCache = perPkgState.elemsCacheMult }
		seq.hasElemAnnotation.makePkg(bag); for _, el = range seq.Elements { subMakeElem(bag, el, elsDone, elCache) }
		for _, elGr = range seq.Groups { subMakeElemGroup(bag, elGr, grsDone) }
	}
	for attGroup, _ = range allAttGroups {
		attGroup.hasElemAnnotation.makePkg(bag)
		bag.appendFmt(true, "\t%v", ustr.PrefixWithSep(perPkgState.attGroupRefImps[attGroup], ".", perPkgState.attGroups[attGroup][(strings.Index(perPkgState.attGroups[attGroup], ".") + 1) :]))
		bag.impsUsed[perPkgState.attGroupRefImps[attGroup]] = true
	}
	for att, _ = range allAtts {
		if key := perPkgState.attsKeys[att]; len(key) > 0 {
			att.hasElemAnnotation.makePkg(bag)
			bag.appendFmt(true, "\t%v", ustr.PrefixWithSep(perPkgState.attRefImps[att], ".", perPkgState.attsCache[key][(strings.Index(perPkgState.attsCache[key], ".") + 1) :]))
			bag.impsUsed[perPkgState.attRefImps[att]] = true
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
	var safeName, typeName, valueType, tmp, key, defVal, impName string
	var asterisk, defName = "", "Default"
	me.elemBase.beforeMakePkg(bag)
	if len(me.Form) == 0 { me.Form = bag.Schema.ElementFormDefault }
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemComplexType.makePkg(bag)
	if len(me.Ref) > 0 {
		key = bag.resolveQnameRef(me.Ref.String(), "", &impName)
		for pref, cache := range map[string]map[string]string { "HasElem_": perPkgState.elemsCacheOnce, "HasElems_": perPkgState.elemsCacheMult } {
			tmp = ustr.PrefixWithSep(impName, ".", idPrefix + pref + bag.safeName(me.Ref.String()[(strings.Index(me.Ref.String(), ":") + 1) :]))
			if perPkgState.elemRefImps[me], perPkgState.elemKeys[me] = impName, key; len(cache[key]) == 0 { cache[key] = tmp }
		}
	} else {
		safeName = bag.safeName(me.Name.String())
		if typeName = me.Type.String(); (len(typeName) == 0) && ((me.ComplexType != nil) || (len(me.SimpleTypes) > 0)) {
			if me.ComplexType != nil { asterisk, typeName = "*", me.ComplexType.Name.String() } else { typeName = me.SimpleTypes[0].Name.String() }
		} else {
			if len(typeName) == 0 { typeName = bag.xsdStringTypeRef() }
			if typeName = bag.resolveQnameRef(typeName, "T", &impName); bag.Schema.globalComplexType(bag, typeName) != nil { asterisk = "*" }
		}
		if defVal = me.Default; len(defVal) == 0 { defName, defVal = "Fixed", me.Fixed }
		if me.Parent() == bag.Schema { key = safeName } else { key = safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal) }
		if valueType = perPkgState.simpleContentValueTypes[typeName]; len(valueType) == 0 { valueType = typeName }
		for pref, cache := range map[string]map[string]string { "HasElem_": perPkgState.elemsCacheOnce, "HasElems_": perPkgState.elemsCacheMult } {
			if tmp = idPrefix + pref + key; !perPkgState.elemsWritten[tmp] {
				perPkgState.elemsWritten[tmp], perPkgState.elemKeys[me] = true, key
				bag.impsUsed[impName] = true
				cache[key] = tmp
				me.hasElemAnnotation.makePkg(bag)
				{	//	these constraints aren't interpreted for makePkg() but their doc annotations might be worth including
					me.hasElemUnique.makePkg(bag)
					me.hasElemsKey.makePkg(bag)
					me.hasElemKeyRef.makePkg(bag)
				}
				bag.appendFmt(false, "type %v struct {", tmp)
				me.hasElemAnnotation.makePkg(bag)
				bag.appendFmt(false, "\t%v %v `xml:\"%v\"`", util.Ifs(pref == "HasElems_", ustr.Pluralize(safeName), safeName), util.Ifs(pref == "HasElems_", "[]" + asterisk + typeName, asterisk + typeName), util.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String() + " ", "") + me.Name.String())
				bag.appendFmt(true, "}")
				fmt.Sprintf(defName)
				if len(defVal) > 0 {
					isPt := bag.isParseType(valueType)
					bag.appendFmt(false, "//\tReturns the %v value for %v -- " + util.Ifs(isPt, "%v", "%#v"), defName, safeName, defVal)
					if isPt {
						if PkgGen.ForceParseForDefaults {
							bag.appendFmt(true, "func (me *%v) %v%v () %v { var x = new(%v); x.SetFromString(%#v); return *x }", tmp, safeName, defName, valueType, valueType, defVal)
						} else {
							bag.appendFmt(true, "func (me *%v) %v%v () %v { return %v(%v) }", tmp, safeName, defName, valueType, valueType, defVal)
						}
					} else {
						bag.appendFmt(true, "func (me *%v) %v%v () %v { return %v(%#v) }", tmp, safeName, defName, valueType, valueType, defVal)
					}
				}
			}
		}
	}



	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
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
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Field) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Group) makePkg (bag *PkgBag) {
	var refName, refImp string
	var choices = []*Choice { me.Choice }
	var seqs = []*Sequence { me.Sequence }
	var el *Element
	var gr *Group
	var elsDone, grsDone = map[string]bool {}, map[string]bool {}
	var elCache map[string]string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemSequence.makePkg(bag)
	if len(me.Ref) > 0 {
		if len(perPkgState.elemGroups[me]) == 0 {
			refName = bag.resolveQnameRef(me.Ref.String(), "", &refImp)
			perPkgState.elemGroups[me] = idPrefix + "HasGroup_" + refName
			perPkgState.elemGroupRefImps[me] = refImp
		}
	} else {
		me.Ref.SetFromString(me.Name.String())
		safeName := bag.safeName(me.Name.String())
		tmp := idPrefix + "HasGroup_" + safeName
		perPkgState.elemGroups[me] = tmp
		me.hasElemAnnotation.makePkg(bag)
		bag.appendFmt(false, "type %v struct {", tmp)
		choices, seqs = Flattened(choices, seqs)
		elCache = perPkgState.elemsCacheOnce
		if me.All != nil { me.All.hasElemAnnotation.makePkg(bag); for _, el = range me.All.Elements { subMakeElem(bag, el, elsDone, elCache) } }
		for _, ch := range choices {
			if ch.MaxOccurs == 1 { elCache = perPkgState.elemsCacheOnce } else { elCache = perPkgState.elemsCacheMult }
			ch.hasElemAnnotation.makePkg(bag); for _, el = range ch.Elements { subMakeElem(bag, el, elsDone, elCache) }
			for _, gr = range ch.Groups { subMakeElemGroup(bag, gr, grsDone) }
		}
		for _, seq := range seqs {
			if seq.MaxOccurs == 1 { elCache = perPkgState.elemsCacheOnce } else { elCache = perPkgState.elemsCacheMult }
			seq.hasElemAnnotation.makePkg(bag); for _, el = range seq.Elements { subMakeElem(bag, el, elsDone, elCache) }
			for _, gr = range seq.Groups { subMakeElemGroup(bag, gr, grsDone) }
		}
		bag.appendFmt(true, "}")
	}

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
			rtr = bag.safeName(bag.AnonName(bag.Stacks.CurSimpleType().Name.String()).String())
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
		perPkgState.simpleContentValueTypes, perPkgState.attsCache, perPkgState.elemsCacheOnce, perPkgState.elemsCacheMult = map[string]string {}, map[string]string {}, map[string]string {}, map[string]string {}
		perPkgState.attGroups, perPkgState.attGroupRefImps = map[*AttributeGroup]string {}, map[*AttributeGroup]string {}
		perPkgState.attsKeys, perPkgState.attRefImps = map[*Attribute]string {}, map[*Attribute]string {}
		perPkgState.elemGroups, perPkgState.elemGroupRefImps = map[*Group]string {}, map[*Group]string {}
		perPkgState.elemKeys, perPkgState.elemRefImps = map[*Element]string {}, map[*Element]string {}
		perPkgState.elemsWritten = map[string]bool {}
		bag.appendFmt(true, "type %vHasCdata struct { CombinedCharDatas string `xml:\",chardata\"` }", idPrefix)
	}
	me.hasElemsNotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsRedefine.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
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
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
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
	me.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(me.MemberTypes, " ")
	for _, st := range me.SimpleTypes { memberTypes = append(memberTypes, st.Name.String()) }
	me.hasElemAnnotation.makePkg(bag)
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

func subMakeElem (bag *PkgBag, el *Element, done map[string]bool, elCache map[string]string) {
	var refImp = perPkgState.elemRefImps[el]
	if len(refImp) > 0 { bag.impsUsed[refImp] = true }
	if refName := perPkgState.elemKeys[el]; (len(refName) > 0) && (!done[refName]) {
		el.hasElemAnnotation.makePkg(bag); done[refName] = true
		if !strings.HasPrefix(elCache[refName], bag.impName + "." + idPrefix) {
			bag.appendFmt(true, "\t%v", elCache[refName])
		}
	}
}

func subMakeElemGroup (bag *PkgBag, gr *Group, done map[string]bool) {
	var refImp string
	if refName := bag.resolveQnameRef(gr.Ref.String(), "", &refImp); !done[refName] {
		gr.hasElemAnnotation.makePkg(bag)
		if done[refName] = true; len(refImp) > 0 {
			if bag.impsUsed[refImp] = true; !strings.HasPrefix(refName, bag.impName + "." + idPrefix) {
				bag.appendFmt(true, "\t%v.%vHasGroup_%v", refImp, idPrefix, refName[(len(refImp) + 1) :])
			}
		} else {
			bag.appendFmt(true, "\t%vHasGroup_%v", idPrefix, refName)
		}
	}
}
