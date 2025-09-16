package МТС

import "fmt"

// Исправить функцию, чтобы она работала. Сигнатуру менять нельзя
func printNumber(ptrToNumber interface{}) {
	//if ptrToNumber != nil {
	//	fmt.Println(*ptrToNumber.(*int))
	//} else {
	//	fmt.Println("nil")
	//}

	// Варик 1
	//if ptrToNumber == nil {
	//	fmt.Println("nil")
	//}
	//
	//ptr, ok := ptrToNumber.(*int)
	//if !ok || ptr == nil {
	//	fmt.Println("nil", ptr)
	//	return
	//}
	// fmt.Println(*ptr)

	// Варик 2
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
}

func main() {
	v := 10
	printNumber(&v)
	var pv *int
	printNumber(pv)
	pv = &v
	printNumber(pv)
}
