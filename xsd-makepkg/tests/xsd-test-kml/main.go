package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	"github.com/metaleap/go-util-misc"

	kml "github.com/metaleap/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
)

type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
}

func main() {
	var (
		dirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
