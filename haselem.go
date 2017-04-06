package xsd

type hasCdata struct {
	CDATA string `xml:",chardata"`
}

type hasElemAll struct {
	All *All `xml:"all"`
}

type hasElemAnnotation struct {
	Annotation *Annotation `xml:"annotation"`
}

type hasElemsAny struct {
	Anys []*Any `xml:"any"`
}

type hasElemsAnyAttribute struct {
	AnyAttributes []*AnyAttribute `xml:"anyAttribute"`
}

type hasElemsAppInfo struct {
	AppInfos []*AppInfo `xml:"appinfo"`
}

type hasElemsAttribute struct {
	Attributes []*Attribute `xml:"attribute"`
}

type hasElemsAttributeGroup struct {
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}

type hasElemChoice struct {
	Choice *Choice `xml:"choice"`
}

type hasElemsChoice struct {
	Choices []*Choice `xml:"choice"`
}

type hasElemComplexContent struct {
	ComplexContent *ComplexContent `xml:"complexContent"`
}

type hasElemComplexType struct {
	ComplexType *ComplexType `xml:"complexType"`
}

type hasElemsComplexType struct {
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type hasElemsDocumentation struct {
	Documentations []*Documentation `xml:"documentation"`
}

type hasElemsElement struct {
	Elements []*Element `xml:"element"`
}

type hasElemsEnumeration struct {
	Enumerations []*RestrictionSimpleEnumeration `xml:"enumeration"`
}

type hasElemExtensionComplexContent struct {
	ExtensionComplexContent *ExtensionComplexContent `xml:"extension"`
}

type hasElemExtensionSimpleContent struct {
	ExtensionSimpleContent *ExtensionSimpleContent `xml:"extension"`
}

type hasElemField struct {
	Field *Field `xml:"field"`
}

type hasElemFractionDigits struct {
	FractionDigits *RestrictionSimpleFractionDigits `xml:"fractionDigits"`
}

type hasElemGroup struct {
	Group *Group `xml:"group"`
}

type hasElemsGroup struct {
	Groups []*Group `xml:"group"`
}

type hasElemsImport struct {
	Imports []*Import `xml:"import"`
}

type hasElemsInclude struct {
	Includes []*Include `xml:"include"`
}

type hasElemsKey struct {
	Keys []*Key `xml:"key"`
}

type hasElemKeyRef struct {
	KeyRef *KeyRef `xml:"keyref"`
}

type hasElemLength struct {
	Length *RestrictionSimpleLength `xml:"length"`
}

type hasElemList struct {
	List *List `xml:"list"`
}

type hasElemMaxExclusive struct {
	MaxExclusive *RestrictionSimpleMaxExclusive `xml:"maxExclusive"`
}

type hasElemMaxInclusive struct {
	MaxInclusive *RestrictionSimpleMaxInclusive `xml:"maxInclusive"`
}

type hasElemMaxLength struct {
	MaxLength *RestrictionSimpleMaxLength `xml:"maxLength"`
}

type hasElemMinExclusive struct {
	MinExclusive *RestrictionSimpleMinExclusive `xml:"minExclusive"`
}

type hasElemMinInclusive struct {
	MinInclusive *RestrictionSimpleMinInclusive `xml:"minInclusive"`
}

type hasElemMinLength struct {
	MinLength *RestrictionSimpleMinLength `xml:"minLength"`
}

type hasElemsNotation struct {
	Notations []*Notation `xml:"notation"`
}

type hasElemPattern struct {
	Pattern *RestrictionSimplePattern `xml:"pattern"`
}

type hasElemsRedefine struct {
	Redefines []*Redefine `xml:"redefine"`
}

type hasElemRestrictionComplexContent struct {
	RestrictionComplexContent *RestrictionComplexContent `xml:"restriction"`
}

type hasElemRestrictionSimpleContent struct {
	RestrictionSimpleContent *RestrictionSimpleContent `xml:"restriction"`
}

type hasElemRestrictionSimpleType struct {
	RestrictionSimpleType *RestrictionSimpleType `xml:"restriction"`
}

type hasElemSelector struct {
	Selector *Selector `xml:"selector"`
}

type hasElemSequence struct {
	Sequence *Sequence `xml:"sequence"`
}

type hasElemsSequence struct {
	Sequences []*Sequence `xml:"sequence"`
}

type hasElemSimpleContent struct {
	SimpleContent *SimpleContent `xml:"simpleContent"`
}

type hasElemsSimpleType struct {
	SimpleTypes []*SimpleType `xml:"simpleType"`
}

type hasElemTotalDigits struct {
	TotalDigits *RestrictionSimpleTotalDigits `xml:"totalDigits"`
}

type hasElemUnion struct {
	Union *Union `xml:"union"`
}

type hasElemUnique struct {
	Unique *Unique `xml:"unique"`
}

type hasElemWhiteSpace struct {
	WhiteSpace *RestrictionSimpleWhiteSpace `xml:"whiteSpace"`
}
