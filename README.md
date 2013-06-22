# xsd
--
    import "github.com/metaleap/go-xsd"


## Usage

```go
var (
	PkgGen = &pkgGen{
		BaseCodePath:             ugo.GopathSrcGithub("metaleap", "go-xsd-pkg"),
		BasePath:                 "github.com/metaleap/go-xsd-pkg",
		ForceParseForDefaults:    false,
		PluralizeSpecialPrefixes: []string{"Library", "Instance"},
		AddWalkers:               true,
	}
)
```

#### func  ClearLoadedSchemasCache

```go
func ClearLoadedSchemasCache()
```

#### func  Flattened

```go
func Flattened(choices []*Choice, seqs []*Sequence) (allChoices []*Choice, allSeqs []*Sequence)
```

#### type All

```go
type All struct {
}
```


#### func (*All) Parent

```go
func (me *All) Parent() element
```

#### func (*All) Value

```go
func (me *All) Value() (l xsdt.Long)
```

#### type Annotation

```go
type Annotation struct {
}
```


#### func (*Annotation) Parent

```go
func (me *Annotation) Parent() element
```

#### type Any

```go
type Any struct {
}
```


#### func (*Any) Parent

```go
func (me *Any) Parent() element
```

#### func (*Any) Value

```go
func (me *Any) Value() (l xsdt.Long)
```

#### type AnyAttribute

```go
type AnyAttribute struct {
}
```


#### func (*AnyAttribute) Parent

```go
func (me *AnyAttribute) Parent() element
```

#### type AppInfo

```go
type AppInfo struct {
}
```


#### func (*AppInfo) Parent

```go
func (me *AppInfo) Parent() element
```

#### type Attribute

```go
type Attribute struct {
}
```


#### func (*Attribute) Parent

```go
func (me *Attribute) Parent() element
```

#### type AttributeGroup

```go
type AttributeGroup struct {
}
```


#### func (*AttributeGroup) Parent

```go
func (me *AttributeGroup) Parent() element
```

#### type Choice

```go
type Choice struct {
}
```


#### func (*Choice) Parent

```go
func (me *Choice) Parent() element
```

#### func (*Choice) Value

```go
func (me *Choice) Value() (l xsdt.Long)
```

#### type ComplexContent

```go
type ComplexContent struct {
}
```


#### func (*ComplexContent) Parent

```go
func (me *ComplexContent) Parent() element
```

#### type ComplexType

```go
type ComplexType struct {
}
```


#### func (*ComplexType) Parent

```go
func (me *ComplexType) Parent() element
```

#### type Documentation

```go
type Documentation struct {
}
```


#### func (*Documentation) Parent

```go
func (me *Documentation) Parent() element
```

#### type Element

```go
type Element struct {
}
```


#### func (*Element) Parent

```go
func (me *Element) Parent() element
```

#### func (*Element) Value

```go
func (me *Element) Value() (l xsdt.Long)
```

#### type ExtensionComplexContent

```go
type ExtensionComplexContent struct {
}
```


#### func (*ExtensionComplexContent) Parent

```go
func (me *ExtensionComplexContent) Parent() element
```

#### type ExtensionSimpleContent

```go
type ExtensionSimpleContent struct {
}
```


#### func (*ExtensionSimpleContent) Parent

```go
func (me *ExtensionSimpleContent) Parent() element
```

#### type Field

```go
type Field struct {
}
```


#### func (*Field) Parent

```go
func (me *Field) Parent() element
```

#### type Group

```go
type Group struct {
}
```


#### func (*Group) Parent

```go
func (me *Group) Parent() element
```

#### func (*Group) Value

```go
func (me *Group) Value() (l xsdt.Long)
```

#### type Import

```go
type Import struct {
}
```


#### func (*Import) Parent

```go
func (me *Import) Parent() element
```

#### type Include

```go
type Include struct {
}
```


#### func (*Include) Parent

```go
func (me *Include) Parent() element
```

#### type Key

```go
type Key struct {
}
```


#### func (*Key) Parent

```go
func (me *Key) Parent() element
```

#### type KeyRef

```go
type KeyRef struct {
}
```


#### func (*KeyRef) Parent

```go
func (me *KeyRef) Parent() element
```

#### type List

```go
type List struct {
}
```


#### func (*List) Parent

```go
func (me *List) Parent() element
```

#### type Notation

```go
type Notation struct {
}
```


#### func (*Notation) Parent

```go
func (me *Notation) Parent() element
```

#### type PkgBag

```go
type PkgBag struct {
	Schema *Schema
	Stacks pkgStacks
}
```


#### func (*PkgBag) AnonName

```go
func (me *PkgBag) AnonName(n string) (an xsdt.NCName)
```

#### type Redefine

```go
type Redefine struct {
}
```


#### func (*Redefine) Parent

```go
func (me *Redefine) Parent() element
```

#### type RestrictionComplexContent

```go
type RestrictionComplexContent struct {
}
```


#### func (*RestrictionComplexContent) Parent

```go
func (me *RestrictionComplexContent) Parent() element
```

#### type RestrictionSimpleContent

```go
type RestrictionSimpleContent struct {
}
```


#### func (*RestrictionSimpleContent) Parent

