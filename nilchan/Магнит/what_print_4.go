package Магнит

import "fmt"

// Что выведется?
func main() {
	slice := []string{"a", "a"} // 2 2

	func(slice []string) {
		slice = append(slice, "b") // 3 4 aab
		slice[0] = "b"             // bab
		slice[1] = "b"             // bbb
		fmt.Println(slice)         // bbb
	}(slice)

	fmt.Println(slice) // aa
}
