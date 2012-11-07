package xsd

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
)

type Schema struct {
	XMLName xml.Name `xml:"schema"`

	hasAttrAttributeFormDefault
	hasAttrBlockDefault
	hasAttrElementFormDefault
	hasAttrFinalDefault
	hasAttrLang
	hasAttrId
	hasAttrSchemaLocation
	hasAttrTargetNamespace
	hasAttrVersion
	hasElemAnnotation
	hasElemsAttribute
	hasElemsAttributeGroup
	hasElemsComplexType
	hasElemsElement
	hasElemsGroup
	hasElemsInclude
	hasElemsImport
	hasElemsNotation
	hasElemsRedefine
	hasElemsSimpleType

	schemaLocations []string
}

	func (me *Schema) MakeGoPkgSrc () string {
		me.makePkg()
		return strings.Join(PkgGen.lines, "\n")
	}

	func (me *Schema) SchemaLocations () []string {
		return me.schemaLocations
	}

func LoadSchema (r io.Reader) (sd *Schema, err error) {
	var data []byte
	if data, err = ioutil.ReadAll(r); err != nil { data = nil } else { sd, err = NewSchema(data) }
	return
}

func LoadSchemaFile (filename string) (sd *Schema, err error) {
	var file *os.File
	if file, err = os.Open(filename); err == nil {
		defer file.Close()
		sd, err = LoadSchema(file)
	}
	return
}

func NewSchema (data []byte) (sd *Schema, err error) {
	sd = new(Schema)
	err = xml.Unmarshal(data, sd)
	sd.schemaLocations = strings.Split(sd.SchemaLocation, " ")
	return
}
