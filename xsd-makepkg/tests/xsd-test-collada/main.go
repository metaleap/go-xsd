package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	util "github.com/metaleap/go-util"

	collada14 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_4_go"
	collada15 "github.com/metaleap/go-xsd-pkg/khronos.org/files/collada_schema_1_5_go"
)

type Col14Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	collada14.TxsdCollada
}

type Col15Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	collada15.TxsdCollada
}

func main() {
	var (
		col14DirBasePath  = util.GopathSrcGithub("metaleap", "go-xsd", "makepkg", "tests", "collada", "1.4.1")
		col14MakeEmptyDoc = func() interface{} { return &Col14Doc{} }
		col15DirBasePath  = util.GopathSrcGithub("metaleap", "go-xsd", "makepkg", "tests", "collada", "1.5")
		col15MakeEmptyDoc = func() interface{} { return &Col15Doc{} }
	)
	tests.TestViaRemarshal(col14DirBasePath, col14MakeEmptyDoc)
	tests.TestViaRemarshal(col15DirBasePath, col15MakeEmptyDoc)
}
