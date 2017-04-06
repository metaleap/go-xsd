package xsd

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type hasAttrAbstract struct {
	Abstract bool `xml:"abstract,attr"`
}

type hasAttrAttributeFormDefault struct {
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`
}

type hasAttrBase struct {
	Base xsdt.Qname `xml:"base,attr"`
}

type hasAttrBlock struct {
	Block string `xml:"block,attr"`
}

type hasAttrBlockDefault struct {
	BlockDefault string `xml:"blockDefault,attr"`
}

type hasAttrDefault struct {
	Default string `xml:"default,attr"`
}

type hasAttrFinal struct {
	Final string `xml:"final,attr"`
}

type hasAttrFinalDefault struct {
	FinalDefault string `xml:"finalDefault,attr"`
}

type hasAttrFixed struct {
	Fixed string `xml:"fixed,attr"`
}

type hasAttrForm struct {
	Form string `xml:"form,attr"`
}

type hasAttrElementFormDefault struct {
	ElementFormDefault string `xml:"elementFormDefault,attr"`
}

type hasAttrId struct {
	Id xsdt.Id `xml:"id,attr"`
}

type hasAttrItemType struct {
	ItemType xsdt.Qname `xml:"itemType,attr"`
}

type hasAttrLang struct {
	Lang xsdt.Language `xml:"lang,attr"`
}

type hasAttrMaxOccurs struct {
	MaxOccurs string `xml:"maxOccurs,attr"`
}

func (me *hasAttrMaxOccurs) Value() (l xsdt.Long) {
	if len(me.MaxOccurs) == 0 {
		l = 1
	} else if me.MaxOccurs == "unbounded" {
		l = -1
	} else {
		l.Set(me.MaxOccurs)
	}
	return
}

type hasAttrMemberTypes struct {
	MemberTypes string `xml:"memberTypes,attr"`
}

type hasAttrMinOccurs struct {
	MinOccurs uint64 `xml:"minOccurs,attr"`
}

type hasAttrMixed struct {
	Mixed bool `xml:"mixed,attr"`
}

type hasAttrName struct {
	Name xsdt.NCName `xml:"name,attr"`
}

type hasAttrNamespace struct {
	Namespace string `xml:"namespace,attr"`
}

type hasAttrNillable struct {
	Nillable bool `xml:"nillable,attr"`
}

type hasAttrProcessContents struct {
	ProcessContents string `xml:"processContents,attr"`
}

type hasAttrPublic struct {
	Public string `xml:"public,attr"`
}

type hasAttrRef struct {
	Ref xsdt.Qname `xml:"ref,attr"`
}

type hasAttrRefer struct {
	Refer xsdt.Qname `xml:"refer,attr"`
}

type hasAttrSchemaLocation struct {
	SchemaLocation xsdt.AnyURI `xml:"schemaLocation,attr"`
}

type hasAttrSource struct {
	Source xsdt.AnyURI `xml:"source,attr"`
}

type hasAttrSubstitutionGroup struct {
	SubstitutionGroup xsdt.Qname `xml:"substitutionGroup,attr"`
}

type hasAttrSystem struct {
	System xsdt.AnyURI `xml:"system,attr"`
}

type hasAttrTargetNamespace struct {
	TargetNamespace xsdt.AnyURI `xml:"targetNamespace,attr"`
}

type hasAttrType struct {
	Type xsdt.Qname `xml:"type,attr"`
}

type hasAttrUse struct {
	Use string `xml:"use,attr"`
}

type hasAttrValue struct {
	Value string `xml:"value,attr"`
}

type hasAttrVersion struct {
	Version xsdt.Token `xml:"version,attr"`
}

type hasAttrXpath struct {
	Xpath string `xml:"xpath,attr"`
}
