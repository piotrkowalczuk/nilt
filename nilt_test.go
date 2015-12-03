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
