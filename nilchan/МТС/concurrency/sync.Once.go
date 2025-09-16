package МТС

import (
	"fmt"
	"sync"
)

type Config struct {
	DatabaseURL string
	Port        int
}

var (
	config   Config
	onceMu   sync.Mutex
	onceDone bool
)

func initConfig() {
	config = Config{
		DatabaseURL: "http://localhost:5432",
		Port:        8080,
	}
}

// Функция, которая вызывает инициализацию конфигурации
func LoadConfig() {
	onceMu.Lock()
	defer onceMu.Unlock()

	if !onceDone {
		initConfig()
		onceDone = true
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			LoadConfig()
		}()
	}
	wg.Wait()
	fmt.Printf("Конфигурация: %+v\n", config)
}
