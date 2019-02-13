package sabotage

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type rwmutexSpy struct {
	state int32
	sema  uint32
}

// IsRWMutexLocked will return true if sync.Mutex is locked.
func IsRWMutexLocked(m *sync.RWMutex) bool {
	spy := (*rwmutexSpy)(unsafe.Pointer(m))
	return atomic.LoadInt32(&spy.state) == 1
}

// UnlockRWMutex will unlock rwmutex.
func UnlockRWMutex(mutex *sync.RWMutex) {
	spy := (*rwmutexSpy)(unsafe.Pointer(mutex))
	atomic.StoreInt32(&spy.state, 0)
}
