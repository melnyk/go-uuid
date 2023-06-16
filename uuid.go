package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
)

// UUID represents a 128-bit Universal Unique Identifier as defined in RFC 4122.
type UUID [16]byte

// Nil UUID
var Nil UUID = UUID{}

var (
	ErrInvalidFormat = errors.New("invalid UUID format")
)

func New() UUID {
	// Return random UUID - described in RFC 4122 as Version 4.

	var uuid UUID

	// Set all bits to randomly (or pseudo-randomly) chosen values.
	_, _ = rand.Read(uuid[:]) // Ignore error

	// Set the bits as described by RFC 4122.
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 10

	return uuid
}

func NewTimeBased() UUID {
	// Return time-based UUID - described in RFC 4122 as Version 4 with first 6 bytes as Unix timestamp.

	var uuid UUID

	// Set first 6 bytes to Unix timestamp.
	tm := time.Now().UnixMilli()
	uuid[0] = byte(tm >> 40)
	uuid[1] = byte(tm >> 32)
	uuid[2] = byte(tm >> 24)
	uuid[3] = byte(tm >> 16)
	uuid[4] = byte(tm >> 8)
	uuid[5] = byte(tm)

	// Set rest of bits to randomly (or pseudo-randomly) chosen values.
	_, _ = rand.Read(uuid[6:]) // Ignore error

	// Set the bits as described by RFC 4122.
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 10

	return uuid
}

// Parse parses a string formatted UUID and returns the UUID.
// Supported formats are "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" and
// "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}".
// If the string is invalid, the UUID returned is Nil and error.
func Parse(str string) (UUID, error) {
	var uuid UUID

	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:
		// Do nothing
	// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
	case 38:
		if str[0] != '{' || str[37] != '}' {
			return Nil, ErrInvalidFormat
		}
		str = str[1:37]
	default:
		return Nil, ErrInvalidFormat
	}

	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return Nil, ErrInvalidFormat
	}

	// Parse groups
	var err error

	// 8 hex digits
	_, err = hex.Decode(uuid[0:4], []byte(str[0:8]))
	if err != nil {
		return Nil, ErrInvalidFormat
	}

	// 4 hex digits
	_, err = hex.Decode(uuid[4:6], []byte(str[9:13]))
	if err != nil {
		return Nil, ErrInvalidFormat
	}

	// 4 hex digits
	_, err = hex.Decode(uuid[6:8], []byte(str[14:18]))
	if err != nil {
		return Nil, ErrInvalidFormat
	}

	// 4 hex digits
	_, err = hex.Decode(uuid[8:10], []byte(str[19:23]))
	if err != nil {
		return Nil, ErrInvalidFormat
	}

	// 12 hex digits
	_, err = hex.Decode(uuid[10:16], []byte(str[24:36]))
	if err != nil {
		return Nil, ErrInvalidFormat
	}

	return uuid, nil
}

func (uuid *UUID) Equal(uuid2 UUID) bool {
	return bytes.Equal(uuid[:], uuid2[:])
}
