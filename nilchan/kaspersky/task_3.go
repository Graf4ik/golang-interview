package kaspersky

import (
	"fmt"
	"reflect"
)

// как проверить интерфейсы на то, что там лежит под капотом?
// Как достать value?

func main() {
	var s *string         // nil указатель на string
	fmt.Println(s == nil) // true
	var i interface{}     // nil интерфейс
	fmt.Println(i == nil) // true
	i = s                 // интерфейс теперь содержит nil указатель типа *string
	fmt.Println(i == nil) // false - потому что интерфейс содержит тип *string

	// ✅ Вариант 1: Type assertion
	if val, ok := i.(*string); ok {
		if val == nil {
			fmt.Println("i содержит *string со значением nil")
		} else {
			fmt.Println("i содержит *string со значением:", *val)
		}
	}

	// Type Switch:
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case *string:
		fmt.Println("Указатель на string:", v)
	case string:
		fmt.Println("Строка:", v)
	case int:
		fmt.Println("Число:", v)
	default:
		fmt.Println("Неизвестный тип")
	}

	// ✅ Вариант 2: Пакет reflect
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	fmt.Println("тип в интерфейсе:", t)
	fmt.Println("значение в интерфейсе:", v)

	if v.IsNil() {
		fmt.Println("значение nil")
	} else {
		fmt.Println("значение не nil")
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

/*
⚠️ Почему i == nil становится false, даже если s == nil?
Потому что interface{} теперь содержит:
тип: *string
значение: nil
Интерфейс считается nil, только если оба этих поля — nil.
Поэтому:
var x interface{} = nil       // x == nil -> true
var y *string = nil
var z interface{} = y         // z != nil
*/

/*
Пакет reflect в Go предназначен для рефлексии — то есть для возможности программы во время выполнения
узнавать и изменять информацию о типах и значениях данных.

Основные возможности reflect:
Узнать тип переменной (Type)
Узнать значение переменной (Value)
Изменять значения переменных динамически
Вызывать методы объектов по имени
Работать с произвольными структурами, срезами, картами и интерфейсами без жёсткой привязки к типу

Зачем это нужно?
Когда нужен обобщённый код (например, в библиотеках сериализации, десериализации, ORM, фреймворках)
Для написания инструментов, которые работают с типами динамически
Для проверки типов и значений во время выполнения (например, в тестах, или для реализации DeepEqual)
*/
