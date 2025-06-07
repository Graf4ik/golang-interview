package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

//===========================================================
//Задача 2
//1. Нужно написать функцию генератор паролей, которая принимает целое число n, а на выходе строка длины n из букв a-zA-Z и 0-9
//2. Что тут можно улучшить?
//3. Какие тесты ты бы написал для нее? Есть какая-нибудь возможность угадать, какая строка будет генерироваться, чтобы писать тесты?
//===========================================================

func main() {
	fmt.Println(generatePassword(10))
}

func generatePassword(n int) string {
	if n <= 0 {
		return ""
	}

	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	password := make([]rune, n)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

// С использованием strings.Builder (более эффективен для очень длинных паролей):

func generatePassword2(n int) string {
	if n <= 0 {
		return ""
	}

	var builder strings.Builder
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}

	return builder.String()
}

// С криптографически безопасным генератором (для важных паролей):
func securePassword(n int) (string, error) {
	if n <= 0 {
		return "", nil
	}

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = chars[bytes[i]%byte(len(chars))]
	}

	return string(bytes), nil
}

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"Zero length", 0},
		{"Short password", 5},
		{"Normal password", 8},
		{"Long password", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePassword(tt.n)
			if len(got) != tt.n && tt.n > 0 {
				t.Errorf("Expected length %d, got %d", tt.n, len(got))
			}

			// Проверка допустимых символов
			for _, r := range got {
				if !isValidChar(r) {
					t.Errorf("Invalid character in password: %c", r)
				}
			}
		})
	}
}

func isValidChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
