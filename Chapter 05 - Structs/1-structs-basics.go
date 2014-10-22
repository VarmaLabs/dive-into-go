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
	fmt.Println(User{"Krishnan", 30, "Mysore", "Karnataka", "India"})
}

/*

Output
------
{Krishnan 30 Mysore Karnataka India}

Notes
-----

Struct in Go is just like it in C.
It is a colleciton of fields.
Most helpful if you want to create your own data types.

*/
