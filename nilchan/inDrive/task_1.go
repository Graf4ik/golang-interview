package inDrive

import "fmt"

// что выведет
func main() {
	var strData string
	for i := 0; i < 10; i++ {
		strData += fmt.Sprintf("%d", i)
	}
	fmt.Println(strData) // 0123456789
}

/*
Объяснение:

Переменная strData инициализируется пустой строкой.
В цикле от 0 до 9 к strData последовательно добавляются строки с числами от 0 до 9.
В итоге получается строка "0123456789".
Она выводится через fmt.Println.
*/
