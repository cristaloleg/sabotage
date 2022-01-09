package sabotage

import (
	_ "unsafe"
)

// Pin the current P, returns pid.
func Pin() int {
	return runtime_procPin()
}

// Unpin the current P.
func Unpin() {
	runtime_procUnpin()
}

// Pid returns the id of a current {}.
func Pid() int {
	id := runtime_procPin()
	runtime_procUnpin()
	return id
}

//go:noescape
//go:linkname runtime_procPin runtime.procPin
func runtime_procPin() int

//go:noescape
//go:linkname runtime_procUnpin runtime.procUnpin
func runtime_procUnpin()
