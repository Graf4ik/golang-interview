package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//==========================================
//Задача 1
//1. Написать функцию, которая принимает число N и возвращает слайс размера N с уникальными числами.
//2. Идеи как тестировать функцию?
//==========================================

func main() {
	res := uniqueDigits(4)
	fmt.Println(res)
}

func uniqueDigits(n int) []int {
	if n <= 0 {
		return []int{}
	}

	res := make([]int, 0, n)
	unique := make(map[int]bool)

	rand.Seed(time.Now().UnixNano())
	for len(res) < n {
		num := rand.Intn(100)
		if !unique[num] {
			unique[num] = true
			res = append(res, num)
		}
	}
	return res
}

func TestUniqueDigits(t *testing.T) {
	// Проверка размера результата
	result := uniqueDigits(5)
	if len(result) != 5 {
		t.Errorf("Expected 5 numbers, got %d", len(result))
	}

	// Проверка уникальности
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			t.Errorf("Duplicate number found: %d", num)
		}
		seen[num] = true
	}

	// Проверка граничных случаев
	if len(uniqueDigits(0)) != 0 {
		t.Error("Expected empty slice for n=0")
	}
}
