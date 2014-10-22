package main

import "fmt"

// Wrong Declaration
x int

func main() {
	fmt.Println(x)
}

/*

Output
------
# command-line-arguments
Chapter 1 - Variables/2-variable-declaration.go:5: non-declaration statement outside function body
Chapter 1 - Variables/2-variable-declaration.go:5: syntax error: unexpected name, expecting semicolon or newline

Notes
-----
var keyword is mandatory if you are declaring it outside a function
This will result in a compilation error "non-declaration statement outside function body"
In order to make this working, you should make it "var x int"

*/