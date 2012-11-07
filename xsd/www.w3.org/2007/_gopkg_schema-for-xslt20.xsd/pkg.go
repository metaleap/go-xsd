package goxsdpkg

//	This is a schema for XSLT 2.0 stylesheets.
//	It defines all the elements that appear in the XSLT namespace; it also
//	provides hooks that allow the inclusion of user-defined literal result elements,
//	extension instructions, and top-level data elements.
//	The schema is derived (with kind permission) from a schema for XSLT 1.0 stylesheets
//	produced by Asir S Vedamuthu of WebMethods Inc.
//	This schema is available for use under the conditions of the W3C Software License
//	published at http://www.w3.org/Consortium/Legal/copyright-software-19980720
//	The schema is organized as follows:
//	PART A: definitions of complex types and model groups used as the basis
//	for element definitions
//	PART B: definitions of individual XSLT elements
//	PART C: definitions for literal result elements
//	PART D: definitions of simple types used in attribute definitions
//	This schema does not attempt to define all the constraints that apply to a valid
//	XSLT 2.0 stylesheet module. It is the intention that all valid stylesheet modules
//	should conform to this schema; however, the schema is non-normative and in the event
//	of any conflict, the text of the Recommendation takes precedence.
//	This schema does not implement the special rules that apply when a stylesheet
//	has sections that use forwards-compatible-mode. In this mode, setting version="3.0"
//	allows elements from the XSLT namespace to be used that are not defined in XSLT 2.0.
//	Simplified stylesheets (those with a literal result element as the outermost element)
//	will validate against this schema only if validation starts in lax mode.
//	This version is dated 2007-03-16
//	Authors: Michael H Kay, Saxonica Limited
//	Jeni Tennison, Jeni Tennison Consulting Ltd.
//	2007-03-15: added xsl:document element
//	revised xsl:sequence element
//	see http://www.w3.org/Bugs/Public/show_bug.cgi?id=4237
//	PART A: definitions of complex types and model groups used as the basis
//	for element definitions
//	PART B: definitions of individual XSLT elements
//	Elements are listed in alphabetical order.
//	PART C: definition of literal result elements
//	There are three ways to define the literal result elements
//	permissible in a stylesheet.
//	(a) do nothing. This allows any element to be used as a literal
//	result element, provided it is not in the XSLT namespace
//	(b) declare all permitted literal result elements as members
//	of the xsl:literal-result-element substitution group
//	(c) redefine the model group xsl:result-elements to accommodate
//	all permitted literal result elements.
//	Literal result elements are allowed to take certain attributes
//	in the XSLT namespace. These are defined in the attribute group
//	literal-result-element-attributes, which can be included in the
//	definition of any literal result element.
//	PART D: definitions of simple types used in stylesheet attributes
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

//	This type is used for all attributes that allow an attribute value template.
//	The general rules for the syntax of attribute value templates, and the specific
//	rules for each such attribute, are described in the XSLT 2.0 Recommendation.
type Avt xs:string

//	A string containing exactly one character.
type Char xs:string

//	An XPath 2.0 expression.
type Expression xs:token

//	Describes how type annotations in source documents are handled.
type Input-type-annotations-type xs:token

//	The level attribute of xsl:number:
//	one of single, multiple, or any.
type Level xs:NCName

//	The mode attribute of xsl:apply-templates:
//	either a QName, or #current, or #default.

//	The mode attribute of xsl:template:
//	either a list, each member being either a QName or #default;
//	or the value #all

//	A list of NameTests, as defined in the XPath 2.0 Recommendation.
//	Each NameTest is either a QName, or "*", or "prefix:*", or "*:localname"




//	The method attribute of xsl:output:
//	Either one of the recognized names "xml", "xhtml", "html", "text",
//	or a QName that must include a prefix.

//	A match pattern as defined in the XSLT 2.0 Recommendation.
//	The syntax for patterns is a restricted form of the syntax for
//	XPath 2.0 expressions.
type Pattern xsl:expression

//	Either a namespace prefix, or #default.
//	Used in the xsl:namespace-alias element.

//	A list of QNames.
//	Used in the [xsl:]use-attribute-sets attribute of various elements,
//	and in the cdata-section-elements attribute of xsl:output

//	A QName.
//	This schema does not use the built-in type xs:QName, but rather defines its own
//	QName type. Although xs:QName would define the correct validation on these attributes,
//	a schema processor would expand unprefixed QNames incorrectly when constructing the PSVI,
//	because (as defined in XML Schema errata) an unprefixed xs:QName is assumed to be in
//	the default namespace, which is not the correct assumption for XSLT.
//	The data type is defined as a restriction of the built-in type Name, restricted
//	so that it can only contain one colon which must not be the first or last character.
type QName xs:Name

//	The description of a data type, conforming to the
//	SequenceType production defined in the XPath 2.0 Recommendation
type Sequence-type xs:token


//	Describes different ways of type-annotating an element or attribute.
type Validation-strip-or-preserve xsl:validation-type

//	Describes different ways of type-annotating an element or attribute.
type Validation-type xs:token

//	One of the values "yes" or "no".
type Yes-or-no xs:token

//	One of the values "yes" or "no" or "omit".
type Yes-or-no-or-omit xs:token