```go
func (me *RestrictionSimpleContent) Parent() element
```

#### type RestrictionSimpleEnumeration

```go
type RestrictionSimpleEnumeration struct {
}
```


#### func (*RestrictionSimpleEnumeration) Parent

```go
func (me *RestrictionSimpleEnumeration) Parent() element
```

#### type RestrictionSimpleFractionDigits

```go
type RestrictionSimpleFractionDigits struct {
}
```


#### func (*RestrictionSimpleFractionDigits) Parent

```go
func (me *RestrictionSimpleFractionDigits) Parent() element
```

#### type RestrictionSimpleLength

```go
type RestrictionSimpleLength struct {
}
```


#### func (*RestrictionSimpleLength) Parent

```go
func (me *RestrictionSimpleLength) Parent() element
```

#### type RestrictionSimpleMaxExclusive

```go
type RestrictionSimpleMaxExclusive struct {
}
```


#### func (*RestrictionSimpleMaxExclusive) Parent

```go
func (me *RestrictionSimpleMaxExclusive) Parent() element
```

#### type RestrictionSimpleMaxInclusive

```go
type RestrictionSimpleMaxInclusive struct {
}
```


#### func (*RestrictionSimpleMaxInclusive) Parent

```go
func (me *RestrictionSimpleMaxInclusive) Parent() element
```

#### type RestrictionSimpleMaxLength

```go
type RestrictionSimpleMaxLength struct {
}
```


#### func (*RestrictionSimpleMaxLength) Parent

```go
func (me *RestrictionSimpleMaxLength) Parent() element
```

#### type RestrictionSimpleMinExclusive

```go
type RestrictionSimpleMinExclusive struct {
}
```


#### func (*RestrictionSimpleMinExclusive) Parent

```go
func (me *RestrictionSimpleMinExclusive) Parent() element
```

#### type RestrictionSimpleMinInclusive

```go
type RestrictionSimpleMinInclusive struct {
}
```


#### func (*RestrictionSimpleMinInclusive) Parent

```go
func (me *RestrictionSimpleMinInclusive) Parent() element
```

#### type RestrictionSimpleMinLength

```go
type RestrictionSimpleMinLength struct {
}
```


#### func (*RestrictionSimpleMinLength) Parent

```go
func (me *RestrictionSimpleMinLength) Parent() element
```

#### type RestrictionSimplePattern

```go
type RestrictionSimplePattern struct {
}
```


#### func (*RestrictionSimplePattern) Parent

```go
func (me *RestrictionSimplePattern) Parent() element
```

#### type RestrictionSimpleTotalDigits

```go
type RestrictionSimpleTotalDigits struct {
}
```


#### func (*RestrictionSimpleTotalDigits) Parent

```go
func (me *RestrictionSimpleTotalDigits) Parent() element
```

#### type RestrictionSimpleType

```go
type RestrictionSimpleType struct {
}
```


#### func (*RestrictionSimpleType) Parent

```go
func (me *RestrictionSimpleType) Parent() element
```

#### type RestrictionSimpleWhiteSpace

```go
type RestrictionSimpleWhiteSpace struct {
}
```


#### func (*RestrictionSimpleWhiteSpace) Parent

```go
func (me *RestrictionSimpleWhiteSpace) Parent() element
```

#### type Schema

```go
type Schema struct {
	XMLName            xml.Name          `xml:"schema"`
	XMLNamespacePrefix string            `xml:"-"`
	XMLNamespaces      map[string]string `xml:"-"`
	XMLIncludedSchemas []*Schema         `xml:"-"`
	XSDNamespacePrefix string            `xml:"-"`
	XSDParentSchema    *Schema           `xml:"-"`
}
```


#### func  LoadSchema

```go
func LoadSchema(uri string, localCopy bool) (sd *Schema, err error)
```

#### func (*Schema) MakeGoPkgSrcFile

```go
func (me *Schema) MakeGoPkgSrcFile() (goOutFilePath string, err error)
```

#### func (*Schema) Parent

```go
func (me *Schema) Parent() element
```

#### func (*Schema) RootSchema

```go
func (me *Schema) RootSchema() *Schema
```

#### type Selector

```go
type Selector struct {
}
```


#### func (*Selector) Parent

```go
func (me *Selector) Parent() element
```

#### type Sequence

```go
type Sequence struct {
}
```


#### func (*Sequence) Parent

```go
func (me *Sequence) Parent() element
```

#### func (*Sequence) Value

```go
func (me *Sequence) Value() (l xsdt.Long)
```

#### type SimpleContent

```go
type SimpleContent struct {
}
```


#### func (*SimpleContent) Parent

```go
func (me *SimpleContent) Parent() element
```

#### type SimpleType

```go
type SimpleType struct {
}
```


#### func (*SimpleType) Parent

```go
func (me *SimpleType) Parent() element
```

#### type Union

```go
type Union struct {
}
```


#### func (*Union) Parent

```go
func (me *Union) Parent() element
```

#### type Unique

```go
type Unique struct {
}
```


#### func (*Unique) Parent

```go
func (me *Unique) Parent() element
```

--
**godocdown** http://github.com/robertkrimen/godocdown