package parallel

import (
	"fmt"
	"github.com/AleksandraShmakova/backend/pkg/mutex"
)

func Say() {
	m := mutex.NewMutex()

	m.Lock()
	fmt.Println("Locked the mutex")
	go func() {
		defer m.Unlock()
		fmt.Println("Unlocked from goroutine")
	}()

	m.Lock() // Ждем, пока горутина освободит мьютекс
	fmt.Println("Locked again")
	m.Unlock()
	fmt.Println("Unlocked the mutex")
}
