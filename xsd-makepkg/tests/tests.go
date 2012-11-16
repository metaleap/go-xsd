//	A simple test function shared by the various test programs inside this directory (rss, atom, collada, svg etc.)
package tests

import (
	"encoding/xml"
	"path/filepath"
	"strings"

	uio "github.com/metaleap/go-util/io"
)

//	Attempts to xml.Unmarshal() all files in the "infiles" sub-directory of the specified directory path into the interface{} structure returned by the specified constructor.
//	For each such input file, then attempts to xml.MarshalIndent() said structure back into a new output XML file with the same name, in the "outfiles" sub-directory of the specified directory path.
func TestViaRemarshal (dirPath string, makeEmptyDoc func () interface{}) {
	var dirPathInFiles = filepath.Join(dirPath, "infiles")
	var dirPathOutFiles = filepath.Join(dirPath, "outfiles")
	var loadXmlDocFile = func (filename string, keepRecursing bool) bool {
		var (
			outFileName = filepath.Join(dirPathOutFiles, filepath.Base(filename))
			data = uio.ReadBinaryFile(filename, true)
			err error
			doc = makeEmptyDoc()
		)
		if err = xml.Unmarshal(data, doc); err != nil { panic(err) }
		if data, err = xml.MarshalIndent(doc, "", "\t"); err != nil { panic(err) }
		uio.WriteTextFile(outFileName, strings.Trim(string(data), " \r\n\t"))
		return keepRecursing
	}
	uio.WalkDirectory(dirPathInFiles, "", loadXmlDocFile, true)
}
