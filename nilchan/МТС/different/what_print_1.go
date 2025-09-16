package МТС

import (
	"fmt"
	"unsafe"
)

type User struct {
	Valid  bool  // 1 байт + 7 байт паддинг
	Id     int64 // 8 байт
	Number int32 // 4 байта + 4 байта паддинг (если 64-бит)
}

type CUser struct {
	Id     int64 // 8 байт
	Number int32 // 4 байта
	Valid  bool  // 1 байт + 3 байт паддинг
}

// Что выведет код?
func main() {
	user := User{}
	cuser := CUser{}
	fmt.Println(unsafe.Sizeof(user) == unsafe.Sizeof(cuser))
}
