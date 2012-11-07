package goxsdpkg

//	XML Schema for RSS v2.0 feed files.
//	Project home: http://www.codeplex.com/rss2schema/
//	Based on the RSS 2.0 specification document at http://cyber.law.harvard.edu/rss/rss.html
//	Author: Jorgen Thelin
//	Revision: 16
//	Date: 01-Nov-2008
//	Feedback to: http://www.codeplex.com/rss2schema/WorkItem/List.aspx
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

//	A time in GMT when aggregators should not request the channel data. The hour beginning at midnight is hour zero.
type SkipHour xs:nonNegativeInteger

//	A day when aggregators should not request the channel data.
type SkipDay xs:string

//	The height of the image in pixels.
type ImageHeight xs:positiveInteger

//	The width of the image in pixels.
type ImageWidth xs:positiveInteger

type CloudProtocol xs:string

//	Using the regexp definiton of E-Mail Address by Lucadean from the .NET RegExp Pattern Repository at http://www.3leaf.com/default/NetRegExpRepository.aspx
type EmailAddress xs:string

//	A date-time displayed in RFC-822 format.
//	Using the regexp definiton of rfc-822 date by Sam Ruby at http://www.intertwingly.net/blog/1360.html
type Rfc822FormatDate xs:string
