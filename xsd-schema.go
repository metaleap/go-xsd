package xsd

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	uio "github.com/metaleap/go-util/io"
	unet "github.com/metaleap/go-util/net"
	ustr "github.com/metaleap/go-util/str"
)

const (
	goPkgPrefix = ""
	goPkgSuffix = "_gopkg"
	protSep = "://"
	xsdNamespaceUri = "http://www.w3.org/2001/XMLSchema"
)

var (
	loadedSchemas = map[string]*Schema {}
)

type Schema struct {
	elemBase
	XMLName xml.Name `xml:"schema"`
	XMLNamespaces map[string]string `xml:"-"`
	XMLIncludedSchemas []*Schema `xml:"-"`
	XSDNamespace string `xml:"-"`
	XSDParentSchema *Schema `xml:"-"`

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

	loadLocalPath, loadUri string
}

	func (me *Schema) onLoad (rootAtts []xml.Attr, loadUri, localPath string) (err error) {
		var tmpUrl string
		var sd *Schema
		loadedSchemas[loadUri] = me
		me.loadLocalPath, me.loadUri = localPath, loadUri
		me.XMLNamespaces = map[string]string {}
		for _, att := range rootAtts {
			if att.Name.Space == "xmlns" {
				me.XMLNamespaces[att.Name.Local] = att.Value
			} else if len(att.Name.Space) > 0 {

			} else if att.Name.Local == "xmlns" {
				me.XMLNamespaces[""] = att.Value
			}
		}
		for k, v := range me.XMLNamespaces { if v == xsdNamespaceUri { me.XSDNamespace = k } }
		if len(me.XMLNamespaces["xml"]) == 0 { me.XMLNamespaces["xml"] = "http://www.w3.org/XML/1998/namespace" }
		me.XMLIncludedSchemas = []*Schema {}
		for _, inc := range me.Includes {
			if tmpUrl = inc.SchemaLocation.String(); strings.Index(tmpUrl, protSep) < 0 {
				tmpUrl = path.Join(path.Dir(loadUri), tmpUrl)
			}
			if sd = loadedSchemas[tmpUrl]; sd == nil {
				if sd, err = LoadSchema(tmpUrl, len(localPath) > 0); err != nil { return }
			}
			sd.XSDParentSchema = me
			me.XMLIncludedSchemas = append(me.XMLIncludedSchemas, sd)
		}
		me.initElement(nil)
		return
	}

	func (me *Schema) MakeGoPkgSrc () string {
		var bag = &PkgBag { Schema: me }
		me.makePkg(bag)
		return strings.Join(bag.lines, "\n")
	}

	func (me *Schema) MakeGoPkgSrcFile () (goOutFilePath string, err error) {
		var goOutDirPath = filepath.Join(filepath.Dir(me.loadLocalPath), goPkgPrefix + filepath.Base(me.loadLocalPath) + goPkgSuffix)
		goOutFilePath = filepath.Join(goOutDirPath, path.Base(me.loadUri) + ".go")
		if err = me.MakeGoPkgSrcFileAt(goOutFilePath); err == nil {
			for _, inc := range me.XMLIncludedSchemas {
				var s = strings.Replace(inc.loadUri, strings.Trim(path.Dir(me.loadUri), "/") + "/", "", -1)
				if err = inc.MakeGoPkgSrcFileAt(filepath.Join(goOutDirPath, ustr.Replace(s, map[string]string { "/": "_", ":": "_", "..": "__" }) + ".go")); err != nil { break }
			}
		}
		return
	}

	func (me *Schema) MakeGoPkgSrcFileAt (goOutFilePath string) (err error) {
		if err = uio.EnsureDirExists(filepath.Dir(goOutFilePath)); err == nil {
			err = uio.WriteTextFile(goOutFilePath, me.MakeGoPkgSrc())
		}
		return
	}

	func (me *Schema) RootSchema () *Schema {
		if me.XSDParentSchema != nil { return me.XSDParentSchema.RootSchema() }
		return me
	}

func ClearLoadedSchemasCache () {
	loadedSchemas = map[string]*Schema {}
}

func loadSchema (r io.Reader, loadUri, localPath string) (sd *Schema, err error) {
	var data []byte
	var rootAtts []xml.Attr
	if data, err = ioutil.ReadAll(r); err == nil {
		var t xml.Token; sd = new(Schema)
		for xd := xml.NewDecoder(bytes.NewReader(data)); err == nil; {
			if t, err = xd.Token(); err == nil {
				if startEl, ok := t.(xml.StartElement); ok { rootAtts = startEl.Attr; break }
			}
		}
		if err = xml.Unmarshal(data, sd); err == nil { err = sd.onLoad(rootAtts, loadUri, localPath) }
		if err != nil { sd = nil }
	}
	return
}

func loadSchemaFile (filename string, loadUri string) (sd *Schema, err error) {
	var file *os.File
	if file, err = os.Open(filename); err == nil {
		defer file.Close()
		sd, err = loadSchema(file, loadUri, filename)
	}
	return
}

func LoadSchema (uri string, localCopy bool) (sd *Schema, err error) {
	var protocol, localPath string
	var rc io.ReadCloser

	if pos := strings.Index(uri, protSep); pos < 0 { protocol = "http" + protSep } else { protocol = uri[: pos + len(protSep)]; uri = uri[pos + len(protSep) :] }
	if localCopy {
		if localPath = filepath.Join(PkgGen.BaseCodePath, uri); !uio.FileExists(localPath) {
			if err = uio.EnsureDirExists(filepath.Dir(localPath)); err == nil { err = unet.DownloadFile(protocol + uri, localPath) }
		}
		if err == nil { if sd, err = loadSchemaFile(localPath, uri); sd != nil { sd.loadLocalPath = localPath } }
	} else if rc, err = unet.OpenRemoteFile(protocol + uri); err == nil {
		defer rc.Close()
		sd, err = loadSchema(rc, uri, "")
	}
	return
}
