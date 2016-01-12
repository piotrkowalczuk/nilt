package nilt

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"strconv"

	"github.com/golang/protobuf/proto"
)

// String represents a string that may be nil.
type String struct {
	String string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *String) Reset() { *m = String{} }

//func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage() {}

// StringOr returns given string value if receiver is nil or invalid.
func (s *String) StringOr(or string) string {
	if s == nil {
		return or
	}
	if !s.Valid {
		return or
	}

	return s.String
}

// MarshalJSON implements json.Marshaler interface.
func (ns *String) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return nil, nil
	}

	return json.Marshal(ns.String)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ns *String) UnmarshalJSON(data []byte) error {
	if data == nil {
		ns.String, ns.Valid = "", false
		return nil
	}

	ns.Valid = true

	return json.Unmarshal(data, &ns.String)
}

// Value implements the driver Valuer interface.
func (n String) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.String, s.Valid = "", false
		return nil
	}
	s.Valid = true

	switch v := value.(type) {
	case []byte:
		s.String = string(v)
	case string:
		s.String = v
	default:
		return fmt.Errorf("nilt: unsuported type (%T) passed to String.Scan", value)
	}

	return nil
}

func (s *String) Appear() bool {
	return s != nil && s.Valid
}

// Int64 represents a int64 that may be nil.
type Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (ni *Int64) Reset() { *ni = Int64{} }

// String implements proto.Message interface.
func (ni *Int64) String() string { return proto.CompactTextString(ni) }

// ProtoMessage implements proto.Message interface.
func (*Int64) ProtoMessage() {}

// Int64Or returns given int64 value if receiver is nil or invalid.
func (i *Int64) Int64Or(or int64) int64 {
	if i == nil {
		return or
	}
	if !i.Valid {
		return or
	}

	return i.Int64
}

// Value implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

// Scan implements the Scanner interface.
func (i *Int64) Scan(value interface{}) (err error) {
	if value == nil {
		i.Int64, i.Valid = 0, false
		return nil
	}
	i.Valid = true

	switch v := value.(type) {
	case []byte:
		i.Int64, err = strconv.ParseInt(string(v), 10, 64)
	case string:
		i.Int64, err = strconv.ParseInt(v, 10, 64)
	case int64:
		i.Int64 = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Int64.Scan", value)
	}

	return
}

func (i *Int64) Appear() bool {
	return i != nil && i.Valid
}

// Int32 represents a int32 that may be nil.
type Int32 struct {
	Int32 int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (ni *Int32) Reset() { *ni = Int32{} }

// String implements proto.Message interface.
func (ni *Int32) String() string { return proto.CompactTextString(ni) }

// ProtoMessage implements proto.Message interface.
func (*Int32) ProtoMessage() {}

// Int32Or returns given int32 value if receiver is nil or invalid.
func (i *Int32) Int32Or(or int32) int32 {
	if i == nil {
		return or
	}
	if !i.Valid {
		return or
	}

	return i.Int32
}

// Value implements the driver Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int32, nil
}

// Scan implements the Scanner interface.
func (i *Int32) Scan(value interface{}) (err error) {
	if value == nil {
		i.Int32, i.Valid = 0, false
		return nil
	}
	i.Valid = true

	var tmp int64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseInt(string(v), 10, 32)
		i.Int32 = int32(tmp)
	case string:
		tmp, err = strconv.ParseInt(v, 10, 32)
		i.Int32 = int32(tmp)
	case int32:
		i.Int32 = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Int32.Scan", value)
	}

	return
}

func (i *Int32) Appear() bool {
	return i != nil && i.Valid
}

