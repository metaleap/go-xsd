package xsd

func (me *hasElemAnnotation) makePkg (bag *makerBag) {
	if me.Annotation != nil { me.Annotation.makePkg(bag) }
}

func (me *hasElemsDocumentation) makePkg (bag *makerBag) {
	for _, doc := range me.Documentations { doc.makePkg(bag) }
}

func (me *hasElemsEnumeration) makePkg (bag *makerBag) {
	for _, enum := range me.Enumerations { enum.makePkg(bag) }
}

func (me *hasElemsImport) makePkg (bag *makerBag) {
	for _, imp := range me.Imports { imp.makePkg(bag) }
}

func (me *hasElemList) makePkg (bag *makerBag) {
	if me.List != nil { me.List.makePkg(bag) }
}

func (me *hasElemsNotation) makePkg (bag *makerBag) {
	if len(me.Notations) > 0 {
		PkgGen.impsUsed[PkgGen.impName] = true
		PkgGen.appendFmt(false, "var Notations = new(%v.Notations)\n\nfunc init () {", PkgGen.impName)
		for _, not := range me.Notations { not.makePkg(bag) }
		PkgGen.appendFmt(true, "}")
	}
}

func (me *hasElemsSimpleType) makePkg (bag *makerBag) {
	for _, st := range me.SimpleTypes { st.makePkg(bag) }
}
