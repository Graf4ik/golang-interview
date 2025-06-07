package livecoding_checklist

import (
	"errors"
	"fmt"
	"os"
)

// 9. Создание ошибки, возврат ее из функции. Вложенные ошибки.

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

// 2. Форматированные ошибки (fmt.Errorf)
func loadConfig(path string) error {
	if path == "" {
		return fmt.Errorf("путь к конфигу не может быть пустым")
	}
	// ...
	return fmt.Errorf("не удалось загрузить конфиг из %s: %v", path, err)
}

func processFile(path string) error {
	content, err := readFile(path)
	if err != nil {
		return fmt.Errorf("ошибка обработки файла: %w", err)
	}
	// ...
}

func readFile(path string) ([]byte, error) {
	return nil, fmt.Errorf("файл не найден: %s", path)
}

type MyError struct {
	Code    int
	Message string
	Err     error
}

func (e *MyError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s (код: %d): %v", e.Message, e.Code, e.Err)
	}
	return fmt.Sprintf("%s (код: %d)", e.Message, e.Code)
}

func (e *MyError) Unwrap() error {
	return e.Err
}

func someOperation() error {
	return &MyError{
		Code:    404,
		Message: "ресурс не найден",
		Err:     os.ErrNotExist,
	}
}

func main() {
	// 1. Создание и возврат ошибок
	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err) // Ошибка: деление на ноль
		return
	}
	fmt.Println("Результат:", result)

	// 3. Вложенные ошибки (Go 1.13+)
	err3 := processFile("config.txt")
	if err3 != nil {
		fmt.Println(err3) // ошибка обработки файла: файл не найден: config.txt

		// Распаковка вложенной ошибки
		if nestedErr := errors.Unwrap(err3); nestedErr != nil {
			fmt.Println("Вложенная ошибка:", nestedErr)
		}

		// Проверка типа вложенной ошибки
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Println("Файл не существует")
		}
	}

	// 4. Кастомные типы ошибок
	err4 := someOperation()
	if err4 != nil {
		var myErr *MyError
		if errors.As(err4, &myErr) {
			fmt.Printf("Кастомная ошибка: код %d, сообщение: %s\n",
				myErr.Code, myErr.Message)
		}
	}

}

// 5. Проверка ошибок (errors.Is и errors.As)
func handleError(err error) {
	// Проверка на конкретную ошибку
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Файл не существует")
		return
	}

	// Проверка на тип ошибки
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Ошибка пути: %v, операция: %s, путь: %s\n",
			pathError.Err, pathError.Op, pathError.Path)
		return
	}

	fmt.Println("Неизвестная ошибка:", err)
}

/*
6. Лучшие практики работы с ошибками
Всегда проверяйте ошибки:

if err != nil {
    // обработайте ошибку
}
Добавляйте контекст к ошибкам:

if err != nil {
    return fmt.Errorf("не удалось обработать запрос: %w", err)
}
Используйте errors.Is для проверки конкретных ошибок:

if errors.Is(err, sql.ErrNoRows) {
    // обработка отсутствия строк
}
Используйте errors.As для проверки типов ошибок:

var netErr net.Error
if errors.As(err, &netErr) && netErr.Timeout() {
    // обработка таймаута
}
Создавайте кастомные ошибки для важных сценариев:

var ErrInvalidInput = errors.New("неверные входные данные")
Документируйте возвращаемые ошибки:

// ParseConfig парсит конфиг и возвращает:
// - ErrInvalidFormat если формат неверный
// - ErrMissingField если обязательные поля отсутствуют
func ParseConfig(data []byte) (*Config, error) {
    // ...
}
*/
