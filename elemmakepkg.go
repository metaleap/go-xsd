package xsd

import (
	"fmt"
	"path"
	"strings"
	"unicode"

	"github.com/metaleap/go-util-str"

	xsdt "github.com/metaleap/go-xsd/types"
)

const (
	idPrefix = "XsdGoPkg"
)

func (me *All) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Annotation) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAppInfo.makePkg(bag)
	me.hasElemsDocumentation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Any) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AnyAttribute) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AppInfo) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Attribute) makePkg(bag *PkgBag) {
	var safeName, typeName, tmp, key, defVal, impName string
	var defName = "Default"
	me.elemBase.beforeMakePkg(bag)
	if len(me.Form) == 0 {
		me.Form = bag.Schema.AttributeFormDefault
	}
	me.hasElemsSimpleType.makePkg(bag)
	if len(me.Ref) > 0 {
		key = bag.resolveQnameRef(me.Ref.String(), "", &impName)
		tmp = ustr.PrefixWithSep(impName, ".", idPrefix+"HasAttr_"+bag.safeName(me.Ref.String()[(strings.Index(me.Ref.String(), ":")+1):]))
		if bag.attRefImps[me], bag.attsKeys[me] = impName, key; len(bag.attsCache[key]) == 0 {
			bag.attsCache[key] = tmp
		}
	} else {
		safeName = bag.safeName(me.Name.String())
		if typeName = me.Type.String(); (len(typeName) == 0) && (len(me.SimpleTypes) > 0) {
			typeName = me.SimpleTypes[0].Name.String()
		} else {
			if len(typeName) == 0 {
				typeName = bag.xsdStringTypeRef()
			}
			typeName = bag.resolveQnameRef(typeName, "T", &impName)
		}
		if defVal = me.Default; len(defVal) == 0 {
			defName, defVal = "Fixed", me.Fixed
		}
		if me.Parent() == bag.Schema {
			key = safeName
		} else {
			key = safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal)
		}
		if len(bag.attsCache[key]) == 0 {
			tmp = idPrefix + "HasAttr_" + key
			bag.attsKeys[me] = key
			bag.attsCache[key] = tmp
			var td = bag.addType(me, tmp, "", me.Annotation)
			td.addField(me, safeName, typeName, ustr.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String()+" ", "")+me.Name.String()+",attr", me.Annotation)
			if isPt := bag.isParseType(typeName); len(defVal) > 0 {
				doc := sfmt("Returns the %v value for %v -- "+ustr.Ifs(isPt, "%v", "%#v"), strings.ToLower(defName), safeName, defVal)
				if isPt {
					if PkgGen.ForceParseForDefaults {
						td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("var x = new(%v); x.Set(%#v); return *x", typeName, defVal), doc)
					} else {
						td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("return %v(%v)", typeName, defVal), doc)
					}
				} else {
					td.addMethod(nil, tmp, safeName+defName, typeName, sfmt("return %v(%#v)", typeName, defVal), doc)
				}
			}
		} else {
			bag.attsKeys[me] = key
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *AttributeGroup) makePkg(bag *PkgBag) {
	var refName, refImp string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	if len(me.Ref) > 0 {
		if len(bag.attGroups[me]) == 0 {
			refName = bag.resolveQnameRef(me.Ref.String(), "", &refImp)
			bag.attGroups[me] = idPrefix + "HasAtts_" + refName
			bag.attGroupRefImps[me] = refImp
		}
	} else {
		safeName := bag.safeName(me.Name.String())
		tmp := idPrefix + "HasAtts_" + safeName
		var td = bag.addType(me, tmp, "", me.Annotation)
		bag.attGroups[me] = tmp
		for _, ag := range me.AttributeGroups {
			if len(ag.Ref) == 0 {
				ag.Ref.Set(ag.Name.String())
			}
			if refName = bag.resolveQnameRef(ag.Ref.String(), "", &refImp); len(refImp) > 0 {
				td.addEmbed(ag, refImp+"."+idPrefix+"HasAtts_"+refName[(len(refImp)+1):], ag.Annotation)
			} else {
				td.addEmbed(ag, idPrefix+"HasAtts_"+refName, ag.Annotation)
			}
		}
		for _, att := range me.Attributes {
			if key := bag.attsKeys[att]; len(key) > 0 {
				td.addEmbed(att, bag.attsCache[key], att.Annotation)
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Choice) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexContent) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemExtensionComplexContent.makePkg(bag)
	me.hasElemRestrictionComplexContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexType) makePkg(bag *PkgBag) {
	var att *Attribute
	var attGroup *AttributeGroup
	var ctBaseType, ctValueType, typeSafeName string
	var allAtts = map[*Attribute]bool{}
	var allAttGroups = map[*AttributeGroup]bool{}
	var allElems = map[*Element]bool{}
	var allElemGroups = map[*Group]bool{}
	var elsDone, grsDone = map[string]bool{}, map[string]bool{}
	var allChoices, tmpChoices = []*Choice{}, []*Choice{me.Choice}
	var allSeqs, tmpSeqs = []*Sequence{}, []*Sequence{me.Sequence}
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
	if len(me.Name) == 0 {
		me.Name = bag.AnonName(me.longSafeName(bag))
	}
	typeSafeName = bag.safeName(ustr.PrependIf(me.Name.String(), "T"))
	var td = bag.addType(me, typeSafeName, "", me.Annotation)
	for _, att = range me.Attributes {
		allAtts[att] = true
	}
	for _, attGroup = range me.AttributeGroups {
		allAttGroups[attGroup] = true
	}
	allChoices, allSeqs = Flattened(tmpChoices, tmpSeqs)
	if me.All != nil {
		for _, el = range me.All.Elements {
			allElems[el] = true
		}
	}
	if me.Group != nil {
		allElemGroups[me.Group] = true
	}
	if mixed = me.Mixed; me.ComplexContent != nil {
		mixed = mixed || me.ComplexContent.Mixed
		td.addAnnotations(me.ComplexContent.Annotation)
		if me.ComplexContent.ExtensionComplexContent != nil {
			td.addAnnotations(me.ComplexContent.ExtensionComplexContent.Annotation)
			if me.ComplexContent.ExtensionComplexContent.All != nil {
				for _, el = range me.ComplexContent.ExtensionComplexContent.All.Elements {
					allElems[el] = true
				}
			}
			for _, elGr = range me.ComplexContent.ExtensionComplexContent.Groups {
				allElemGroups[elGr] = true
			}
			tmpChoices, tmpSeqs = Flattened(me.ComplexContent.ExtensionComplexContent.Choices, me.ComplexContent.ExtensionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices...), append(allSeqs, tmpSeqs...)
			for _, att = range me.ComplexContent.ExtensionComplexContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range me.ComplexContent.ExtensionComplexContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if len(me.ComplexContent.ExtensionComplexContent.Base) > 0 {
				ctBaseType = me.ComplexContent.ExtensionComplexContent.Base.String()
			}
		}
		if me.ComplexContent.RestrictionComplexContent != nil {
			td.addAnnotations(me.ComplexContent.RestrictionComplexContent.Annotation)
			if me.ComplexContent.RestrictionComplexContent.All != nil {
				for _, el = range me.ComplexContent.RestrictionComplexContent.All.Elements {
					allElems[el] = true
				}
			}
			tmpChoices, tmpSeqs = Flattened(me.ComplexContent.RestrictionComplexContent.Choices, me.ComplexContent.RestrictionComplexContent.Sequences)
			allChoices, allSeqs = append(allChoices, tmpChoices...), append(allSeqs, tmpSeqs...)
			for _, att = range me.ComplexContent.RestrictionComplexContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range me.ComplexContent.RestrictionComplexContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if len(me.ComplexContent.RestrictionComplexContent.Base) > 0 {
				ctBaseType = me.ComplexContent.RestrictionComplexContent.Base.String()
			}
		}
	}
	if me.SimpleContent != nil {
		td.addAnnotations(me.SimpleContent.Annotation)
		if me.SimpleContent.ExtensionSimpleContent != nil {
			if len(me.SimpleContent.ExtensionSimpleContent.Base) > 0 {
				ctBaseType = me.SimpleContent.ExtensionSimpleContent.Base.String()
			}
			td.addAnnotations(me.SimpleContent.ExtensionSimpleContent.Annotation)
			for _, att = range me.SimpleContent.ExtensionSimpleContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range me.SimpleContent.ExtensionSimpleContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if (len(ctValueType) == 0) && (len(me.SimpleContent.ExtensionSimpleContent.Base) > 0) {
				ctValueType = me.SimpleContent.ExtensionSimpleContent.Base.String()
			}
		}
		if me.SimpleContent.RestrictionSimpleContent != nil {
			if len(me.SimpleContent.RestrictionSimpleContent.Base) > 0 {
				ctBaseType = me.SimpleContent.RestrictionSimpleContent.Base.String()
			}
			td.addAnnotations(me.SimpleContent.RestrictionSimpleContent.Annotation)
			for _, att = range me.SimpleContent.RestrictionSimpleContent.Attributes {
				allAtts[att] = true
			}
			for _, attGroup = range me.SimpleContent.RestrictionSimpleContent.AttributeGroups {
				allAttGroups[attGroup] = true
			}
			if (len(ctValueType) == 0) && (len(me.SimpleContent.RestrictionSimpleContent.Base) > 0) {
				ctValueType = me.SimpleContent.RestrictionSimpleContent.Base.String()
			}
			if (len(ctValueType) == 0) && (len(me.SimpleContent.RestrictionSimpleContent.SimpleTypes) > 0) {
				ctValueType = me.SimpleContent.RestrictionSimpleContent.SimpleTypes[0].Name.String()
			}
			for _, enum := range me.SimpleContent.RestrictionSimpleContent.Enumerations {
				println("ENUMTODO!?! Whoever sees this message, please post an issue at github.com/metaleap/go-xsd with a link to the XSD..." + enum.selfName().String())
			}
		}
	}
	if ctBaseType = bag.resolveQnameRef(ctBaseType, "T", nil); len(ctBaseType) > 0 {
		td.addEmbed(nil, bag.safeName(ctBaseType))
	} else if ctValueType = bag.resolveQnameRef(ctValueType, "T", nil); len(ctValueType) > 0 {
		bag.simpleContentValueTypes[typeSafeName] = ctValueType
		td.addField(nil, idPrefix+"Value", ctValueType, ",chardata")
		chain := sfmt("me.%vValue", idPrefix)
		td.addMethod(nil, "*"+typeSafeName, sfmt("To%v", bag.safeName(ctValueType)), ctValueType, sfmt("return %v", chain), sfmt("Simply returns the value of its %vValue field.", idPrefix))
		ttn := ctValueType
		for ttd := bag.declTypes[ctValueType]; ttd != nil; ttd = bag.declTypes[ttn] {
			if ttd != nil {
				bag.declConvs[ttd.Name] = true
			}
			if ttn = ttd.Type; len(ttn) > 0 {
				chain += sfmt(".To%v()", bag.safeName(ttn))
				td.addMethod(nil, "*"+typeSafeName, sfmt("To%v", bag.safeName(ttn)), ttn, sfmt("return %v", chain), sfmt("Returns the value of its %vValue field as a %v (which %v is just aliasing).", idPrefix, ttn, ctValueType))
			} else {
				break
			}
		}
		if (!strings.HasPrefix(ctValueType, "xsdt.")) && (bag.declTypes[ctValueType] == nil) {
			println("NOTFOUND: " + ctValueType)
		}
	} else if mixed {
		td.addEmbed(nil, idPrefix+"HasCdata")
	}
	for elGr, _ = range allElemGroups {
		subMakeElemGroup(bag, td, elGr, grsDone, anns(nil, me.ComplexContent)...)
	}
	for el, _ = range allElems {
		subMakeElem(bag, td, el, elsDone, 1, anns(me.All, nil)...)
	}
	for _, ch := range allChoices {
		for _, el = range ch.Elements {
			subMakeElem(bag, td, el, elsDone, ch.hasAttrMaxOccurs.Value(), ch.Annotation)
		}
		for _, elGr = range ch.Groups {
			subMakeElemGroup(bag, td, elGr, grsDone, ch.Annotation)
		}
	}
	for _, seq := range allSeqs {
		for _, el = range seq.Elements {
			subMakeElem(bag, td, el, elsDone, seq.hasAttrMaxOccurs.Value(), seq.Annotation)
		}
		for _, elGr = range seq.Groups {
			subMakeElemGroup(bag, td, elGr, grsDone, seq.Annotation)
		}
	}
	for attGroup, _ = range allAttGroups {
		td.addEmbed(attGroup, ustr.PrefixWithSep(bag.attGroupRefImps[attGroup], ".", bag.attGroups[attGroup][(strings.Index(bag.attGroups[attGroup], ".")+1):]), attGroup.Annotation)
	}

	for att, _ = range allAtts {
		if key := bag.attsKeys[att]; len(key) > 0 {
			td.addEmbed(att, ustr.PrefixWithSep(bag.attRefImps[att], ".", bag.attsCache[key][(strings.Index(bag.attsCache[key], ".")+1):]), att.Annotation)
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Documentation) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	if len(me.CDATA) > 0 {
		var s, ln string
		for _, ln = range ustr.Split(me.CDATA, "\n") {
			if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 {
				bag.appendFmt(false, "//\t%s", s)
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Element) makePkg(bag *PkgBag) {
	var (
		safeName, typeName, valueType, tmp, key, defVal, impName string
		subEl                                                    *Element
	)
	asterisk, defName, doc := "", "Default", ""
	me.elemBase.beforeMakePkg(bag)
	if len(me.Form) == 0 {
		me.Form = bag.Schema.ElementFormDefault
	}
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemComplexType.makePkg(bag)
	if len(me.Ref) > 0 {
		key = bag.resolveQnameRef(me.Ref.String(), "", &impName)
		for pref, cache := range map[string]map[string]string{"HasElem_": bag.elemsCacheOnce, "HasElems_": bag.elemsCacheMult} {
			tmp = ustr.PrefixWithSep(impName, ".", idPrefix+pref+bag.safeName(me.Ref.String()[(strings.Index(me.Ref.String(), ":")+1):]))
			if bag.elemRefImps[me], bag.elemKeys[me] = impName, key; len(cache[key]) == 0 {
				cache[key] = tmp
			}
		}
	} else {
		safeName = bag.safeName(me.Name.String())
		if typeName = me.Type.String(); (len(typeName) == 0) && ((me.ComplexType != nil) || (len(me.SimpleTypes) > 0)) {
			if me.ComplexType != nil {
				asterisk, typeName = "*", me.ComplexType.Name.String()
			} else {
				typeName = me.SimpleTypes[0].Name.String()
			}
		} else {
			if len(typeName) == 0 {
				typeName = bag.xsdStringTypeRef()
			}
			loadedSchemas := make(map[string]bool)
			if typeName = bag.resolveQnameRef(typeName, "T", &impName); bag.Schema.RootSchema([]string{bag.Schema.loadUri}).globalComplexType(bag, typeName, loadedSchemas) != nil {
				asterisk = "*"
			}
		}
		if defVal = me.Default; len(defVal) == 0 {
			defName, defVal = "Fixed", me.Fixed
		}
		if me.Parent() == bag.Schema {
			key = safeName
		} else {
			key = bag.safeName(bag.Stacks.FullName()) + "_" + safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal)
		}
		if valueType = bag.simpleContentValueTypes[typeName]; len(valueType) == 0 {
			valueType = typeName
		}
		isPt := bag.isParseType(valueType)
		if _, isChoice := me.Parent().(*Choice); isChoice && isPt {
			asterisk = "*"
		}
		for pref, cache := range map[string]map[string]string{"HasElem_": bag.elemsCacheOnce, "HasElems_": bag.elemsCacheMult} {
			if tmp = idPrefix + pref + key; !bag.elemsWritten[tmp] {
				bag.elemsWritten[tmp], bag.elemKeys[me] = true, key
				cache[key] = tmp
				var td = bag.addType(me, tmp, "", me.Annotation)
				td.addField(me, ustr.Ifs(pref == "HasElems_", pluralize(safeName), safeName), ustr.Ifs(pref == "HasElems_", "[]"+asterisk+typeName, asterisk+typeName), ustr.Ifs(len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String()+" ", "")+me.Name.String(), me.Annotation)
				if me.parent == bag.Schema {
					loadedSchemas := make(map[string]bool)
					for _, subEl = range bag.Schema.RootSchema([]string{bag.Schema.loadUri}).globalSubstitutionElems(me, loadedSchemas) {
						td.addEmbed(subEl, idPrefix+pref+bag.safeName(subEl.Name.String()), subEl.Annotation)
					}
				}
				if len(defVal) > 0 {
					doc = sfmt("Returns the %v value for %v -- "+ustr.Ifs(isPt, "%v", "%#v"), strings.ToLower(defName), safeName, defVal)
					if isPt {
						if PkgGen.ForceParseForDefaults {
							td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("var x = new(%v); x.Set(%#v); return *x", valueType, defVal), doc)
						} else {
							td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("return %v(%v)", valueType, defVal), doc)
						}
					} else {
						td.addMethod(nil, tmp, safeName+defName, valueType, sfmt("return %v(%#v)", valueType, defVal), doc)
					}
				}
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionComplexContent) makePkg(bag *PkgBag) {
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

func (me *ExtensionSimpleContent) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Field) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Group) makePkg(bag *PkgBag) {
	var refName, refImp string
	var choices = []*Choice{me.Choice}
	var seqs = []*Sequence{me.Sequence}
	var el *Element
	var gr *Group
	var elsDone, grsDone = map[string]bool{}, map[string]bool{}
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemSequence.makePkg(bag)
	if len(me.Ref) > 0 {
		if len(bag.elemGroups[me]) == 0 {
			refName = bag.resolveQnameRef(me.Ref.String(), "", &refImp)
			bag.elemGroups[me] = idPrefix + "HasGroup_" + refName
			bag.elemGroupRefImps[me] = refImp
		}
	} else {
		me.Ref.Set(me.Name.String())
		safeName := bag.safeName(me.Name.String())
		tmp := idPrefix + "HasGroup_" + safeName
		bag.elemGroups[me] = tmp
		var td = bag.addType(me, tmp, "", me.Annotation)
		choices, seqs = Flattened(choices, seqs)
		if me.All != nil {
			for _, el = range me.All.Elements {
				subMakeElem(bag, td, el, elsDone, 1, me.All.Annotation)
			}
		}
		for _, ch := range choices {
			for _, el = range ch.Elements {
				subMakeElem(bag, td, el, elsDone, ch.hasAttrMaxOccurs.Value(), ch.Annotation)
			}
			for _, gr = range ch.Groups {
				subMakeElemGroup(bag, td, gr, grsDone, ch.Annotation)
			}
		}
		for _, seq := range seqs {
			for _, el = range seq.Elements {
				subMakeElem(bag, td, el, elsDone, seq.hasAttrMaxOccurs.Value(), seq.Annotation)
			}
			for _, gr = range seq.Groups {
				subMakeElemGroup(bag, td, gr, grsDone, seq.Annotation)
			}
		}
	}

	me.elemBase.afterMakePkg(bag)
}

func (me *Import) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	var impName, impPath string
	var pos int
	me.hasElemAnnotation.makePkg(bag)
	for k, v := range bag.Schema.XMLNamespaces {
		if v == me.Namespace {
			impName = safeIdentifier(k)
			break
		}
	}
	if len(impName) > 0 {
		if pos, impPath = strings.Index(me.SchemaLocation.String(), protSep), me.SchemaLocation.String(); pos > 0 {
			impPath = impPath[pos+len(protSep):]
		} else {
			impPath = path.Join(path.Dir(bag.Schema.loadUri), impPath)
		}
		impPath = path.Join(path.Dir(impPath), goPkgPrefix+path.Base(impPath)+goPkgSuffix)
		bag.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Key) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *KeyRef) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *List) makePkg(bag *PkgBag) {
	var safeName string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	rtr := bag.resolveQnameRef(me.ItemType.String(), "T", nil)
	if len(rtr) == 0 {
		rtr = me.SimpleTypes[0].Name.String()
	}
	st := bag.Stacks.CurSimpleType()
	safeName = bag.safeName(ustr.PrependIf(st.Name.String(), "T"))
	body, doc := "", sfmt("%v declares a String containing a whitespace-separated list of %v values. This Values() method creates and returns a slice of all elements in that list", safeName, rtr)
	body = sfmt("svals := %v.ListValues(string(me)); list = make([]%v, len(svals)); for i, s := range svals { list[i].Set(s) }; return", bag.impName, rtr)
	bag.ctd.addMethod(me, safeName, "Values", sfmt("(list []%v)", rtr), body, doc+".", me.Annotation)
	for baseType := bag.simpleBaseTypes[rtr]; len(baseType) > 0; baseType = bag.simpleBaseTypes[baseType] {
		body = sfmt("svals := %v.ListValues(string(me)); list = make([]%v, len(svals)); for i, s := range svals { list[i].Set(s) }; return", bag.impName, baseType)
		bag.ctd.addMethod(me, safeName, "Values"+bag.safeName(baseType), sfmt("(list []%v)", baseType), body, sfmt("%s, typed as %s.", doc, baseType), me.Annotation)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Notation) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(false, "%vNotations.Add(%#v, %#v, %#v, %#v)", idPrefix, me.Id, me.Name, me.Public, me.System)
	me.elemBase.afterMakePkg(bag)
}

func (me *Redefine) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionComplexContent) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleContent) makePkg(bag *PkgBag) {
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

func (me *RestrictionSimpleEnumeration) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	safeName := bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T"))
	var doc = sfmt("Returns true if the value of this enumerated %v is %#v.", safeName, me.Value)
	bag.ctd.addMethod(me, safeName, "Is"+bag.safeName(me.Value), "bool", sfmt("return me.String() == %#v", me.Value), doc)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleFractionDigits) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleLength) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxExclusive) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxInclusive) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxLength) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinExclusive) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinInclusive) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinLength) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimplePattern) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleTotalDigits) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleType) makePkg(bag *PkgBag) {
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
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleWhiteSpace) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Schema) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsImport.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsRedefine.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Selector) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Sequence) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleContent) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemExtensionSimpleContent.makePkg(bag)
	me.hasElemRestrictionSimpleContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleType) makePkg(bag *PkgBag) {
	var typeName = me.Name
	var baseType, safeName = "", ""
	var resolve = true
	var isPt bool
	if len(typeName) == 0 {
		typeName = bag.AnonName(me.longSafeName(bag))
		me.Name = typeName
	} else {
		me.Name = typeName
	}
	typeName = xsdt.NCName(ustr.PrependIf(typeName.String(), "T"))
	me.elemBase.beforeMakePkg(bag)
	bag.Stacks.SimpleType.Push(me)
	safeName = bag.safeName(typeName.String())
	if me.RestrictionSimpleType != nil {
		if baseType = me.RestrictionSimpleType.Base.String(); (len(baseType) == 0) && (len(me.RestrictionSimpleType.SimpleTypes) > 0) {
			resolve, baseType = false, me.RestrictionSimpleType.SimpleTypes[0].Name.String()
		}
	}
	if len(baseType) == 0 {
		baseType = bag.xsdStringTypeRef()
	}
	if resolve {
		baseType = bag.resolveQnameRef(baseType, "T", nil)
	}
	bag.simpleBaseTypes[safeName] = baseType
	if isPt = bag.isParseType(baseType); isPt {
		bag.parseTypes[safeName] = true
	}
	var td = bag.addType(me, safeName, baseType, me.Annotation)
	var doc string
	if isPt {
		doc = sfmt("Since %v is a non-string scalar type (either boolean or numeric), sets the current value obtained from parsing the specified string.", safeName)
	} else {
		doc = sfmt("Since %v is just a simple String type, this merely sets the current value from the specified string.", safeName)
	}
	td.addMethod(nil, "*"+safeName, "Set (s string)", "", sfmt("(*%v)(me).Set(s)", baseType), doc)
	if isPt {
		doc = sfmt("Returns a string representation of this %v's current non-string scalar value.", safeName)
	} else {
		doc = sfmt("Since %v is just a simple String type, this merely returns the current string value.", safeName)
	}
	td.addMethod(nil, safeName, "String", "string", sfmt("return %v(me).String()", baseType), doc)
	doc = sfmt("This convenience method just performs a simple type conversion to %v's alias type %v.", safeName, baseType)
	td.addMethod(nil, safeName, "To"+bag.safeName(baseType), baseType, sfmt("return %v(me)", baseType), doc)
	me.hasElemRestrictionSimpleType.makePkg(bag)
	me.hasElemList.makePkg(bag)
	me.hasElemUnion.makePkg(bag)
	bag.Stacks.SimpleType.Pop()
	me.elemBase.afterMakePkg(bag)
}

