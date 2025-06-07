package livecoding_checklist

import (
	"errors"
	"fmt"
	"os"
)

// 7. Полный пример
var (
	ErrFileNotFound   = errors.New("файл не найден")
	ErrInvalidContent = errors.New("неверное содержимое файла")
)

type ParseError struct {
	Line    int
	Column  int
	Message string
	Err     error
}

func (e *ParseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("ошибка парсинга (строка %d, колонка %d): %s: %v",
			e.Line, e.Column, e.Message, e.Err)
	}
	return fmt.Sprintf("ошибка парсинга (строка %d, колонка %d): %s",
		e.Line, e.Column, e.Message)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

func parseFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotFound, path)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	if len(content) == 0 {
		return fmt.Errorf("%w: файл пуст", ErrInvalidContent)
	}

	// Имитация ошибки парсинга
	return &ParseError{
		Line:    10,
		Column:  5,
		Message: "неожиданный символ",
		Err:     ErrInvalidContent,
	}
}

func main() {
	err := parseFile("missing.txt")
	if err != nil {
		switch {
		case errors.Is(err, ErrFileNotFound):
			fmt.Println("Создаем новый файл...")
			// обработка отсутствия файла

		case errors.Is(err, ErrInvalidContent):
			fmt.Println("Ошибка содержимого:", err)
			// обработка неверного содержимого

		default:
			var parseErr *ParseError
			if errors.As(err, &parseErr) {
				fmt.Printf("Детали ошибки парсинга: строка %d, колонка %d\n",
					parseErr.Line, parseErr.Column)
			} else {
				fmt.Println("Неизвестная ошибка:", err)
			}
		}
	}
}

/*
Этот пример демонстрирует:
Создание базовых ошибок
Обертывание ошибок с добавлением контекста
Создание кастомного типа ошибки
Проверку ошибок с помощью errors.Is и errors.As
Различные способы обработки ошибок
*/
