package main

import "fmt"

// ===========================================================
// Задача 3
// Исправить функцию, чтобы она работала. Сигнатуру менять нельзя
// ===========================================================

/*
ункция printNumber ожидает указатель на число (*int),
но принимает его как interface{}. При попытке разыменования
через type assertion (*ptrToNumber.(*int)) возникает паника,
если переданный интерфейс содержит nil-указатель (хотя сам интерфейс не nil).
*/

/*
Ключевые изменения:

	Проверка на nil самого интерфейса
	Безопасный type assertion с проверкой (ok)
	Дополнительная проверка на nil после приведения типа
	Ранний возврат в случае ошибок
*/
func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	ptr, ok := ptrToNumber.(*int)
	if !ok || ptr == nil {
		fmt.Println("nil")
		return
	}

	/*
	 Вариант 2
	 switch v := ptrToNumber.(type) {
	    case *int:
	        if v != nil {
	            fmt.Println(*v)
	        } else {
	            fmt.Println("nil")
	        }
	    default:
	        fmt.Println("nil")
	    }
	*/

	fmt.Println(*ptr)
}

func main() {
	v := 10
	printNumber(&v)
	var pv *int
	printNumber(pv)
	pv = &v
	printNumber(pv)
}
