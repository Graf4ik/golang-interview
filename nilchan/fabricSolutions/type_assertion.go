package fabricSolutions

import "fmt"

// добавить код, который выведет в терминал
// тип переменной whoami

func printType(whoami interface{}) {
	if whoami == nil {
		return
	}

	switch whoami.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	}
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
