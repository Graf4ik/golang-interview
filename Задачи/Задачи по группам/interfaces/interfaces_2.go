package main

import (
	"fmt"
	"reflect"
)

// ===========================================================
// Задача 2
// 1. Добавить код, который выведет тип переменной whoami
// ===========================================================

func printType(whoami interface{}) {
	fmt.Printf("%T\n", whoami)

	// Способ 1: Использование fmt.Printf с %T
	fmt.Printf("Способ 1: Тип = %T\n", whoami)

	// Способ 2: Использование reflect.TypeOf
	fmt.Println("Способ 2: Тип =", reflect.TypeOf(whoami))

	// Способ 3: Type switch
	switch v := whoami.(type) {
	case int:
		fmt.Println("Способ 3: Это int")
	case string:
		fmt.Println("Способ 3: Это string")
	case bool:
		fmt.Println("Способ 3: Это bool")
	default:
		fmt.Printf("Способ 3: Неизвестный тип: %T\n", v)
	}

	fmt.Println("-----")
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
