package uuid

import (
	"testing"
)

func TestString(t *testing.T) {
	var id UUID
	if id.String() != "00000000-0000-0000-0000-000000000000" {
		t.Errorf("String() returned unexpected value: %s", id.String())
	}

	var id2 UUID = [16]byte{0x12, 0x34, 0x56, 0x78}
	if id2.String() != "12345678-0000-0000-0000-000000000000" {
		t.Errorf("String() returned unexpected value: %s", id2.String())
	}
}
