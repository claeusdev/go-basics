package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Country string
	City    string
}

func main() {

	//Variables
	var i int = 5
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	name := "Hello"
	fmt.Println(name)

	//Pointers
	var some *string = new(string)
	*some = "Hello"
	fmt.Println(*some)

	//Constants
	const pi = 3.14
	fmt.Println(pi)

	//Array
	var arr [3]int
	fmt.Println(arr)

	//Slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)
	fmt.Println(slice[1:3])
	fmt.Println(slice[:2])

	//Maps
	m := map[string]int{"key": 5}
	m["random"] = 3
	fmt.Println(m)

	//Structs
	var user User
	user.FirstName = "Nana"
	user.LastName = "Manu"

	nana := User{
		ID:        1,
		FirstName: "Nana",
		LastName:  "Manu",
		Address: Address{
			Country: "Ghana",
			City:    "Accra",
		},
	}
	fmt.Println(user, nana)

	//Functions
	Print("Hello")
}

func Print(msg string) {
	fmt.Println(msg)
}

func sum(a int, b int) int {
	return a + b
}
