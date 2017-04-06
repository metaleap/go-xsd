package tests

import (
	"encoding/xml"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util-fs"

	xmlx "github.com/go-forks/go-pkg-xmlx"
)

var (
	OnDocLoaded func(interface{})
)

func verifyDocs(origData, faksData []byte) (errs []error) {
	orig, faks := xmlx.New(), xmlx.New()
	err := orig.LoadBytes(origData, nil)
	if err == nil {
		if err = faks.LoadBytes(faksData, nil); err == nil {
			errs = verifyNode(orig.Root, faks.Root)
		}
	}
	return
}

func verifyNode(orig, faks *xmlx.Node) (errs []error) {
	type both struct {
		origNodes, faksNodes []*xmlx.Node
	}
	var (
		curBoth *both
		cn      *xmlx.Node
		tmp     string
		i       int
		subErrs []error
	)
	attVal := func(xn *xmlx.Node, att *xmlx.Attr) (v string) {
		if v = xn.As("", att.Name.Local); len(v) == 0 {
			v = xn.As(att.Name.Space, att.Name.Local)
		}
		return
	}
	cleanNodes := func(xns ...*xmlx.Node) {
		for _, xn := range xns {
			for _, cn = range xn.Children {
				if cn.Type != xmlx.NT_ELEMENT {
					xn.RemoveChild(cn)
				}
			}
		}
	}
	cleanNodes(orig, faks)
	for _, a := range orig.Attributes {
		if tmp = attVal(faks, a); tmp != a.Value {
			errs = append(errs, fmt.Errorf("Attribute '%s:%s' of <%s> element: different values (orig='%s' faks='%s')", a.Name.Space, a.Name.Local, orig.Name.Local, a.Value, tmp))
		}
	}
	if len(orig.Children) > len(faks.Children) {
		errs = append(errs, fmt.Errorf("Orig <%s> element has %v children, but faks has %v.", orig.Name.Local, len(orig.Children), len(faks.Children)))
	}
	if orig.Value != faks.Value {
		errs = append(errs, fmt.Errorf("Orig <%s> element value differs from faks value.", orig.Name.Local))
	}
	namedNodes := map[string]*both{}
	for _, cn = range orig.Children {
		if curBoth = namedNodes[cn.Name.Local]; curBoth == nil {
			curBoth = &both{}
			namedNodes[cn.Name.Local] = curBoth
		}
		curBoth.origNodes = append(curBoth.origNodes, cn)
	}
	for _, cn = range faks.Children {
		if curBoth = namedNodes[cn.Name.Local]; curBoth != nil {
			curBoth.faksNodes = append(curBoth.faksNodes, cn)
		}
	}
	for tmp, curBoth = range namedNodes {
		if len(curBoth.origNodes) != len(curBoth.faksNodes) {
			errs = append(errs, fmt.Errorf("Orig <%s> element has %v <%s> elements but faks <%s> element has %v.", orig.Name.Local, len(curBoth.origNodes), tmp, faks.Name.Local, len(curBoth.faksNodes)))
		} else if len(curBoth.origNodes) == 1 {
			errs = append(errs, verifyNode(curBoth.origNodes[0], curBoth.faksNodes[0])...)
		} else {
			for i, cn = range curBoth.origNodes {
				if subErrs = verifyNode(cn, curBoth.faksNodes[i]); len(subErrs) > 0 {
					errs = append(errs, subErrs...)
				}
			}
		}
	}
	return
}

//	Attempts to xml.Unmarshal() all files in the "infiles" sub-directory of the specified directory path into the interface{} structure returned by the specified constructor.
//	For each such input file, then attempts to xml.MarshalIndent() said structure back into a new output XML file with the same name, in the "outfiles" sub-directory of the specified directory path.
func TestViaRemarshal(dirPath string, makeEmptyDoc func() interface{}) {
	var dirPathInFiles = filepath.Join(dirPath, "infiles")
	var dirPathOutFiles = filepath.Join(dirPath, "outfiles")
	var loadXmlDocFile = func(filename string) bool {
		log.Printf("Loading %s", filename)
		doc, dataOrig := makeEmptyDoc(), ufs.ReadBinaryFile(filename, true)
		err := xml.Unmarshal(dataOrig, doc)
		if err != nil {
			panic(err)
		}
		if OnDocLoaded != nil {
			OnDocLoaded(doc)
		}
		outFileName := filepath.Join(dirPathOutFiles, filepath.Base(filename))
		log.Printf("Writing %s", outFileName)
		dataFaks, err := xml.MarshalIndent(doc, "", "\t")
		if err != nil {
			panic(err)
		}
		ufs.WriteTextFile(outFileName, strings.Trim(string(dataFaks), " \r\n\t"))
		log.Printf("Verifying...")
		if errs := verifyDocs(dataOrig, dataFaks); len(errs) > 0 {
			for _, err = range errs {
				log.Printf("%v", err)
			}
		}
		return true
	}
	if errs := ufs.NewDirWalker(false, nil, loadXmlDocFile).Walk(dirPathInFiles); len(errs) > 0 {
		panic(errs[0])
	}
}
