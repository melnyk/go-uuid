package uuid

import "encoding/hex"

func (u UUID) String() string {
	// format as xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	// 8-4-4-4-12 byte groups (32 hex digits + 4 dashes)
	var buf [36]byte

	hex.Encode(buf[:8], u[:4]) // 8 hex digits
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6]) // 4 hex digits
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8]) // 4 hex digits
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10]) // 4 hex digits
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:]) // 12 hex digits

	return string(buf[:])

}
