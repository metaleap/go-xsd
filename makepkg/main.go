package main

import (
	"fmt"

	xsd "github.com/metaleap/go-xsd"
)

func main () {
	var sd *xsd.Schema
	var err error
	var outFilePath string
	var schemas = []string {
		"docs.oasis-open.org/election/external/xAL.xsd",
		"www.w3.org/2001/03/xml.xsd",
		"docbook.org/xml/5.0/xsd/docbook.xsd",
		"kbcafe.com/rss/atom.xsd.xml",
		"thearchitect.co.uk/schemas/rss-2_0.xsd",
		"www.w3.org/2007/schema-for-xslt20.xsd",
		"www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd",
		"www.w3.org/Math/XMLSchema/mathml2/mathml2.xsd",
		"www.w3.org/TR/2007/REC-voicexml21-20070619/vxml.xsd",
		"khronos.org/files/collada_schema_1_4",
		"khronos.org/files/collada_schema_1_5",
		"schemas.opengis.net/kml/2.2.0/ogckml22.xsd",
		"schemas.opengis.net/kml/2.2.0/atom-author-link.xsd",
	}
	for _, s := range schemas {
		fmt.Printf("LOAD: %v\n", s)
		if sd, err = xsd.LoadSchema(s, true); err != nil {
			panic(err)
		} else if sd != nil {
			if outFilePath, err = sd.MakeGoPkgSrcFile(); err == nil {
				fmt.Printf("\tGEN: %v\n", outFilePath)
			} else {
				panic(err)
			}
		}
	}
}
