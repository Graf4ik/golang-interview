package burgerKing

import "fmt"

func main() {
	{
		type User struct {
			ID int
		}
		users := map[int]User{
			1: {ID: 1},
		}
		var user *User
		user = &users[1]
		fmt.Println(user) // invalid operation: cannot take address of users[1] (map index expression of struct type User)
	}

	fmt.Println("_____")
}
