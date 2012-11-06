package xsd

import (
	"encoding/xml"
)

type All struct {
	XMLName xml.Name `xml:"all"`
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsElement
}

type Annotation struct {
	XMLName xml.Name `xml:"annotation"`
	hasElemsAppInfo
	hasElemsDocumentation
}

type Any struct {
	XMLName xml.Name `xml:"any"`
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrNamespace
	hasAttrProcessContents
	hasElemAnnotation
}

type AnyAttribute struct {
	XMLName xml.Name `xml:"anyAttribute"`
	hasAttrId
	hasAttrNamespace
	hasAttrProcessContents
	hasElemAnnotation
}

type AppInfo struct {
	XMLName xml.Name `xml:"appinfo"`
	hasAttrSource
	hasCdata
}

type Attribute struct {
	XMLName xml.Name `xml:"attribute"`
	hasAttrDefault
	hasAttrFixed
	hasAttrForm
	hasAttrId
	hasAttrName
	hasAttrRef
	hasAttrType
	hasAttrUse
	hasElemAnnotation
	hasElemsSimpleType
}

type AttributeGroup struct {
	XMLName xml.Name `xml:"attributeGroup"`
	hasAttrId
	hasAttrName
	hasAttrRef
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
}

type Choice struct {
	XMLName xml.Name `xml:"choice"`
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsAny
	hasElemsChoice
	hasElemsElement
	hasElemsGroup
	hasElemsSequence
}

type ComplexContent struct {
	XMLName xml.Name `xml:"complexContent"`
	hasAttrId
	hasAttrMixed
	hasElemAnnotation
	hasElemExtensionComplexContent
	hasElemRestrictionComplexContent
}

type ComplexType struct {
	XMLName xml.Name `xml:"complexType"`
	hasAttrAbstract
	hasAttrBlock
	hasAttrFinal
	hasAttrId
	hasAttrMixed
	hasAttrName
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemChoice
	hasElemComplexContent
	hasElemsGroup
	hasElemsSequence
	hasElemSimpleContent
}

type Documentation struct {
	XMLName xml.Name `xml:"documentation"`
	hasAttrLang
	hasAttrSource
	hasCdata
}

type Element struct {
	XMLName xml.Name `xml:"element"`
	hasAttrAbstract
	hasAttrBlock
	hasAttrDefault
	hasAttrFinal
	hasAttrFixed
	hasAttrForm
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrNillable
	hasAttrRef
	hasAttrSubstitutionGroup
	hasAttrType
	hasElemAnnotation
	hasElemComplexType
	hasElemsKey
	hasElemKeyRef
	hasElemsSimpleType
	hasElemUnique
}

type ExtensionComplexContent struct {
	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrId
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsGroup
	hasElemsSequence
}

type ExtensionSimpleContent struct {
	XMLName xml.Name `xml:"extension"`
	hasAttrBase
	hasAttrId
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	hasAttrId
	hasAttrXpath
	hasElemAnnotation
}

type Group struct {
	XMLName xml.Name `xml:"group"`
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasAttrName
	hasAttrRef
	hasElemAll
	hasElemAnnotation
	hasElemChoice
	hasElemSequence
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	hasAttrId
	hasAttrSchemaLocation
	hasElemAnnotation
}

type Import struct {
	XMLName xml.Name `xml:"import"`
	hasAttrId
	hasAttrNamespace
	hasAttrSchemaLocation
	hasElemAnnotation
}

type Key struct {
	XMLName xml.Name `xml:"key"`
	hasAttrId
	hasAttrName
	hasElemAnnotation
	hasElemField
	hasElemSelector
}

type KeyRef struct {
	XMLName xml.Name `xml:"keyref"`
	hasAttrId
	hasAttrName
	hasAttrRefer
	hasElemAnnotation
	hasElemField
	hasElemSelector
}

type List struct {
	XMLName xml.Name `xml:"list"`
	hasAttrId
	hasAttrItemType
	hasElemAnnotation
	hasElemsSimpleType
}

