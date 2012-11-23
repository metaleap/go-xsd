package xsd

func (me *All) initElement(parent element) {
	me.elemBase.init(parent, me, "all", &me.hasAttrId, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsElement.initChildren(me)
}

func (me *Annotation) initElement(parent element) {
	me.elemBase.init(parent, me, "annotation")
	me.hasElemsAppInfo.initChildren(me)
	me.hasElemsDocumentation.initChildren(me)
}

func (me *Any) initElement(parent element) {
	me.elemBase.init(parent, me, "any", &me.hasAttrId, &me.hasAttrNamespace, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs, &me.hasAttrProcessContents)
	me.hasElemAnnotation.initChildren(me)
}

func (me *AnyAttribute) initElement(parent element) {
	me.elemBase.init(parent, me, "anyAttribute", &me.hasAttrId, &me.hasAttrNamespace, &me.hasAttrProcessContents)
	me.hasElemAnnotation.initChildren(me)
}

func (me *AppInfo) initElement(parent element) {
	me.elemBase.init(parent, me, "appInfo", &me.hasAttrSource)
}

func (me *Attribute) initElement(parent element) {
	me.elemBase.init(parent, me, "attribute", &me.hasAttrDefault, &me.hasAttrFixed, &me.hasAttrForm, &me.hasAttrId, &me.hasAttrName, &me.hasAttrRef, &me.hasAttrType, &me.hasAttrUse)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *AttributeGroup) initElement(parent element) {
	me.elemBase.init(parent, me, "attributeGroup", &me.hasAttrId, &me.hasAttrName, &me.hasAttrRef)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
}

func (me *Choice) initElement(parent element) {
	me.elemBase.init(parent, me, "choice", &me.hasAttrId, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsAny.initChildren(me)
	me.hasElemsChoice.initChildren(me)
	me.hasElemsElement.initChildren(me)
	me.hasElemsGroup.initChildren(me)
	me.hasElemsSequence.initChildren(me)
}

func (me *ComplexContent) initElement(parent element) {
	me.elemBase.init(parent, me, "complexContent", &me.hasAttrId, &me.hasAttrMixed)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemExtensionComplexContent.initChildren(me)
	me.hasElemRestrictionComplexContent.initChildren(me)
}

func (me *ComplexType) initElement(parent element) {
	me.elemBase.init(parent, me, "complexType", &me.hasAttrAbstract, &me.hasAttrBlock, &me.hasAttrFinal, &me.hasAttrId, &me.hasAttrMixed, &me.hasAttrName)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemAll.initChildren(me)
	me.hasElemChoice.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemGroup.initChildren(me)
	me.hasElemSequence.initChildren(me)
	me.hasElemComplexContent.initChildren(me)
	me.hasElemSimpleContent.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
}

func (me *Documentation) initElement(parent element) {
	me.elemBase.init(parent, me, "documentation", &me.hasAttrLang, &me.hasAttrSource)
}

func (me *Element) initElement(parent element) {
	me.elemBase.init(parent, me, "element", &me.hasAttrAbstract, &me.hasAttrBlock, &me.hasAttrDefault, &me.hasAttrFinal, &me.hasAttrFixed, &me.hasAttrForm, &me.hasAttrId, &me.hasAttrName, &me.hasAttrNillable, &me.hasAttrRef, &me.hasAttrType, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs, &me.hasAttrSubstitutionGroup)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemUnique.initChildren(me)
	me.hasElemsKey.initChildren(me)
	me.hasElemComplexType.initChildren(me)
	me.hasElemKeyRef.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *ExtensionComplexContent) initElement(parent element) {
	me.elemBase.init(parent, me, "extension", &me.hasAttrBase, &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemAll.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsChoice.initChildren(me)
	me.hasElemsGroup.initChildren(me)
	me.hasElemsSequence.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
}

func (me *ExtensionSimpleContent) initElement(parent element) {
	me.elemBase.init(parent, me, "extension", &me.hasAttrBase, &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
}

func (me *Field) initElement(parent element) {
	me.elemBase.init(parent, me, "field", &me.hasAttrId, &me.hasAttrXpath)
	me.hasElemAnnotation.initChildren(me)
}

func (me *Group) initElement(parent element) {
	me.elemBase.init(parent, me, "group", &me.hasAttrId, &me.hasAttrName, &me.hasAttrRef, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemAll.initChildren(me)
	me.hasElemChoice.initChildren(me)
	me.hasElemSequence.initChildren(me)
}

func (me *Import) initElement(parent element) {
	me.elemBase.init(parent, me, "import", &me.hasAttrId, &me.hasAttrNamespace, &me.hasAttrSchemaLocation)
	me.hasElemAnnotation.initChildren(me)
}

func (me *Key) initElement(parent element) {
	me.elemBase.init(parent, me, "key", &me.hasAttrId, &me.hasAttrName)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemField.initChildren(me)
	me.hasElemSelector.initChildren(me)
}

func (me *KeyRef) initElement(parent element) {
	me.elemBase.init(parent, me, "keyref", &me.hasAttrId, &me.hasAttrName, &me.hasAttrRefer)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemField.initChildren(me)
	me.hasElemSelector.initChildren(me)
}

func (me *List) initElement(parent element) {
	me.elemBase.init(parent, me, "list", &me.hasAttrId, &me.hasAttrItemType)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *Notation) initElement(parent element) {
	me.elemBase.init(parent, me, "notation", &me.hasAttrId, &me.hasAttrName, &me.hasAttrPublic, &me.hasAttrSystem)
	me.hasElemAnnotation.initChildren(me)
}

func (me *Redefine) initElement(parent element) {
	me.elemBase.init(parent, me, "redefine", &me.hasAttrId, &me.hasAttrSchemaLocation)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsGroup.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
	me.hasElemsComplexType.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *RestrictionComplexContent) initElement(parent element) {
	me.elemBase.init(parent, me, "restriction", &me.hasAttrBase, &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemAll.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsChoice.initChildren(me)
	me.hasElemsSequence.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
}

func (me *RestrictionSimpleContent) initElement(parent element) {
	me.elemBase.init(parent, me, "restriction", &me.hasAttrBase, &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemLength.initChildren(me)
	me.hasElemPattern.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsEnumeration.initChildren(me)
	me.hasElemFractionDigits.initChildren(me)
	me.hasElemMaxExclusive.initChildren(me)
	me.hasElemMaxInclusive.initChildren(me)
	me.hasElemMaxLength.initChildren(me)
	me.hasElemMinExclusive.initChildren(me)
	me.hasElemMinInclusive.initChildren(me)
	me.hasElemMinLength.initChildren(me)
	me.hasElemTotalDigits.initChildren(me)
	me.hasElemWhiteSpace.initChildren(me)
	me.hasElemsAnyAttribute.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *RestrictionSimpleEnumeration) initElement(parent element) {
	me.elemBase.init(parent, me, "enumeration", &me.hasAttrValue)
}

func (me *RestrictionSimpleFractionDigits) initElement(parent element) {
	me.elemBase.init(parent, me, "fractionDigits", &me.hasAttrValue)
}

func (me *RestrictionSimpleLength) initElement(parent element) {
	me.elemBase.init(parent, me, "length", &me.hasAttrValue)
}

func (me *RestrictionSimpleMaxExclusive) initElement(parent element) {
	me.elemBase.init(parent, me, "maxExclusive", &me.hasAttrValue)
}

func (me *RestrictionSimpleMaxInclusive) initElement(parent element) {
	me.elemBase.init(parent, me, "maxInclusive", &me.hasAttrValue)
}

func (me *RestrictionSimpleMaxLength) initElement(parent element) {
	me.elemBase.init(parent, me, "maxLength", &me.hasAttrValue)
}

func (me *RestrictionSimpleMinExclusive) initElement(parent element) {
	me.elemBase.init(parent, me, "minExclusive", &me.hasAttrValue)
}

func (me *RestrictionSimpleMinInclusive) initElement(parent element) {
	me.elemBase.init(parent, me, "minInclusive", &me.hasAttrValue)
}

func (me *RestrictionSimpleMinLength) initElement(parent element) {
	me.elemBase.init(parent, me, "minLength", &me.hasAttrValue)
}

func (me *RestrictionSimplePattern) initElement(parent element) {
	me.elemBase.init(parent, me, "pattern", &me.hasAttrValue)
}

func (me *RestrictionSimpleTotalDigits) initElement(parent element) {
	me.elemBase.init(parent, me, "totalDigits", &me.hasAttrValue)
}

func (me *RestrictionSimpleType) initElement(parent element) {
	me.elemBase.init(parent, me, "restriction", &me.hasAttrBase, &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemLength.initChildren(me)
	me.hasElemPattern.initChildren(me)
	me.hasElemsEnumeration.initChildren(me)
	me.hasElemFractionDigits.initChildren(me)
	me.hasElemMaxExclusive.initChildren(me)
	me.hasElemMaxInclusive.initChildren(me)
	me.hasElemMaxLength.initChildren(me)
	me.hasElemMinExclusive.initChildren(me)
	me.hasElemMinInclusive.initChildren(me)
	me.hasElemMinLength.initChildren(me)
	me.hasElemTotalDigits.initChildren(me)
	me.hasElemWhiteSpace.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *RestrictionSimpleWhiteSpace) initElement(parent element) {
	me.elemBase.init(parent, me, "whiteSpace", &me.hasAttrValue)
}

func (me *Schema) initElement(parent element) {
	me.elemBase.init(parent, me, "schema", &me.hasAttrId, &me.hasAttrLang, &me.hasAttrVersion, &me.hasAttrBlockDefault, &me.hasAttrFinalDefault, &me.hasAttrSchemaLocation, &me.hasAttrTargetNamespace, &me.hasAttrAttributeFormDefault, &me.hasAttrElementFormDefault)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsAttribute.initChildren(me)
	me.hasElemsElement.initChildren(me)
	me.hasElemsGroup.initChildren(me)
	me.hasElemsImport.initChildren(me)
	me.hasElemsNotation.initChildren(me)
	me.hasElemsRedefine.initChildren(me)
	me.hasElemsAttributeGroup.initChildren(me)
	me.hasElemsComplexType.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *Selector) initElement(parent element) {
	me.elemBase.init(parent, me, "selector", &me.hasAttrId, &me.hasAttrXpath)
	me.hasElemAnnotation.initChildren(me)
}

func (me *Sequence) initElement(parent element) {
	me.elemBase.init(parent, me, "sequence", &me.hasAttrId, &me.hasAttrMaxOccurs, &me.hasAttrMinOccurs)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsAny.initChildren(me)
	me.hasElemsChoice.initChildren(me)
	me.hasElemsElement.initChildren(me)
	me.hasElemsGroup.initChildren(me)
	me.hasElemsSequence.initChildren(me)
}

func (me *SimpleContent) initElement(parent element) {
	me.elemBase.init(parent, me, "simpleContent", &me.hasAttrId)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemExtensionSimpleContent.initChildren(me)
	me.hasElemRestrictionSimpleContent.initChildren(me)
}

func (me *SimpleType) initElement(parent element) {
	me.elemBase.init(parent, me, "simpleType", &me.hasAttrFinal, &me.hasAttrId, &me.hasAttrName)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemRestrictionSimpleType.initChildren(me)
	me.hasElemList.initChildren(me)
	me.hasElemUnion.initChildren(me)
}

func (me *Union) initElement(parent element) {
	me.elemBase.init(parent, me, "union", &me.hasAttrId, &me.hasAttrMemberTypes)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemsSimpleType.initChildren(me)
}

func (me *Unique) initElement(parent element) {
	me.elemBase.init(parent, me, "unique", &me.hasAttrId, &me.hasAttrName)
	me.hasElemAnnotation.initChildren(me)
	me.hasElemField.initChildren(me)
	me.hasElemSelector.initChildren(me)
}
