package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	util "github.com/metaleap/go-util"

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
		entryDirBasePath  = util.BaseCodePathGithub("metaleap", "go-xsd", "makepkg", "tests", "atom", "entry")
		entryMakeEmptyDoc = func() interface{} { return &AtomEntryDoc{} }
		feedDirBasePath   = util.BaseCodePathGithub("metaleap", "go-xsd", "makepkg", "tests", "atom", "feed")
		feedMakeEmptyDoc  = func() interface{} { return &AtomFeedDoc{} }
	)
	tests.TestViaRemarshal(entryDirBasePath, entryMakeEmptyDoc)
	tests.TestViaRemarshal(feedDirBasePath, feedMakeEmptyDoc)
}
