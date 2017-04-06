# tests
--
    import "github.com/metaleap/go-xsd/xsd-makepkg/tests"

	A simple test function shared by the various test programs inside this directory (rss, atom, collada, svg etc.)

## Usage

```go
var (
	OnDocLoaded func(interface{})
)
```

#### func  TestViaRemarshal

```go
func TestViaRemarshal(dirPath string, makeEmptyDoc func() interface{})
```
Attempts to xml.Unmarshal() all files in the "infiles" sub-directory of the
specified directory path into the interface{} structure returned by the
specified constructor. For each such input file, then attempts to
xml.MarshalIndent() said structure back into a new output XML file with the same
name, in the "outfiles" sub-directory of the specified directory path.

--
**godocdown** http://github.com/robertkrimen/godocdown