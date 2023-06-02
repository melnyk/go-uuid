package uuid

// MarshalText implements encoding.TextMarshaler.
func (uuid UUID) MarshalText() ([]byte, error) {
	buf := []byte(uuid.String())
	return buf, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (uuid *UUID) UnmarshalText(data []byte) error {
	id, err := Parse(string(data))
	if err != nil {
		return err
	}
	*uuid = id
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (uuid UUID) MarshalBinary() ([]byte, error) {
	return uuid[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != len(uuid) {
		return ErrInvalidFormat
	}
	copy(uuid[:], data)
	return nil
}
