package nilt

import (
	"database/sql/driver"
	"encoding/json"
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
