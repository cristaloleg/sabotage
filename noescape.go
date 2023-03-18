package sabotage

import "unsafe"

// Noescape hides a pointer from escape analysis.
//
// Noescape is copy/pasted from Go's runtime/stubs.go:noescape(), and is valid as of Go 1.20.
// There is a chance that this approach stops working in future and parameter `p` may escape.
//
// Docs from runtime/stubs.go:
//
// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//
//go:nosplit
func Noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
