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
	user := User{"Lakshmi", 22, "Mumbai", "Maharashtra", "India"}
	ptr := &user
	ptr.Name = "Gouri"
	ptr.Age = 23
	ptr.City = "Calicut"
	ptr.State = "Kerala"
	fmt.Println(ptr)
	fmt.Println(user)
}

/*

Output
------
&{Gouri 23 Calicut Kerala India}
{Gouri 23 Calicut Kerala India}

Notes
-----

The if syntax doens't require the paranthesis() and flower brackets {} are mandatory

*/
