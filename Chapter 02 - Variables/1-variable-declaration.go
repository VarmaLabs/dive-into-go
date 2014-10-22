package main

import "fmt"

// Declaring the variable
var i int

func main() {
	// Initialising the variable
	i = 17
	fmt.Println(i)
}

/*

Output
------
17

Notes
-----
Go is similar to C when it comes to syntax. Unlike C, Go's declarations read left to right.
Go was having the semicolon in its declaration initially and it was dropped later considering the readability and many other factors.

*/
