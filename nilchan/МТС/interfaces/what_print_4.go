package МТС

import "fmt"

// Что выведет код и почему?
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
	var err *MyError
	errorHandler(err) // nil
	err = &MyError{}
	errorHandler(err) // MyError!
}
