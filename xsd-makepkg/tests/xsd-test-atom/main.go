package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	"github.com/metaleap/go-util-misc"

	atom "github.com/metaleap/go-xsd-pkg/kbcafe.com/rss/atom.xsd.xml_go"
)

type AtomEntryDoc struct {
	XMLName xml.Name `xml:"entry"`
	atom.TentryType
}

type AtomFeedDoc struct {
	XMLName xml.Name `xml:"feed"`
	atom.TfeedType
}

func main() {
	var (
		entryDirBasePath  = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "entry")
		entryMakeEmptyDoc = func() interface{} { return &AtomEntryDoc{} }
		feedDirBasePath   = ugo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-atom", "feed")
		feedMakeEmptyDoc  = func() interface{} { return &AtomFeedDoc{} }
	)
	tests.TestViaRemarshal(entryDirBasePath, entryMakeEmptyDoc)
	tests.TestViaRemarshal(feedDirBasePath, feedMakeEmptyDoc)
}
