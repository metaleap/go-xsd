package xsd

func (me *hasElemAll) initChildren(p element) {
	if me.All != nil {
		me.All.initElement(p)
	}
}

func (me *hasElemAnnotation) initChildren(p element) {
	if me.Annotation != nil {
		me.Annotation.initElement(p)
	}
}

func (me *hasElemsAny) initChildren(p element) {
	for _, any := range me.Anys {
		any.initElement(p)
	}
}

func (me *hasElemsAnyAttribute) initChildren(p element) {
	for _, aa := range me.AnyAttributes {
		aa.initElement(p)
	}
}

func (me *hasElemsAppInfo) initChildren(p element) {
	for _, ai := range me.AppInfos {
		ai.initElement(p)
	}
}

func (me *hasElemsAttribute) initChildren(p element) {
	for _, ea := range me.Attributes {
		ea.initElement(p)
	}
}

func (me *hasElemsAttributeGroup) initChildren(p element) {
	for _, ag := range me.AttributeGroups {
		ag.initElement(p)
	}
}

func (me *hasElemChoice) initChildren(p element) {
	if me.Choice != nil {
		me.Choice.initElement(p)
	}
}

func (me *hasElemsChoice) initChildren(p element) {
	for _, ch := range me.Choices {
		ch.initElement(p)
	}
}

func (me *hasElemComplexContent) initChildren(p element) {
	if me.ComplexContent != nil {
		me.ComplexContent.initElement(p)
	}
}

func (me *hasElemComplexType) initChildren(p element) {
	if me.ComplexType != nil {
		me.ComplexType.initElement(p)
	}
}

func (me *hasElemsComplexType) initChildren(p element) {
	for _, ct := range me.ComplexTypes {
		ct.initElement(p)
	}
}

func (me *hasElemsDocumentation) initChildren(p element) {
	for _, doc := range me.Documentations {
		doc.initElement(p)
	}
}

func (me *hasElemsElement) initChildren(p element) {
	for _, el := range me.Elements {
		el.initElement(p)
	}
}

func (me *hasElemsEnumeration) initChildren(p element) {
	for _, enum := range me.Enumerations {
		enum.initElement(p)
	}
}

func (me *hasElemExtensionComplexContent) initChildren(p element) {
	if me.ExtensionComplexContent != nil {
		me.ExtensionComplexContent.initElement(p)
	}
}

func (me *hasElemExtensionSimpleContent) initChildren(p element) {
	if me.ExtensionSimpleContent != nil {
		me.ExtensionSimpleContent.initElement(p)
	}
}

func (me *hasElemField) initChildren(p element) {
	if me.Field != nil {
		me.Field.initElement(p)
	}
}

func (me *hasElemFractionDigits) initChildren(p element) {
	if me.FractionDigits != nil {
		me.FractionDigits.initElement(p)
	}
}

func (me *hasElemGroup) initChildren(p element) {
	if me.Group != nil {
		me.Group.initElement(p)
	}
}

func (me *hasElemsGroup) initChildren(p element) {
	for _, gr := range me.Groups {
		gr.initElement(p)
	}
}

func (me *hasElemsImport) initChildren(p element) {
	for _, imp := range me.Imports {
		imp.initElement(p)
	}
}

func (me *hasElemsKey) initChildren(p element) {
	for _, k := range me.Keys {
		k.initElement(p)
	}
}

func (me *hasElemKeyRef) initChildren(p element) {
	if me.KeyRef != nil {
		me.KeyRef.initElement(p)
	}
}

func (me *hasElemLength) initChildren(p element) {
	if me.Length != nil {
		me.Length.initElement(p)
	}
}

func (me *hasElemList) initChildren(p element) {
	if me.List != nil {
		me.List.initElement(p)
	}
}

func (me *hasElemMaxExclusive) initChildren(p element) {
	if me.MaxExclusive != nil {
		me.MaxExclusive.initElement(p)
	}
}

func (me *hasElemMaxInclusive) initChildren(p element) {
	if me.MaxInclusive != nil {
		me.MaxInclusive.initElement(p)
	}
}

func (me *hasElemMaxLength) initChildren(p element) {
	if me.MaxLength != nil {
		me.MaxLength.initElement(p)
	}
}

func (me *hasElemMinExclusive) initChildren(p element) {
	if me.MinExclusive != nil {
		me.MinExclusive.initElement(p)
	}
}

func (me *hasElemMinInclusive) initChildren(p element) {
	if me.MinInclusive != nil {
		me.MinInclusive.initElement(p)
	}
}

func (me *hasElemMinLength) initChildren(p element) {
	if me.MinLength != nil {
		me.MinLength.initElement(p)
	}
}

func (me *hasElemsNotation) initChildren(p element) {
	for _, not := range me.Notations {
		not.initElement(p)
	}
}

func (me *hasElemPattern) initChildren(p element) {
	if me.Pattern != nil {
		me.Pattern.initElement(p)
	}
}

func (me *hasElemsRedefine) initChildren(p element) {
	for _, rd := range me.Redefines {
		rd.initElement(p)
	}
}

func (me *hasElemRestrictionComplexContent) initChildren(p element) {
	if me.RestrictionComplexContent != nil {
		me.RestrictionComplexContent.initElement(p)
	}
}

func (me *hasElemRestrictionSimpleContent) initChildren(p element) {
	if me.RestrictionSimpleContent != nil {
		me.RestrictionSimpleContent.initElement(p)
	}
}

func (me *hasElemRestrictionSimpleType) initChildren(p element) {
	if me.RestrictionSimpleType != nil {
		me.RestrictionSimpleType.initElement(p)
	}
}

func (me *hasElemSelector) initChildren(p element) {
	if me.Selector != nil {
		me.Selector.initElement(p)
	}
}

func (me *hasElemSequence) initChildren(p element) {
	if me.Sequence != nil {
		me.Sequence.initElement(p)
	}
}

func (me *hasElemsSequence) initChildren(p element) {
	for _, seq := range me.Sequences {
		seq.initElement(p)
	}
}

func (me *hasElemSimpleContent) initChildren(p element) {
	if me.SimpleContent != nil {
		me.SimpleContent.initElement(p)
	}
}

func (me *hasElemsSimpleType) initChildren(p element) {
	for _, st := range me.SimpleTypes {
		st.initElement(p)
	}
}

func (me *hasElemTotalDigits) initChildren(p element) {
	if me.TotalDigits != nil {
		me.TotalDigits.initElement(p)
	}
}

func (me *hasElemUnion) initChildren(p element) {
	if me.Union != nil {
		me.Union.initElement(p)
	}
}

func (me *hasElemUnique) initChildren(p element) {
	if me.Unique != nil {
		me.Unique.initElement(p)
	}
}

func (me *hasElemWhiteSpace) initChildren(p element) {
	if me.WhiteSpace != nil {
		me.WhiteSpace.initElement(p)
	}
}
