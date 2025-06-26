package smartWay

import "fmt"

type MyStruct struct{}

func (s MyStruct) Hello() {
	fmt.Println("Hello World")
}

// что выведет в консоль
func main() {
	a := MyStruct{}
	b := &a
	b.Hello()
}

/*
✅ Почему это работает:
В Go метод Hello() определён для значения (MyStruct), а не для указателя (*MyStruct), но Go автоматически делает допустимое преобразование:
Когда ты вызываешь b.Hello() и b — это *MyStruct,
Компилятор Go разыменовывает указатель b в *b и вызывает метод Hello() на значении.
Это называется auto-dereferencing при вызове метода.

📌 Правило:
Метод со значением (func (s MyStruct)) можно вызвать как на значении, так и на указателе.
Метод с указателем (func (s *MyStruct)) — нельзя вызвать на значении, если Go не может автоматически взять адрес.
*/
