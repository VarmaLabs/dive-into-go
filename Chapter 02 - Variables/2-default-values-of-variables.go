package main

import "fmt"

// Declaring the variables
var a int
var b float64
var c string
var d bool
var e *int

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

/*

Output
------
0
0

false
<nil>

Notes
-----
Memory is allocated when the variable is declared or when new method is called.
During this time, each variable is set to a default value which is called zero value for its type
int - 0
float64 - 0.0
string - ""
bool - false
pointer - nil
More here : http://golang.org/ref/spec#The_zero_value

*/