type Notation struct {
	XMLName xml.Name `xml:"notation"`
	hasAttrId
	hasAttrName
	hasAttrPublic
	hasAttrSystem
	hasElemAnnotation
}

type Redefine struct {
	XMLName xml.Name `xml:"redefine"`
	hasAttrId
	hasAttrSchemaLocation
	hasElemAnnotation
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsGroup
	hasElemsSimpleType
}

type RestrictionComplexContent struct {
	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrId
	hasElemAll
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsChoice
	hasElemsSequence
}

type RestrictionSimpleContent struct {
	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrId
	hasElemAnnotation
	hasElemsAnyAttribute
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemPattern
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemWhiteSpace
}

type RestrictionSimpleEnumeration struct {
	XMLName xml.Name `xml:"enumeration"`
	hasAttrValue
}

type RestrictionSimpleFractionDigits struct {
	XMLName xml.Name `xml:"fractionDigits"`
	hasAttrValue
}

type RestrictionSimpleLength struct {
	XMLName xml.Name `xml:"length"`
	hasAttrValue
}

type RestrictionSimpleMaxExclusive struct {
	XMLName xml.Name `xml:"maxExclusive"`
	hasAttrValue
}

type RestrictionSimpleMaxInclusive struct {
	XMLName xml.Name `xml:"maxInclusive"`
	hasAttrValue
}

type RestrictionSimpleMaxLength struct {
	XMLName xml.Name `xml:"maxLength"`
	hasAttrValue
}

type RestrictionSimpleMinExclusive struct {
	XMLName xml.Name `xml:"minExclusive"`
	hasAttrValue
}

type RestrictionSimpleMinInclusive struct {
	XMLName xml.Name `xml:"minInclusive"`
	hasAttrValue
}

type RestrictionSimpleMinLength struct {
	XMLName xml.Name `xml:"minLength"`
	hasAttrValue
}

type RestrictionSimplePattern struct {
	XMLName xml.Name `xml:"pattern"`
	hasAttrValue
}

type RestrictionSimpleTotalDigits struct {
	XMLName xml.Name `xml:"totalDigits"`
	hasAttrValue
}

type RestrictionSimpleType struct {
	XMLName xml.Name `xml:"restriction"`
	hasAttrBase
	hasAttrId
	hasElemAnnotation
	hasElemsEnumeration
	hasElemFractionDigits
	hasElemLength
	hasElemMaxExclusive
	hasElemMaxInclusive
	hasElemMaxLength
	hasElemMinExclusive
	hasElemMinInclusive
	hasElemMinLength
	hasElemPattern
	hasElemsSimpleType
	hasElemTotalDigits
	hasElemWhiteSpace
}

type RestrictionSimpleWhiteSpace struct {
	XMLName xml.Name `xml:"whiteSpace"`
	hasAttrValue
}

type Selector struct {
	XMLName xml.Name `xml:"selector"`
	hasAttrId
	hasAttrXpath
	hasElemAnnotation
}

type Sequence struct {
	XMLName xml.Name `xml:"sequence"`
	hasAttrId
	hasAttrMaxOccurs
	hasAttrMinOccurs
	hasElemAnnotation
	hasElemsAny
	hasElemsChoice
	hasElemsElement
	hasElemsGroup
	hasElemsSequence
}

type SimpleContent struct {
	XMLName xml.Name `xml:"simpleContent"`
	hasAttrId
	hasElemAnnotation
	hasElemExtensionSimpleContent
	hasElemRestrictionSimpleContent
}

type SimpleType struct {
	XMLName xml.Name `xml:"simpleType"`
	hasAttrFinal
	hasAttrId
	hasAttrName
	hasElemAnnotation
	hasElemList
	hasElemRestrictionSimpleType
	hasElemUnion
}

type Union struct {
	XMLName xml.Name `xml:"union"`
	hasAttrId
	hasAttrMemberTypes
	hasElemAnnotation
	hasElemsSimpleType
}

type Unique struct {
	XMLName xml.Name `xml:"unique"`
	hasAttrId
	hasAttrName
	hasElemAnnotation
	hasElemField
	hasElemSelector
}
