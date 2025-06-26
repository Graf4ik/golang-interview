package custom

import (
	"fmt"
	"testing"
)

//===========================================================
//Задача 3
//Написать функцию, которая устанавливает i-ый бит числа в 0
//===========================================================

func main() {
	var num int = 15
	var bitPos uint = 2

	result := setBitToZero(num, bitPos)
	fmt.Printf("Число %d (%b) после установки %d-го бита в 0: %d (%b)\n",
		num, num, bitPos, result, result)
}

func setBitToZero(num int, i uint) int {
	if i >= 64 { // Для 64-битных чисел
		return num
	}
	return num &^ (1 << i) //  создаёт маску, где только i-й бит установлен в 1
	// Например, для i=2: 0100 (в двоичном виде)
	/*
		&^ - это оператор "AND NOT" (битовое И-НЕ):
		Сначала инвертирует биты правого операнда
		Затем выполняет побитовое И с левым операндом
		Эффективно сбрасывает только указанный бит
	*/
}

func TestSetBitToZero(t *testing.T) {
	tests := []struct {
		num    int
		bitPos uint
		want   int
	}{
		{15, 2, 11},    // 1111 -> 1011
		{255, 0, 254},  // 11111111 -> 11111110
		{8, 3, 0},      // 1000 -> 0000
		{0, 5, 0},      // 0000 -> 0000
		{1023, 9, 511}, // 1111111111 -> 0111111111
	}

	for _, tt := range tests {
		got := setBitToZero(tt.num, tt.bitPos)
		if got != tt.want {
			t.Errorf("setBitToZero(%d, %d) = %d (%b), want %d (%b)",
				tt.num, tt.bitPos, got, got, tt.want, tt.want)
		}
	}
}
