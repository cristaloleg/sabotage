package sabotage

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type mutexSpy struct {
	state int32
	sema  uint32
}

// IsMutexLocked will return true if sync.Mutex is locked.
func IsMutexLocked(m *sync.Mutex) bool {
	spy := (*mutexSpy)(unsafe.Pointer(m))
	return atomic.LoadInt32(&spy.state) == 1
}

// UnlockMutex ...
func UnlockMutex(mutex *sync.Mutex) {
	spy := (*mutexSpy)(unsafe.Pointer(mutex))
	atomic.StoreInt32(&spy.state, 0)
}
