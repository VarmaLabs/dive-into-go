package main

import "fmt"

type User struct {
	Name    string
	Age     int
	City    string
	State   string
	Country string
}

func main() {
	user := User{"Sita", 21, "New Delhi", "Delhi"}
	fmt.Println(user)
}

/*

Output
------
# command-line-arguments
Chapter 5 - Structs/4-struct-initialisation.go:14: too few values in struct initializer

Notes
-----

You need to pass all the arguments if you are using the comma separated syntax.
Otherwise, it will throw "too few values in struct initializer" error

*/
