package xsd

func (me *hasElemAnnotation) makePkg (bag *makerBag) {
	if me.Annotation != nil { me.Annotation.makePkg(bag) }
}

func (me *hasElemsDocumentation) makePkg (bag *makerBag) {
	for _, doc := range me.Documentations { doc.makePkg(bag) }
}

func (me *hasElemsImport) makePkg (bag *makerBag) {
	for _, imp := range me.Imports { imp.makePkg(bag) }
}

func (me *hasElemsSimpleType) makePkg (bag *makerBag) {
	for _, st := range me.SimpleTypes { st.makePkg(bag); PkgGen.append("") }
}
