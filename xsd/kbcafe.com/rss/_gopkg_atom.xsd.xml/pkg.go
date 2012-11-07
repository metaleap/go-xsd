package goxsdpkg

//	This version of the Atom schema is based on version 1.0 of the format specifications,
//	found here http://www.atomenabled.org/developers/syndication/atom-format-spec.php.
//	An Atom document may have two root elements, feed and entry, as defined in section 2.
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

//	Schema definition for an email address.
type EmailType xs:normalizedString
