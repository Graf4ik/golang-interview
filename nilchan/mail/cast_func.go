package mail

import (
	"fmt"
)

// напишите и протестируйте функцию преобразования числа (integer) в
// в десятичную строку("string") 123 -> "123"
// Встроенные функции проеобразования использовать нельзя.
// StringBuilder использовать нельзя, fmt только в тестах
func main() {
	tests := []int{0, 5, 123, -456, 1000, -987654321}

	for _, n := range tests {
		s := IntToString(n)
		fmt.Printf("IntToString(%d) = \"%s\"\n", n, s)
	}
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	// Сохраняем цифры в обратном порядке
	var digits []byte
	for n > 0 {
		d := byte(n % 10)
		digits = append(digits, '0'+d)
		n /= 10
	}

	// Если число отрицательное — добавляем минус
	if negative {
		digits = append(digits, '-')
	}

	// Разворачиваем слайс
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
