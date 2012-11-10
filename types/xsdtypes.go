package xsdt

import (
	"strconv"
)

type AnyURI string

	func (me *AnyURI) SetFromString (v string) {
		*me = AnyURI(v)
	}

	func (me AnyURI) String () string {
		return string(me)
	}

type Base64Binary string // []byte

	func (me *Base64Binary) SetFromString (v string) {
		*me = Base64Binary(v)
	}

	func (me Base64Binary) String () string {
		return string(me)
	}

type Boolean bool

	func (me *Boolean) SetFromString (v string) {
		b, _ := strconv.ParseBool(v); *me = Boolean(b)
	}

	func (me Boolean) String () string {
		return strconv.FormatBool(bool(me))
	}

type Byte byte

	func (me *Byte) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 8); *me = Byte(v)
	}

	func (me Byte) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type Date string // time.Time

	func (me *Date) SetFromString (v string) {
		*me = Date(v)
	}

	func (me Date) String () string {
		return string(me)
	}

type DateTime string // time.Time

	func (me *DateTime) SetFromString (v string) {
		*me = DateTime(v)
	}

	func (me DateTime) String () string {
		return string(me)
	}

type Decimal string // complex128

	func (me *Decimal) SetFromString (v string) {
		*me = Decimal(v)
	}

	func (me Decimal) String () string {
		return string(me)
	}

type Double float64

	func (me *Double) SetFromString (s string) {
		v, _ := strconv.ParseFloat(s, 64); *me = Double(v)
	}

	func (me Double) String () string {
		return strconv.FormatFloat(float64(me), 'f', 8, 64)
	}

type Duration string // time.Duration

	func (me *Duration) SetFromString (v string) {
		*me = Duration(v)
	}

	func (me Duration) String () string {
		return string(me)
	}

type Entities string

	func (me *Entities) SetFromString (v string) {
		*me = Entities(v)
	}

	func (me Entities) String () string {
		return string(me)
	}

	func (me Entities) Values () (list []Entity) {
		var btv = new(Entity)
		var spl = ListValues(string(me))
		list = make([]Entity, len(spl))
		for i, s := range spl { btv.SetFromString(s); list[i] = *btv }
		return
	}

type Entity string

	func (me *Entity) SetFromString (v string) {
		*me = Entity(v)
	}

	func (me Entity) String () string {
		return string(me)
	}

type Float float32

	func (me *Float) SetFromString (s string) {
		v, _ := strconv.ParseFloat(s, 32); *me = Float(v)
	}

	func (me Float) String () string {
		return strconv.FormatFloat(float64(me), 'f', 8, 32)
	}

type GDay string

	func (me *GDay) SetFromString (v string) {
		*me = GDay(v)
	}

	func (me GDay) String () string {
		return string(me)
	}

type GYear string

	func (me *GYear) SetFromString (v string) {
		*me = GYear(v)
	}

	func (me GYear) String () string {
		return string(me)
	}

type GYearMonth string

	func (me *GYearMonth) SetFromString (v string) {
		*me = GYearMonth(v)
	}

	func (me GYearMonth) String () string {
		return string(me)
	}

type GMonth string

	func (me *GMonth) SetFromString (v string) {
		*me = GMonth(v)
	}

	func (me GMonth) String () string {
		return string(me)
	}

type GMonthDay string

	func (me *GMonthDay) SetFromString (v string) {
		*me = GMonthDay(v)
	}

	func (me GMonthDay) String () string {
		return string(me)
	}

type HexBinary string // []byte

	func (me *HexBinary) SetFromString (v string) {
		*me = HexBinary(v)
	}

	func (me HexBinary) String () string {
		return string(me)
	}

type Id string

	func (me *Id) SetFromString (v string) {
		*me = Id(v)
	}

	func (me Id) String () string {
		return string(me)
	}

type Idref string

	func (me *Idref) SetFromString (v string) {
		*me = Idref(v)
	}

	func (me Idref) String () string {
		return string(me)
	}

type Idrefs string

	func (me *Idrefs) SetFromString (v string) {
		*me = Idrefs(v)
	}

	func (me Idrefs) String () string {
		return string(me)
	}

	func (me Idrefs) Values () (list []Idref) {
		var btv = new(Idref)
		var spl = ListValues(string(me))
		list = make([]Idref, len(spl))
		for i, s := range spl { btv.SetFromString(s); list[i] = *btv }
		return
	}

