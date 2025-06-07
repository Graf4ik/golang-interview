package main

// напишите функцию, которая бы возвращала ошибку не импортируя для этого никакх пакетов
func main() {
	println(handle().Error())
}

type myError string

func (e myError) Error() string {
	return string(e)
}

func handle() error {
	return myError("Что то пошло не так")
}

/*
myError — это тип string, который реализует интерфейс error, так как у него есть метод Error() string.
handle() возвращает значение типа myError, который ведёт себя как error.
И всё это — без импорта errors или других внешних пакетов.
*/
