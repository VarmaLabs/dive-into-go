package main

import "fmt"

// The variable will take the type of the initial value.
var c, python, java = true, false

func main() {
	fmt.Println(c, python, java)
}

/*

Output
------
# command-line-arguments
Chapter 1 - Variables/6-variable-initialisation.go:6: missing expression in var declaration
Chapter 1 - Variables/6-variable-initialisation.go:9: undefined: java

Notes
-----
Either the type of the variables should be there at the time of declaring or all the variables should be initialised so that the types are assigned by the compiler automatically.

*/
