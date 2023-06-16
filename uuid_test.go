package uuid

import (
	"fmt"
	"testing"
)

// Test for Equal() method
func ExampleUUID_Equal() {
	var id UUID
	var id2 UUID
	var id3 UUID

	id = New()
	id2 = New()

	if id.Equal(id2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	if id3.Equal(Nil) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	// Output: Not Equal
	// Equal
}

// Test for Parse method
func ExampleParse() {
	str := "12345678-1234-1234-aBcD-123456789012"

	id, err := Parse(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	// Output: 12345678-1234-1234-abcd-123456789012
}

func TestNewTimeBased(t *testing.T) {
	id1 := NewTimeBased()
	id2 := NewTimeBased()
	if id1.Equal(id2) {
		t.Error("NewTimeBased() returned non-unique UUIDs")
	}
}

func TestParse(t *testing.T) {
	var good []string = []string{
		"12345678-1234-1234-aBcD-123456789012",
		"{12345678-1234-1234-aBcD-123456789012}",
		"12345678-1234-1234-ABCD-123456789012",
		"{12345678-1234-1234-ABCD-123456789012}",
	}

	var goodUUID []UUID = []UUID{
		{0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x12, 0x34, 0xab, 0xcd, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12},
		{0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x12, 0x34, 0xab, 0xcd, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12},
		{0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x12, 0x34, 0xab, 0xcd, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12},
		{0x12, 0x34, 0x56, 0x78, 0x12, 0x34, 0x12, 0x34, 0xab, 0xcd, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12},
	}

	var bad []string = []string{
		"12345678-1234-1234-aBcD-1234567890123",
		"{12345678-1234-1234-aBcD-1234567890123}",
		"{1234567812341234aBcD1234567890123}",
		"12345678-1234-1234-ABCD!123456789012",
		"{12345678-1234-1234-ABCD!123456789012}",
		"12345678-1234-1234-aBcD-12345678901",
		"{12345678-1234-1234-aBcD-12345678901}",
		"12345678-1234-1234!ABCD-123456789012",
		"{12345678-1234-1234!ABCD-123456789012}",
		"12345678-1234!1234-aBcD-123456789012",
		"{12345678-1234!1234-aBcD-123456789012}",
		"12345678!1234-1234-aBcD-123456789012",
		"{12345678!1234-1234-aBcD-123456789012}",
		"#12345678-1234-1234-aBcD-123456789012#",
		"123456X8-123M-1234-aBcD-123456789012",
		"12345678-1X34-1234-aBcD-123456789012",
		"12345678-1234-12X4-aBcD-123456789012",
		"12345678-1234-1234-aBcX-123456789012",
		"12345678-1234-1234-aBcD-12345678X012",
	}

	for i, s := range good {
		u, err := Parse(s)
		if err != nil {
			t.Errorf("Parse(%s) returned error: %s", s, err)
		}
		if !u.Equal(goodUUID[i]) {
			t.Errorf("Parse(%s) returned %s, expected %s", s, u, goodUUID[i])
		}
	}

	for _, s := range bad {
		u, err := Parse(s)
		if err == nil {
			t.Errorf("Parse(%s) returned %s, expected error", s, u)
		}
	}
}
