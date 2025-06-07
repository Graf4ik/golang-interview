package main

import "fmt"

//===========================================================
//Задача 4
//Что выведет код и почему?
//===========================================================

/*
Что происходит:
Первая проверка (var err *MyError):
	err — это nil указатель на MyError (тип *MyError)
	При передаче в errorHandler происходит преобразование в интерфейс error:
		Тип: *MyError
		 Значение: nil
	Проверка err != nil в errorHandler вернет true, потому что интерфейс содержит тип (*MyError), даже если значение nil.
Вторая проверка (err = &MyError{}):
err теперь указывает на реальный экземпляр MyError
При передаче в errorHandler интерфейс error содержит:
Тип: *MyError
Значение: адрес структуры
Проверка err != nil вернет true.

Почему первая проверка (nil *MyError) выводит "Error: MyError!"?
	В Go интерфейс error считается nil только если и тип, и значение nil.
	В первом случае:
		Тип: *MyError (не nil)
		Значение: nil (указатель nil)
	Интерфейс error не nil, поэтому err != nil возвращает true.
	При вызове err.Error() вызывается метод для типа *MyError, который возвращает "MyError!".
*/

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	var err *MyError  // nil указатель на MyError
	errorHandler(err) // Передаем nil *MyError как error
	err = &MyError{}  // Создаем реальный экземпляр
	errorHandler(err) // Передаем &MyError{} как error
}
