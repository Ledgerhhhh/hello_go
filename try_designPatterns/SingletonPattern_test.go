package try_designPatterns

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

var (
	SingletonInstance *Singleton
	once              sync.Once
)

type Singleton struct {
	data string
}

func getInstance() {
	for i := 0; i < 10; i++ {
		once.Do(func() {
			SingletonInstance = &Singleton{
				data: "Hello from Singleton" + strconv.Itoa(i),
			}
		})
	}
	fmt.Println(SingletonInstance.data)
}

func TestGetOne(t *testing.T) {
	getInstance()
}
