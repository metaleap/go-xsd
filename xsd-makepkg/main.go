package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"

	"github.com/metaleap/go-util-misc"

	xsd "github.com/metaleap/go-xsd"
)

var (
	flagGoFmt      = flag.Bool("gofmt", true, "Run 'gofmt' against the generated Go wrapper package?")
	flagGoInst     = flag.Bool("goinst", true, "Run 'go-buildrun' against the generated Go wrapper package?")
	flagSchema     = flag.String("uri", "", "The XML Schema Definition file URIs to generate a Go wrapper packages from, whitespace-separated. (For each, the protocol prefix can be omitted, it then defaults to http://. Only protocols understood by the net/http package are supported.)")
	flagLocalCopy  = flag.Bool("local", true, "Local copy: only downloads if file does not exist locally")
	flagForceParse = flag.Bool("parse", false, "Not necessary unless the generated Go wrapper package won't compile.")
	flagBasePath   = flag.String("basepath", "", "Defaults to "+xsd.PkgGen.BasePath+". A $GOPATH/src/-relative path (always a slash-style path, even on Windows) where XSD files are downloaded to / loaded from and generated Go wrapper packages are created. Any XSD imports are also rewritten as Go imports from that path (but are not otherwise auto-magically processed in any way).")

	//	if no schemas are specified in *flagSchema, we run the pkg-maker through a default series of various XSDs...
	schemas = []string{
		"www.w3.org/2001/xml.xsd",
		"www.w3.org/2001/03/xml.xsd",
		"www.w3.org/TR/2002/WD-SVG11-20020108/xml.xsd",
		"www.w3.org/TR/2002/WD-SVG11-20020108/xlink.xsd",
		"www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd",
		"www.w3.org/2007/schema-for-xslt20.xsd",
		"www.w3.org/Math/XMLSchema/mathml2/common/xlink-href.xsd",
		"www.w3.org/Math/XMLSchema/mathml2/mathml2.xsd",
		"docs.oasis-open.org/election/external/xAL.xsd",
		"docbook.org/xml/5.0/xsd/xml.xsd",
		"docbook.org/xml/5.0/xsd/xlink.xsd",
		"docbook.org/xml/5.0/xsd/docbook.xsd",
		"kbcafe.com/rss/atom.xsd.xml",
		"thearchitect.co.uk/schemas/rss-2_0.xsd",
		"schemas.opengis.net/kml/2.2.0/atom-author-link.xsd",
		"schemas.opengis.net/kml/2.2.0/ogckml22.xsd",
		"khronos.org/files/collada_schema_1_4",
		"khronos.org/files/collada_schema_1_5",
	}
)

func main() {
	var (
		sd          *xsd.Schema
		err         error
		raw         []byte
		outFilePath string
	)
	flag.Parse()
	if len(*flagSchema) > 0 {
		schemas = strings.Split(*flagSchema, " ")
	}
	if len(*flagBasePath) > 0 {
		xsd.PkgGen.BasePath, xsd.PkgGen.BaseCodePath = *flagBasePath, ugo.GopathSrc(strings.Split(*flagBasePath, "/")...)
	}
	for _, s := range schemas {
		log.Printf("LOAD:\t%v\n", s)
		if sd, err = xsd.LoadSchema(s, *flagLocalCopy); err != nil {
			log.Printf("\tERROR: %v\n", err)
		} else if sd != nil {
			xsd.PkgGen.ForceParseForDefaults = *flagForceParse || (s == "schemas.opengis.net/kml/2.2.0/ogckml22.xsd") // KML schema uses 0 and 1 as defaults for booleans...
			if outFilePath, err = sd.MakeGoPkgSrcFile(); err == nil {
				log.Printf("MKPKG:\t%v\n", outFilePath)
				if *flagGoFmt {
					if raw, err = exec.Command("gofmt", "-w=true", "-s=true", "-e=true", outFilePath).CombinedOutput(); len(raw) > 0 {
						log.Printf("GOFMT:\t%s\n", string(raw))
					}
					if err != nil {
						log.Printf("GOFMT:\t%v\n", err)
					}
				}
				if *flagGoInst {
					if raw, err = exec.Command("go-buildrun", "-d=__doc.html", "-f="+outFilePath).CombinedOutput(); len(raw) > 0 {
						println(string(raw))
					}
					if err != nil {
						log.Printf("GOINST:\t%v\n", err)
					}
				}
			} else {
				log.Printf("\tERROR:\t%v\n", err)
			}
		}
	}
}
