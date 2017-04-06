go-xsd
======


A Go package for loading ( **xml.Unmarshal()**ing ) an XML Schema Definition (XSD) document into an **xsd.Schema** structure.

With this, you could probably write an XML validator, or otherwise utilize or further process the loaded XSD --- but the main use-case here was:


go-xsd/xsd-makepkg
==================


A command-line tool to generate Go "XML wrapper" package sources for specified XSD schema URIs.

If no arguments are specified, this tool proceeds to (re)generate all Go packages for various common XML formats in your local $GOPATH-relative directory corresponding to the **http://github.com/metaleap/go-xsd-pkg** repository. For more details on command-line arguments for *xsd-makepkg*: scroll down to the bottom of this readme.

Each generated wrapper package contains the type structures required to easily **xml.Unmarshal()** an XML document based on that XSD.

**XSD simple-types** are represented by the corresponding native Go scalar data type, augmented by utility methods where applicable:

- enumerated simple-types (eg. "value can be 'sunny', 'cloudy', 'rainy'") get handy corresponding **IsXyz() bool** methods (eg. "IsSunny()", "IsCloudy()", "IsRainy()")

- simple-types that define a whitespace-separated list of scalar values get a corresponding, properly typed **Values()** method

- for attributes or elements that define a fixed or default value, their corresponding generated Go simple-type will have a properly typed *ElemnameDefault()* / *ElemnameFixed()* / *AttnameDefault()* / *AttnameFixed()* method (eg. if the *langpref* attribute is defined to default to "Go", then its simple-type will have a *LangprefDefault()* method returning "Go")

**XSD complex-types**, attribute-groups, element-groups, elements etc. are ultimately represented by corresponding generated Go struct types.

**XSD includes** are all loaded and processed together into a single output .go source file.

**XSD imports** are rewritten as Go imports but not otherwise auto-magically processed. If you see the generated .go package importing another "some-xsd-xml-whatever-name-_go" package that will cause a "package not found" compiler error, then to make that import work, you'll first need to also auto-generate that package with *xsd-makepkg* yourself as well.

**XSD documentation annotation** is rewritten as Go // code comments. Yeah, that's rather neat.

Regarding the auto-generated code:

- it's **by necessity not idiomatic** and most likely not as terse/slim as manually-written structs would be. For very simplistic XML formats, writing your own 3 or 4 custom structs might be a tiny bit more efficient. **For highly intricate, unwieldy XML formats, the auto-generated packages beat hand-writing 100s of custom structs, however.** Auto-generated code will never win a code-beauty contest, you're expected to simply import the compiled package rather than having to work inside its generated source files.

- most (XSD-declared) types are prefixed with T -- thanks to the SVG schema which taught me that in XSD, "scriptType" and "ScriptType" are two valid and uniquely different type names. To have all types exported from the generated Go package, then, some kind of prefix is indeed needed.

- most XSDs are chock-full of anonymous types, as well as implicit ones (unions, restrictions, extensions...) Go does support "anonymous types" per se, but I decided against using them. Every type is declared and exported, no anonymous magic. This makes most auto-generated packages "look" even more confusing than their XSD counterparts at first glance. Indeed they may appear quite bloated, and when coding with the imported generated package you'll probably be better off working with the particular XML format's specification document rather than the **godoc** for the generated package... this is not a perfect situation but at least for now I can work with this for the few XML formats I occasionally need to "parse, convert and forget" -- ultimately, most XML formats at my end are mere interchange or legacy formats, and never really the "main topic" at hand.


go-xsd/types
============


A tiny package automatically imported by all **go-xsd** auto-generated packages.
Maps all XSD built-in simple-types to Go types, which affords us easy mapping of any XSD type references in the schema to Go imports: every *xs:string* and *xs:boolean* automatically becomes *xsdt.String* and *xsdt.Boolean* etc.
Types are mapped to Go types depending on how **encoding/xml.Unmarshal()** can handle them: ie. it parses bools and numbers, but dates/durations have too many format mismatches and thus are just declared string types.
Same for base64- and hex-encoded binary data: since **Unmarshal()** won't decode them, we leave them as strings. If you need their binary data, your code needs to import Go's base64/hex codec packages and use them as necessary.


