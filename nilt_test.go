package nilt_test

import (
	"testing"

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
			t.Errorf("booleans are not equal expected %d, got %d", given.Valid, tmp.Valid)
		}
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
		t.Error("wrong valid property value, got %b but expected %b", u.Valid, expectedValid)
	}

	if u.Uint32 != expectedValue {
		t.Error("wrong uint32 property value, got %d but expected %d", u.Uint32, expectedValue)
	}
}
