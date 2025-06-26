package ostrovok

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

/*
У тебя в коде есть серьёзная гонка данных (data race), связанная с переменной rates, которая:
читается из http.HandlerFunc (одновременно, возможно, из многих горутин),
обновляется в background-горутине каждую минуту.
❗ Проблема
Переменная rates — это map[string]float64, и map в Go не потокобезопасна.
Ты:
Читаешь rates[from] в хендлере
Пишешь rates = newrates в другой горутине
*/

var rates atomic.Value // потокобезопасное хранилище

// нужно пофиксить код
func main() {
	initialRates, err := readConversionRates()

	if err != nil {
		log.Fatalf("readConversionRates() failed: %s", err)
	}

	rates.Store(initialRates) // сохранить потокобезопасно

	// background task update conversion rates  фоновая горутина обновления
	go func() {
		for {
			time.Sleep(time.Minute)
			newRates, err := readConversionRates()
			if err != nil {
				log.Printf("readConversionRates() failed: %s", err)
			}
			rates.Store(newRates)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// parse 'from' and 'value' from query params
		from := "RUB"
		val := 140.0

		currentRates, ok := rates.Load().(map[string]float64)
		rate, ok := currentRates[from]
		if !ok {
			http.NotFound(w, r)
			return
		}

		convVal := val / rate

		fmt.Fprint(w, convVal)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
🧠 Почему atomic.Value?
Он потокобезопасен для чтения и записи
Позволяет атомарно заменить всю структуру
Без необходимости в sync.RWMutex
*/

// readConversionRates reads rates from a file or an external service (relatively long running function).
func readConversionRates() (map[string]float64, error) {
	// resp, err := http.Get("https://example.org/conv-rates")
	time.Sleep(100 * time.Millisecond)

	return map[string]float64{
		"USD": 1.0,
		"RUB": 70.0,
		"EUR": 95.0,
	}, nil
}
