package main

import "fmt"

// All global variables need to use the long form.
// var x int
var x int = 2

func main() {
	fmt.Println(x)
}

/*

Output
------
2

Notes
-----
You cannot use the short variable declaration outside the funciton for some reason.
I believe this restriction is put in place to recogonise the statement boundaries while syntax analysis
I have read it like it is possible with a special usage of semicolan.
For now, var (or const or type or func) is necessary at the top level

*/
