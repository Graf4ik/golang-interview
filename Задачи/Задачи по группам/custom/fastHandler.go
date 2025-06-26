package custom

import (
	"math/rand"
	"net/http"
	"time"
)

//===========================================================
//Задача 5
//1. Релизовать ручку так, чтобы она выполнялась быстрее чем за одну секунду
//2. Теперь допустим, что запрашивается температура в каком-то location_id. Опиши, как это реализовать.
//===========================================================

// Есть функция getWeather, которая через нейронную сеть вычисляет температуру за ~1 секунду
// Есть highload ручка /weather/highload с нагрузкой 3k RPS
// Необходимо реализовать код этой ручки

func getWeather() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	http.HandleFunc("/weather/highload", func(resp http.ResponseWriter, req *http.Request) {

	})
}