// Int represents a int that may be nil.
type Int struct {
	Int   int  `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (i *Int) Reset() { *i = Int{} }

// String implements proto.Message interface.
func (i *Int) String() string { return proto.CompactTextString(i) }

// ProtoMessage implements proto.Message interface.
func (*Int) ProtoMessage() {}

// IntOr returns given int value if receiver is nil or invalid.
func (i *Int) IntOr(or int) int {
	if i == nil {
		return or
	}
	if !i.Valid {
		return or
	}

	return i.Int
}

// Value implements the driver Valuer interface.
func (i Int) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int, nil
}

// Scan implements the Scanner interface.
func (i *Int) Scan(value interface{}) (err error) {
	if value == nil {
		i.Int, i.Valid = 0, false
		return nil
	}
	i.Valid = true

	var tmp int64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseInt(string(v), 10, 32)
		i.Int = int(tmp)
	case string:
		tmp, err = strconv.ParseInt(v, 10, 32)
		i.Int = int(tmp)
	case int:
		i.Int = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Int.Scan", value)
	}

	return
}

func (i *Int) Appear() bool {
	return i != nil && i.Valid
}

// Uint32 represents a uint32 that may be nil.
type Uint32 struct {
	Uint32 uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (u *Uint32) Reset() { *u = Uint32{} }

// String implements proto.Message interface.
func (u *Uint32) String() string { return proto.CompactTextString(u) }

// ProtoMessage implements proto.Message interface.
func (*Uint32) ProtoMessage() {}

// Uint32Or returns given uint32 value if receiver is nil or invalid.
func (u *Uint32) Uint32Or(or uint32) uint32 {
	if u == nil {
		return or
	}
	if !u.Valid {
		return or
	}

	return u.Uint32
}

// Value implements the driver Valuer interface.
func (u Uint32) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	return int64(u.Uint32), nil
}

// Scan implements the Scanner interface.
func (u *Uint32) Scan(value interface{}) (err error) {
	if value == nil {
		u.Uint32, u.Valid = 0, false
		return nil
	}
	u.Valid = true

	var tmp uint64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseUint(string(v), 10, 32)
		if tmp > 4294967295 {
			return errors.New("nilt: value passed to Uint32.Scan is out of range")
		}
		u.Uint32 = uint32(tmp)
	case string:
		tmp, err = strconv.ParseUint(v, 10, 32)
		if tmp > 4294967295 {
			return errors.New("nilt: value passed to Uint32.Scan is out of range")
		}
		u.Uint32 = uint32(tmp)
	case int64:
		u.Uint32 = uint32(v)
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Uint32.Scan", value)
	}

	return
}

func (u *Uint32) Appear() bool {
	return u != nil && u.Valid
}

// Float32 represents a flaot64 that may be nil.
type Float32 struct {
	Float32 float32 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (f *Float32) Reset() { *f = Float32{} }

// String implements proto.Message interface.
func (f *Float32) String() string { return proto.CompactTextString(f) }

// ProtoMessage implements proto.Message interface.
func (*Float32) ProtoMessage() {}

// Float32Or returns given Float32 value if receiver is nil or invalid.
func (f *Float32) Float32Or(or float32) float32 {
	if f == nil {
		return or
	}
	if !f.Valid {
		return or
	}

	return f.Float32
}

// Value implements the driver Valuer interface.
func (f Float32) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float32, nil
}

// Scan implements the Scanner interface.
func (f *Float32) Scan(value interface{}) (err error) {
	if value == nil {
		f.Float32, f.Valid = 0.0, false
		return nil
	}
	f.Valid = true

	var tmp float64
	switch v := value.(type) {
	case []byte:
		tmp, err = strconv.ParseFloat(string(v), 32)
		f.Float32 = float32(tmp)
	case string:
		tmp, err = strconv.ParseFloat(v, 32)
		f.Float32 = float32(tmp)
	case float32:
		f.Float32 = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Float32.Scan", value)
	}

	return
}

func (f *Float32) Appear() bool {
	return f != nil && f.Valid
}

// Float64 represents a flaot64 that may be nil.
type Float64 struct {
	Float64 float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (f *Float64) Reset() { *f = Float64{} }

// String implements proto.Message interface.
func (f *Float64) String() string { return proto.CompactTextString(f) }

// ProtoMessage implements proto.Message interface.
func (*Float64) ProtoMessage() {}

// Float64Or returns given float64 value if receiver is nil or invalid.
func (f *Float64) Float64Or(or float64) float64 {
	if f == nil {
		return or
	}
	if !f.Valid {
		return or
	}

	return f.Float64
}

// Value implements the driver Valuer interface.
func (f Float64) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float64, nil
}

// Scan implements the Scanner interface.
func (f *Float64) Scan(value interface{}) (err error) {
	if value == nil {
		f.Float64, f.Valid = 0.0, false
		return nil
	}
	f.Valid = true

	switch v := value.(type) {
	case []byte:
		f.Float64, err = strconv.ParseFloat(string(v), 64)
	case string:
		f.Float64, err = strconv.ParseFloat(v, 64)
	case float64:
		f.Float64 = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Float64.Scan", value)
	}

	return
}

func (f *Float64) Appear() bool {
	return f != nil && f.Valid
}

// Bool represents a bool that may be nil.
type Bool struct {
	Bool  bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (b *Bool) Reset() { *b = Bool{} }

// String implements proto.Message interface.
func (b *Bool) String() string { return proto.CompactTextString(b) }

// ProtoMessage implements proto.Message interface.
func (*Bool) ProtoMessage() {}

// BoolOr returns given bool value if receiver is nil or invalid.
func (b *Bool) BoolOr(or bool) bool {
	if b == nil {
		return or
	}
	if !b.Valid {
		return or
	}

	return b.Bool
}

// Value implements the driver Valuer interface.
func (f Bool) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Bool, nil
}

// Scan implements the Scanner interface.
func (b *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		b.Bool, b.Valid = false, false
		return nil
	}
	b.Valid = true

	switch v := value.(type) {
	case []byte:
		b.Bool, err = strconv.ParseBool(string(v))
	case string:
		b.Bool, err = strconv.ParseBool(v)
	case bool:
		b.Bool = v
	default:
		err = fmt.Errorf("nilt: unsuported type (%T) passed to Bool.Scan", value)
	}

	return
}

func (b *Bool) Appear() bool {
	return b != nil && b.Valid
}
