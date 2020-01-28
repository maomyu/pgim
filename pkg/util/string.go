package util

import "unsafe"

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
