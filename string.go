package sabotage

import (
	"reflect"
	"runtime"
	"unsafe"
)

// BytesToString converts []byte to a string without an allocation.
func BytesToString(bytes []byte) (s string) {
	// Also can be just one line:
	// return *(*string)(unsafe.Pointer(&bs))
	// See: https://github.com/golang/go/issues/25484

	slice := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = slice.Data
	str.Len = slice.Len
	runtime.KeepAlive(&bytes)
	return s
}

// StringToBytes converts string to []byte without an allocation.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
