package uuid

import (
	"database/sql/driver"
	"errors"
)

var (
	ErrUnableToScanType = errors.New("unable to scan type into UUID")
)

// Scan implements sql.Scanner for UUID to read values from databases transparently.
func (uuid *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		return uuid.UnmarshalText([]byte(src))

	case []byte:
		// if slice is 16 bytes (equal to UUID size), we assume it's a UUID in a raw byte format
		if len(src) == len(uuid) {
			return uuid.UnmarshalBinary(src)
		}
		return uuid.UnmarshalText([]byte(src))
	}

	return ErrUnableToScanType
}

// Value implements sql.Valuer for UUID to write values from databases transparently.
// The UUID is encoded as a string.
func (uuid UUID) Value() (driver.Value, error) {
	return uuid.String(), nil
}
