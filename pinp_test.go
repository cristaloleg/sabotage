package sabotage

import (
	"runtime"
	"sync"
	"testing"
)

func TestPin(t *testing.T) {
	n := Pin()
	Unpin()
	t.Log(n)

	gomax := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(gomax)
	for i := 0; i < gomax; i++ {
		go func() {
			defer wg.Done()
			n := Pin()
			Unpin()

			t.Log(n)
		}()
	}
	wg.Wait()
}
