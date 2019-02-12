package sabotage_test

import (
	"sync"
	"testing"

	"github.com/cristaloleg/sabotage"
)

func TestResetSyncOnce(t *testing.T) {
	once := &sync.Once{}

	counter := 0

	work := func(param int) {
		wg := sync.WaitGroup{}
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()

				once.Do(func() {
					counter += param
				})
			}()
		}
		wg.Wait()
	}

	work(42)
	if want := 42; counter != want {
		t.Fatalf("expected %v, got %v", want, counter)
	}

	sabotage.ResetSyncOnce(once)

	counter = 0
	work(78)
	if want := 78; counter != want {
		t.Fatalf("expected %v, got %v", want, counter)
	}
}

func TestIsOnceDone(t *testing.T) {
	once := &sync.Once{}

	if sabotage.IsOnceDone(once) {
		t.Error("should be false")
	}

	once.Do(func() {})

	if !sabotage.IsOnceDone(once) {
		t.Error("should be true")
	}
}
