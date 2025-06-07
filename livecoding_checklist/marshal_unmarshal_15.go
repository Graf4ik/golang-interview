package livecoding_checklist

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// 15. Маршал и анмаршал структуры в json.
/*
В Go для работы с JSON используется пакет encoding/json, который предоставляет простые функции для преобразования структур в JSON и обратно.

Основные функции
json.Marshal() - преобразует структуру в JSON (маршалинг)
json.Unmarshal() - преобразует JSON в структуру (анмаршалинг)
*/

// Маршалинг (структура → JSON) Простая структура
type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person2{Name: "Alice", Age: 25}

	// Маршалинг структуры в JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
	// Вывод: {"name":"Alice","age":25}
}

// Настройки маршалинга
type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author,omitempty"` // Поле не будет включено, если пустое
	Pages    int    `json:"pages,string"`     // Число как строка
	Internal string `json:"-"`                // Поле всегда игнорируется
}

func main7() {
	b := Book{Title: "The Go Programming Language", Pages: 380}

	jsonData, _ := json.Marshal(b)
	fmt.Println(string(jsonData))
	// Вывод: {"title":"The Go Programming Language","pages":"380"}
}

// Анмаршалинг (JSON → структура)
// Базовый пример

func main8() {
	jsonStr := `{"name":"Bob","age":30}`

	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", p)
	// Вывод: {Name:Bob Age:30}
}

// Работа с динамическими данными
func main9() {
	jsonStr := `{"product":"Laptop","price":999.99,"tags":["electronics","tech"]}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
	// Вывод: map[price:999.99 product:Laptop tags:[electronics tech]]

	// Извлечение конкретных значений
	price := data["price"].(float64)
	fmt.Println(price) // 999.99
}

// Продвинутые примеры
// Кастомный маршалинг
type CustomDate struct {
	time.Time
}

func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, cd.Format("2006-01-02"))), nil
}

type Event struct {
	Name string     `json:"name"`
	Date CustomDate `json:"date"`
}

func main10() {
	event := Event{
		Name: "Conference",
		Date: CustomDate{time.Now()},
	}

	jsonData, _ := json.Marshal(event)
	fmt.Println(string(jsonData))
	// Вывод: {"name":"Conference","date":"2023-05-15"}
}

// Обработка неопределенных полей
func main11() {
	type Config struct {
		Port int `json:"port"`
	}

	jsonStr := `{"port": 8080, "host": "localhost"}`

	var cfg Config
	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	decoder.DisallowUnknownFields() // Запрещаем неизвестные поля

	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error:", err)
		// Вывод: Error: json: unknown field "host"
	}
}

// Форматированный вывод JSON
func main12() {
	p := Person{Name: "Alice", Age: 25}

	// Форматированный JSON с отступами
	jsonData, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(jsonData))
	/*
	   Вывод:
	   {
	     "name": "Alice",
	     "age": 25
	   }
	*/
}

/*
Практические советы

Все экспортируемые поля структуры должны начинаться с заглавной буквы
Используйте теги json:"name" для настройки имен полей в JSON
Для обработки динамических JSON-структур используйте map[string]interface{}
Для больших JSON-данных используйте json.Decoder вместо json.Unmarshal
Для красивого вывода используйте json.MarshalIndent
Работа с JSON в Go проста и эффективна благодаря стандартной библиотеке. Эти инструменты покрывают большинство потребностей при работе с JSON-данными.
*/
