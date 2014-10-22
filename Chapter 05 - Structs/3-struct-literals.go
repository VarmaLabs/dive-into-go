package main

import "fmt"
import "reflect"

type User struct {
	Name    string
	Age     int
	City    string
	State   string
	Country string
}

func main() {
	user := User{"Sasi", 22, "Salem", "Tamil Nadu", "India"}
	ptr := &user
	fmt.Println("The type of user is:", reflect.TypeOf(user))
	fmt.Println("The type of ptr is:", reflect.TypeOf(ptr))
}

/*

Output
------
The type of user is: main.User
The type of ptr is: *main.User

Notes
-----

ptr is a pointer to Struct user

reflect - is a library in go where you can find helful meta programming / reflection methods

*/
