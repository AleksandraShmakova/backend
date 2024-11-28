package parallel

import (
	"fmt"
	"github.com/username/mymutexproject/pkg/mutex"
	"strings"
)

// Say выполняет горутины и возвращает результаты их работы как строку.
func Say() string {
	var result strings.Builder
	m := mutex.NewMutex(3)

	// Запуск горутин
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer m.Unlock()
			result.WriteString(fmt.Sprintf("Горутина %d: Привет, мир!\n", i))
		}(i)
	}
	m.Wait()

	return result.String()
}
