package uuid

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

var testUUID = UUID{0x6a, 0x2b, 0x5c, 0x4d, 0x3e, 0x2f, 0x1a, 0x0b, 0x0c, 0x1d, 0x2e, 0x3f, 0x4a, 0x5b, 0x6c, 0x7d}

func TestJSON(t *testing.T) {
	type S struct {
		ID1 UUID
		ID2 UUID
	}
	s1 := S{ID1: testUUID}
	data, err := json.Marshal(&s1)
	if err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := json.Unmarshal(data, &s2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(&s1, &s2) {
		t.Errorf("got %#v, want %#v", s2, s1)
	}
}

func TestMarshal(t *testing.T) {
	_, err := testUUID.MarshalText()
	if err != nil {
		t.Fatal(err)
	}

	b, err := testUUID.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(testUUID[:], b) {
		t.Fatal("MarshalBinary failed - not equal")
	}
}

func TestUnmarshal(t *testing.T) {
	var uuid UUID
	err := uuid.UnmarshalBinary(testUUID[:])
	if err != nil {
		t.Fatal(err)
	}

	if !uuid.Equal(testUUID) {
		t.Fatal("UnmarshalBinary failed - not equal")
	}

	err = uuid.UnmarshalBinary(testUUID[:8])
	if err != ErrInvalidFormat {
		t.Fatal("UnmarshalBinary failed - should have failed")
	}
}
