package main

import "fmt"

func swap(x, y string) {
	return y, x
}

func main() {
	a, b := swap("Ram", "Chandran")
	fmt.Println(a, b)
}

/*

Output
------
# command-line-arguments
Chapter 2 - Functions/4-error.go:6: too many arguments to return
Chapter 2 - Functions/4-error.go:10: swap("Ram", "Chandran") used as value

Notes
-----

You need to mention the type of the return values explicity after declaring the type of its input values
Otherwise, you will get an error "too many arguments to return"

*/
