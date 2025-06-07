package avito

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

/*
- Есть функция, которая через нейронную сеть вычисляет прогноз погоды за ~1 секунду
- Есть highload RPC ручка с нагрузкой 10k RPS
- Необходимо реализовать код этой ручки
*/

// aiWeatherForecast через нейронную сеть вычисляет прогноз погоды за ~1 секунду
func aiWeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

/*
	Преимущества

aiWeatherForecast() вызывается один раз в секунду в фоне
Все 10k запросов получают последнее значение мгновенно (без блокировок, быстро)
Нет гонок благодаря atomic.Int32
*/
var cachedForecast atomic.Int32 //

func main() {
	// Инициализация прогноза перед стартом сервера
	cachedForecast.Store(int32(aiWeatherForecast()))

	// Периодическое обновление прогноза в фоне
	go func() {
		for {
			forecast := aiWeatherForecast()
			cachedForecast.Store(int32(forecast))
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Быстро возвращаем последний кэшированный прогноз
		fmt.Fprintf(w, "temperature %d\n", aiWeatherForecast())
	})

	fmt.Println("Listening on :3333...")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
