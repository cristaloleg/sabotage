package sabotage_test

import (
	"sync"
	"testing"

	"github.com/cristaloleg/sabotage"
)

func TestUnlockRWMutex(t *testing.T) {
	mu := &sync.RWMutex{}
	ch := make(chan struct{})

	go func() {
		mu.Lock()
		ch <- struct{}{}
	}()

	sabotage.UnlockRWMutex(mu)
	<-ch
}

func TestIsRWMutexLocked(t *testing.T) {
	mu := &sync.RWMutex{}

	if sabotage.IsRWMutexLocked(mu) {
		t.Error("should be false")
	}

	mu.Lock()

	if !sabotage.IsRWMutexLocked(mu) {
		t.Error("should be true")
	}
}
