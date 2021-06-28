package single

import (
	"fmt"
	"sync"
	"testing"
)

var wg = sync.WaitGroup{}

func TestSingle(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		cnt := i
		goFunc := func() {
			fmt.Printf("go func: %d ", cnt)
			getInstance()
			wg.Done()
		}
		go goFunc()
	}
	wg.Wait()
}
