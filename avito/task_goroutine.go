package avito

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Написать код, который будет выводить коды ответов на HTTP-запросы по двум URL адресам
(например главная страница Google и главная страница Facebook)
*/

// Пример решения
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		resp, err := http.Get("http://www.google.com")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Google status: %5\n" + resp.Status)
	}()

	go func() {
		defer wg.Done()
		resp, err := http.Get("http://www.facebook.com")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Facebook status: %5\n" + resp.Status)
	}()

	wg.Wait()
}
