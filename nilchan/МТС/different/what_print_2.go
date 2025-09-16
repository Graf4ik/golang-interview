package МТС

import (
	"fmt"
)

func main() {
	str := "Привет"
	str[2] = 'e'
	fmt.Println(str) // cannot assign to str[2] (neither addressable nor a map index expression)
}