type Int int32

	func (me *Int) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 32); *me = Int(v)
	}

	func (me Int) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type Integer int

	func (me *Integer) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 64); *me = Integer(v)
	}

	func (me Integer) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type Language string

	func (me *Language) SetFromString (v string) {
		*me = Language(v)
	}

	func (me Language) String () string {
		return string(me)
	}

type Long int64

	func (me *Long) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 64); *me = Long(v)
	}

	func (me Long) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type Name string

	func (me *Name) SetFromString (v string) {
		*me = Name(v)
	}

	func (me Name) String () string {
		return string(me)
	}

type NCName string

	func (me *NCName) SetFromString (v string) {
		*me = NCName(v)
	}

	func (me NCName) String () string {
		return string(me)
	}

type NegativeInteger int

	func (me *NegativeInteger) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 64); *me = NegativeInteger(v)
	}

	func (me NegativeInteger) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type Nmtoken string

	func (me *Nmtoken) SetFromString (v string) {
		*me = Nmtoken(v)
	}

	func (me Nmtoken) String () string {
		return string(me)
	}

type Nmtokens string

	func (me *Nmtokens) SetFromString (v string) {
		*me = Nmtokens(v)
	}

	func (me Nmtokens) String () string {
		return string(me)
	}

	func (me Nmtokens) Values () (list []Nmtoken) {
		var btv = new(Nmtoken)
		var spl = ListValues(string(me))
		list = make([]Nmtoken, len(spl))
		for i, s := range spl { btv.SetFromString(s); list[i] = *btv }
		return
	}

type NonNegativeInteger uint

	func (me *NonNegativeInteger) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 64); *me = NonNegativeInteger(v)
	}

	func (me NonNegativeInteger) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type NonPositiveInteger int

	func (me *NonPositiveInteger) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 64); *me = NonPositiveInteger(v)
	}

	func (me NonPositiveInteger) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type NormalizedString string

	func (me *NormalizedString) SetFromString (v string) {
		*me = NormalizedString(v)
	}

	func (me NormalizedString) String () string {
		return string(me)
	}

type notation struct {
	Id, Name, Public, System string
}

type Notation string

	func (me *Notation) SetFromString (v string) {
		*me = Notation(v)
	}

	func (me Notation) String () string {
		return string(me)
	}

type Notations map[string]*notation

	func (me Notations) Add (id, name, public, system string) {
		me[name] = &notation { Id: id, Name: name, Public: public, System: system }
	}

type PositiveInteger uint

	func (me *PositiveInteger) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 64); *me = PositiveInteger(v)
	}

	func (me PositiveInteger) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type Qname string

	func (me *Qname) SetFromString (v string) {
		*me = Qname(v)
	}

	func (me Qname) String () string {
		return string(me)
	}

type Short int16

	func (me *Short) SetFromString (s string) {
		v, _ := strconv.ParseInt(s, 0, 16); *me = Short(v)
	}

	func (me Short) String () string {
		return strconv.FormatInt(int64(me), 10)
	}

type String string

	func (me *String) SetFromString (v string) {
		*me = String(v)
	}

	func (me String) String () string {
		return string(me)
	}

type Token string

	func (me *Token) SetFromString (v string) {
		*me = Token(v)
	}

	func (me Token) String () string {
		return string(me)
	}

type UnsignedByte uint8

	func (me *UnsignedByte) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 8); *me = UnsignedByte(v)
	}

	func (me UnsignedByte) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type UnsignedInt uint32

	func (me *UnsignedInt) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 32); *me = UnsignedInt(v)
	}

	func (me UnsignedInt) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type UnsignedLong uint64

	func (me *UnsignedLong) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 64); *me = UnsignedLong(v)
	}

	func (me UnsignedLong) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

type UnsignedShort uint16

	func (me *UnsignedShort) SetFromString (s string) {
		v, _ := strconv.ParseUint(s, 0, 16); *me = UnsignedShort(v)
	}

	func (me UnsignedShort) String () string {
		return strconv.FormatUint(uint64(me), 10)
	}

// why import "strings" just for this one need to Split()...
func ListValues (v string) (spl []string) {
	var cur = ""
	for _, r := range v {
		if r == ' ' {
			if len(cur) > 0 { spl = append(spl, cur) }
			cur = ""
		} else {
			cur += string(r)
		}
	}
	return
}
