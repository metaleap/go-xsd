package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	util "github.com/metaleap/go-util"

	kml "github.com/metaleap/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
)

type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
}

func main() {
	var (
		dirBasePath  = util.BaseCodePathGithub("metaleap", "go-xsd", "makepkg", "tests", "kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
