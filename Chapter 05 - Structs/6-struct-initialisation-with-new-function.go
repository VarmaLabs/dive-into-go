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
	user1 := new(User)
	fmt.Println(user1)

	var user2 *User = new(User)
	fmt.Println(user2)

	user1.Name, user1.Age = "Sita", 22
	fmt.Println(user1)

	user2.Name, user2.Age = "Gita", 23
	fmt.Println(user2)
}

/*

Output
------
&{ 0   }
&{ 0   }
&{Sita 22   }
&{Gita 23   }

Notes
-----

"t := new(T)" is equivalent to "var t *T = new(T)"
new(T) allocates zero value of the struct T and returns a pointer

*/
