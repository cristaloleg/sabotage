package sabotage_test

import (
	"sync"
	"testing"

	"github.com/cristaloleg/sabotage"
)

func TestUnlockMutex(t *testing.T) {
	mu := &sync.Mutex{}
	ch := make(chan struct{})

	go func() {
		mu.Lock()
		ch <- struct{}{}
	}()

	sabotage.UnlockMutex(mu)
	<-ch
}

func TestIsMutexLocked(t *testing.T) {
	mu := &sync.Mutex{}

	if sabotage.IsMutexLocked(mu) {
		t.Error("should be false")
	}

	mu.Lock()

	if !sabotage.IsMutexLocked(mu) {
		t.Error("should be true")
	}
}
