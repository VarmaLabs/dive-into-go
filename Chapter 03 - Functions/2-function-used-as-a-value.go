package main

import "fmt"

func PrintHello() {
	fmt.Println("Hello Gopher!")
}

func main() {
	PrintHello()
	x := PrintHello()
	fmt.Println(PrintHello())
}

/*

Output
------
# command-line-arguments
Chapter 2 - Functions/2-function-used-as-a-value.go:11: PrintHello() used as value
Chapter 2 - Functions/2-function-used-as-a-value.go:12: PrintHello() used as value

Notes
-----
PrintHello() means it accept no arguments and return nothing
You cannot assign the return value of a function to a variable nor pass it to another function if the function is declared to return nothing

*/