How to use auto-generated packages:
===================================


Take a look at the "test progs" under **xsd-makepkg/tests**, they're basically simple usage examples. For unmarshal you need to define just one small custom struct like this --- using the rss package as a simple example, as demonstrated in **xsd-makepkg/tests/xsd-test-rss/main.go**:


    type MyRssDoc struct {
        XMLName xml.Name `xml:"rss"`
        rss.TxsdRss
    }


So your custom struct specifies two things:

- the XML name of the root element in your XML file, as is typical when working with **encoding/xml.Unmarshal()**.

- the Go struct type from the auto-generated package to **embed right inside** your custom struct.

The second part is the only tricky part. XML Schema Definition has no real concept of "root element", partly because they're designed to support use-cases where you embed a full document defined in one XSD deep inside a full document defined in another XSD. So a Collada document may contain a full or partial MathML document somewhere inside it. Some well-designed XSDs define a single top-level element, so we could infer "this is the root element" and generate a "XyzDoc" struct (like the *MyRssDoc* above) for you. But many don't. Some formats may legally have one of two or more possible "root" elements, ie. Atom allegedly may have either a "feed" root element or an "entry" root element. So **go-xsd** does not magically infer which of the XSD's top-level elements might be the root element, you define this by writing a small struct as shown above. The naming of the root element Go type to be embedded is not consistent across different packages, because their naming is directly based on the XSD that was used to generate the package. So for example...

- for rss we have *rss.TxsdRss*
- for atom: *atom.TentryType* and *atom.TfeedType*
- for svg: *svg.TsvgType*
- for Collada: *collada.TxsdCollada*

Seems like Collada and RSS share a certain naming pattern, and yet Atom/SVG share another one? Mere coincidence, the naming is completely arbitrary and up to the XSD author. Ultimately, to find out the proper Go type name to embed, you'll have to dig a bit inside the generated package. That's actually pretty straightforward, here's how you do it:

A) Suppose you have an XML format where the root element (and only that one) is known to be named:


    <gopher>


B) Open the generated Go package source files under **$GOPATH/src/github.com/metaleap/go-xsd-pkg/yourdomain.org/xsd/gopher.xsd_go/*.go** (unless you used custom paths when you ran the **go-xsd/xsd-makepkg** tool)

C) Search for an occurence of either:


    "gopher"`


( quote, gopher, quote, backtick ), or:


     gopher"`


( *whitespace*, gopher, quote, backtick )

D) The found occurence is likely the tag for a field in a type named something like **XsdGoPkgHasElem_Gopher** or **XsdGoPkgHasElems_Gopher**. Ignore that type, instead focus on the type of the field itself. That's the one you're looking for, the one to embed in your tiny custom doc struct.


Command-line flags for *go-xsd/xsd-makepkg* tool:
=================================================


- **-basepath=""**: Defaults to github.com/metaleap/go-xsd-pkg. A $GOPATH/src/-relative path (always a slash-style path, even on Windows) where XSD files are downloaded to / loaded from and generated Go wrapper packages are created. Any XSD imports are also rewritten as Go imports from that path (but are not otherwise auto-magically processed in any way).
- **-local=true**: Local copy -- only downloads if file does not exist locally
- **-parse=false**: Not necessary, unless the generated Go wrapper package fails to compile with either the error "*cannot convert {value} to type {type}*" or "*cannot use {value} (type {type}) as type {type} in return argument*" -- ultimately down to a slightly faulty XSD file, but while rare, those exist (hello, KML using 0 and 1 for *xs:boolean*s that are clearly spec'd to be only ever either *true* or *false*...)
- **-uri=""**: The XML Schema Definition file URIs to generate a Go wrapper packages from, whitespace-separated. (For each, the protocol prefix can be omitted, it then defaults to *http://*. Only protocols understood by the *net/http* package are supported.)
- **-gofmt=true**: Run 'gofmt' against the generated Go wrapper package?
- **-goinst=true**: Run 'go-buildrun' ( http://github.com/metaleap/go-buildrun ) against the generated Go wrapper package?
