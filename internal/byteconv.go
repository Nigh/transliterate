package internal

import "unsafe"

func BytesToString(b []byte) string {
	// #nosec G103
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	// #nosec G103
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
