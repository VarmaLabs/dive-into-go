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
	user := User{Name: "Gita", City: "New Delhi"}
	fmt.Println(user)
}

/*

Output
------
{Gita 0 New Delhi  }

Notes
-----

You can pass a subset of fields by using the 'FieldName: Value' syntax.
The order of the named fields is irrelevant if you use this syntax.
The rest of the fields in the struct will be initialised with the zero values

*/