func (me *Union) makePkg(bag *PkgBag) {
	var memberTypes []string
	var rtn, rtnSafeName, safeName string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(me.MemberTypes, " ")
	for _, st := range me.SimpleTypes {
		memberTypes = append(memberTypes, st.Name.String())
	}
	for _, mt := range memberTypes {
		rtn = bag.resolveQnameRef(mt, "T", nil)
		safeName, rtnSafeName = bag.safeName(ustr.PrependIf(bag.Stacks.CurSimpleType().Name.String(), "T")), bag.safeName(rtn)
		bag.ctd.addMethod(me, safeName, "To"+rtnSafeName, rtn, sfmt(ustr.Ifs(bag.isParseType(rtn), "var x = new(%v); x.Set(me.String()); return *x", "return %v(me)"), rtn), sfmt("%v is an XSD union-type of several types. This is a simple type conversion to %v, but keep in mind the actual value may or may not be a valid %v value.", safeName, rtnSafeName, rtnSafeName), me.Annotation)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Unique) makePkg(bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func anns(a *All, cc *ComplexContent) (anns []*Annotation) {
	if (a != nil) && (a.Annotation != nil) {
		anns = append(anns, a.Annotation)
	}
	if cc != nil {
		if cc.Annotation != nil {
			anns = append(anns, cc.Annotation)
		}
		if ecc := cc.ExtensionComplexContent; (ecc != nil) && (ecc.Annotation != nil) {
			anns = append(anns, ecc.Annotation)
		}
	}
	return
}

func pluralize(s string) string {
	for _, psp := range PkgGen.PluralizeSpecialPrefixes {
		if strings.HasPrefix(s, psp) {
			return ustr.Pluralize(s[len(psp):] + s[:len(psp)])
		}
	}
	return ustr.Pluralize(s)
}

func sfmt(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

// For any rune, return a rune that is a valid in an identifier
func coerceToIdentifierRune(ch rune) rune {
	if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
		return '_'
	}
	return ch
}

// Take any string and convert it to a valid identifier
// Appends an underscore if the first rune is a number
func safeIdentifier(s string) string {
	s = strings.Map(coerceToIdentifierRune, s)
	if unicode.IsNumber([]rune(s)[0]) {
		s = fmt.Sprint("_", s)
	}
	return s
}

func subMakeElem(bag *PkgBag, td *declType, el *Element, done map[string]bool, parentMaxOccurs xsdt.Long, anns ...*Annotation) {
	var elCache map[string]string
	anns = append(anns, el.Annotation)
	if refName := bag.elemKeys[el]; (len(refName) > 0) && (!done[refName]) {
		if done[refName], elCache = true, ustr.Ifm((parentMaxOccurs == 1) && (el.hasAttrMaxOccurs.Value() == 1), bag.elemsCacheOnce, bag.elemsCacheMult); !strings.HasPrefix(elCache[refName], bag.impName+"."+idPrefix) {
			td.addEmbed(el, elCache[refName], anns...)
		}
	}
}

func subMakeElemGroup(bag *PkgBag, td *declType, gr *Group, done map[string]bool, anns ...*Annotation) {
	var refImp string
	anns = append(anns, gr.Annotation)
	if refName := bag.resolveQnameRef(gr.Ref.String(), "", &refImp); !done[refName] {
		if done[refName] = true; len(refImp) > 0 {
			if !strings.HasPrefix(refName, bag.impName+"."+idPrefix) {
				td.addEmbed(gr, refImp+"."+idPrefix+"HasGroup_"+refName[(len(refImp)+1):], anns...)
			}
		} else {
			td.addEmbed(gr, idPrefix+"HasGroup_"+refName, anns...)
		}
	}
}
