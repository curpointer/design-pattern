package single

import (
	"fmt"
	"sync"
	"testing"
)

var wgs = sync.WaitGroup{}

func TestSingleOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		wgs.Add(1)
		cnt := i
		goFunc := func() {
			fmt.Printf("go func: %d ", cnt)
			getOtherInstance()
			wgs.Done()
		}
		go goFunc()
	}
	wgs.Wait()
}
