package uuid

import (
	"testing"
)

func TestValue(t *testing.T) {
	test := "12345678-1234-1234-abcd-123456789012"
	id, _ := Parse(test)
	val, _ := id.Value()
	if test != val.(string) {
		t.Errorf("Value() did not return expected string (got %s, expected %s)", val, test)
	}
}

func TestScan(t *testing.T) {
	stringTest := "12345678-1234-1234-abcd-123456789012"
	invalidTest := "invalid uuid"

	// good type tests
	var uuid UUID
	err := (&uuid).Scan(stringTest)
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan([]byte(stringTest))
	if err != nil {
		t.Fatal(err)
	}

	byteTest := make([]byte, len(uuid))
	copy(byteTest, uuid[:])
	err = (&uuid).Scan(byteTest)
	if err != nil {
		t.Fatal(err)
	}

	// bad tests
	err = (&uuid).Scan(1)
	if err == nil {
		t.Error("incorrectly parsed int")
	}
	if err != ErrUnableToScanType {
		t.Errorf("incorrect error returned when parsing int: %s", err.Error())
	}

	err = (&uuid).Scan(invalidTest)
	if err == nil {
		t.Error("incorrectly parsed invalid uuid string")
	}
	if err != ErrInvalidFormat {
		t.Errorf("incorrect error returned when parsing int: %s", err.Error())
	}

	var emptySlice []byte
	err = (&uuid).Scan(emptySlice)
	if err == nil {
		t.Error("incorrecly parsed empty bytes slice")
	}
	if err != ErrInvalidFormat {
		t.Errorf("incorrect error returned when parsing int: %s", err.Error())
	}

	err = (&uuid).Scan(byteTest[:len(byteTest)-1])
	if err == nil {
		t.Error("incorrectly parsed invalid bytes slice")
	}
	if err != ErrInvalidFormat {
		t.Errorf("incorrect error returned when parsing int: %s", err.Error())
	}
}
