package single

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleother struct {
}

var singleOtherInstance *singleother

func getOtherInstance() *singleother {
	if singleOtherInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleOtherInstance = &singleother{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleOtherInstance
}
