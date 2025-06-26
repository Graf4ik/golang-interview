package CRTEX

import "fmt"

// Что выведет программа?

func main() {
	m := make(map[string]int)
	m["1"] = 0
	fmt.Printf("%v", &m["1"])
}

/*
invalid operation: cannot take address of m["1"] (map index expression of type int)
У элемента нельзя брать адрес, потому что из за возможной эвакуации данных адреса
элементов поменяются, поэтому брать адрес элемента мапы не безопасно
*/
