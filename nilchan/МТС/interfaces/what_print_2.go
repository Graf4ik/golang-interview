package МТС

import "fmt"

// Добавить код, который выведет тип переменной whoami
func printType(whoami interface{}) {
	switch whoami.(type) {
	case string:
		fmt.Println(whoami.(string))
	case int:
		fmt.Println(whoami.(int))
	case bool:
		fmt.Println(whoami.(bool))
	}
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
