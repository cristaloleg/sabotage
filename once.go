package sabotage

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type onceSpy struct {
	m    sync.Mutex
	done uint32
}

// ResetSyncOnce will reset sync.Once to default state.
func ResetSyncOnce(once *sync.Once) {
	spy := (*onceSpy)(unsafe.Pointer(once))
	atomic.StoreUint32(&spy.done, 0)
}

// IsOnceDone will return true if sync.Once.Do was invoked.
func IsOnceDone(once *sync.Once) bool {
	spy := (*onceSpy)(unsafe.Pointer(once))
	return atomic.LoadUint32(&spy.done) == 1
}
