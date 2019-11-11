package sabotage

import "unsafe"

// HasOverlap reports whether x and y share memory at any index.
// The memory beyond the slice length is ignored.
//
// Based on golang.org/x/crypto/internal/subtle
//
func HasOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 &&
		uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}
