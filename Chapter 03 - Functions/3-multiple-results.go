package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Ram", "Chandran")
	fmt.Println(a, b)
}

/*

Output
------
Chandran Ram

Notes
-----

A Go function can return multiple values
However you need to mention the type of the return values explicity after declaring the type of its input values

*/
