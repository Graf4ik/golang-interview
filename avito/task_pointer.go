package avito

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
	// person.Name = "Alice" // Меняем поле, на которое указывает указатель
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

// Bob Bob
// Как модифицировать программу, чтобы вывелось: Bob, Alice
