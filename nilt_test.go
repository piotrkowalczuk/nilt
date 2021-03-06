package nilt_test

import (
	"testing"

	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/piotrkowalczuk/nilt"
)

func TestInt64_ProtoMessage(t *testing.T) {
	var (
		buf []byte
		err error
		tmp nilt.Int64
	)
	success := []nilt.Int64{
		{Int64: 1, Valid: true},
		{Int64: 0, Valid: false},
		{Int64: 13123, Valid: false},
		{Int64: -1241223, Valid: true},
	}

	for _, given := range success {
		buf, err = proto.Marshal(&given)
		if err != nil {
			t.Errorf("marshal returned unexpected error: %s", err.Error())
			continue
		}

		err = proto.Unmarshal(buf, &tmp)
		if err != nil {
			t.Errorf("unmarshal returned unexpected error: %s", err.Error())
			continue
		}

		if tmp.Int64 != given.Int64 {
			t.Errorf("integers are not equal expected %d, got %d", given.Int64, tmp.Int64)
		}

		if tmp.Valid != given.Valid {
			t.Errorf("booleans are not equal expected %t, got %t", given.Valid, tmp.Valid)
		}
	}
}

func TestInt64_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *nilt.Int64
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &nilt.Int64{},
			expected: "null",
		},
		"valid": {
			given:    &nilt.Int64{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &nilt.Int64{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &nilt.Int64{Int64: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &nilt.Int64{Int64: 9223372036854775807, Valid: true},
			expected: "9223372036854775807",
		},
		"non zero valid min value": {
			given:    &nilt.Int64{Int64: -9223372036854775808, Valid: true},
			expected: "-9223372036854775808",
		},
		"non zero invalid max value": {
			given:    &nilt.Int64{Int64: 9223372036854775807, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &nilt.Int64{Int64: -9223372036854775808, Valid: false},
			expected: "null",
		},
	}

CasesLoop:
	for d, c := range cases {
		b, err := json.Marshal(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", d, err.Error())
			continue CasesLoop
		}

		if string(b) != c.expected {
			t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
		} else {
			t.Logf("%s: %s", d, string(b))
		}
	}

	type within struct {
		ID *nilt.Int64 `json:"id"`
	}
}

func TestUint32_Scan(t *testing.T) {
	testUint32_Scan_success(t, nil, 0, false)
}

func TestUint32_Scan_string(t *testing.T) {
	success := map[uint32]string{
		100:        "100",
		4294967295: "4294967295",
		0:          "0",
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func TestUint32_Scan_int64(t *testing.T) {
	success := map[uint32]int64{
		100:        100,
		4294967295: 4294967295,
		0:          0,
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func TestUint32_Scan_bytes(t *testing.T) {
	success := map[uint32][]byte{
		100:        []byte("100"),
		4294967295: []byte("4294967295"),
		0:          []byte("0"),
	}

	for expected, given := range success {
		testUint32_Scan_success(t, given, expected, true)
	}
}

func testUint32_Scan_success(t *testing.T, given interface{}, expectedValue uint32, expectedValid bool) {
	var u nilt.Uint32
	if err := u.Scan(given); err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}

	if u.Valid != expectedValid {
		t.Errorf("wrong valid property value, got %t but expected %t", u.Valid, expectedValid)
	}

	if u.Uint32 != expectedValue {
		t.Errorf("wrong uint32 property value, got %d but expected %d", u.Uint32, expectedValue)
	}
}

func TestUint32_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		given    *nilt.Uint32
		expected string
	}{

		"nil": {
			given:    nil,
			expected: "null",
		},
		"zero value": {
			given:    &nilt.Uint32{},
			expected: "null",
		},
		"valid": {
			given:    &nilt.Uint32{Valid: true},
			expected: "0",
		},
		"invalid": {
			given:    &nilt.Uint32{Valid: false},
			expected: "null",
		},
		"non zero valid value": {
			given:    &nilt.Uint32{Uint32: 123, Valid: true},
			expected: "123",
		},
		"non zero valid max value": {
			given:    &nilt.Uint32{Uint32: 4294967295, Valid: true},
			expected: "4294967295",
		},
		"non zero valid min value": {
			given:    &nilt.Uint32{Uint32: 0, Valid: true},
			expected: "0",
		},
		"non zero invalid max value": {
			given:    &nilt.Uint32{Uint32: 4294967295, Valid: false},
			expected: "null",
		},
		"non zero invalid min value": {
			given:    &nilt.Uint32{Uint32: 0, Valid: false},
			expected: "null",
		},
	}

CasesLoop:
	for d, c := range cases {
		b, err := json.Marshal(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", d, err.Error())
			continue CasesLoop
		}

		if string(b) != c.expected {
			t.Errorf("%s: wrong output, expected %s but got %s", d, c.expected, string(b))
		} else {
			t.Logf("%s: %s", d, string(b))
		}
	}

	type within struct {
		ID *nilt.Uint32 `json:"id"`
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	cases := map[string]*nilt.Bool{
		"nil":         nil,
		"zero value":  &nilt.Bool{},
		"valid":       &nilt.Bool{Valid: true},
		"invalid":     &nilt.Bool{Valid: false},
		"true true":   &nilt.Bool{Bool: true, Valid: true},
		"true false":  &nilt.Bool{Bool: true, Valid: false},
		"false false": &nilt.Bool{Bool: false, Valid: false},
		"false true":  &nilt.Bool{Bool: false, Valid: true},
	}

SimpleLoop:
	for d, c := range cases {
		b, err := json.Marshal(c)
		if err != nil {
			t.Errorf("simple: %s: unexpected error: %s", d, err.Error())
			continue SimpleLoop
		}

		t.Logf("simple: %s: %s", d, string(b))
	}

	type within struct {
		Exists *nilt.Bool `json:"exists"`
	}

WithinLoop:
	for d, c := range cases {
		w := within{Exists: c}
		b, err := json.Marshal(w)
		if err != nil {
			t.Errorf("within: %s: unexpected error: %s", d, err.Error())
			continue WithinLoop
		}

		t.Logf("within: %s: %s", d, string(b))
	}
}
