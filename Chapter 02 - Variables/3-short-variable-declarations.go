package main

import "fmt"
import "reflect"

func main() {
	x := 1
	i, j := "Ram", 23.0
	fmt.Println(x)
	fmt.Println(i, j)
	fmt.Println("The type of x is:", reflect.TypeOf(x))
	fmt.Println("The type of i is:", reflect.TypeOf(i))
	fmt.Println("The type of j is:", reflect.TypeOf(j))
}

/*

Output
------
1
Ram 23
The type of x is: int
The type of i is: string
The type of j is: float64

Notes
-----
Short Variable Declarations can appear only inside functions.
It can also be used to declare local variables inside the scope of if, for and switch statements
More here - http://golang.org/ref/spec#Short_variable_declarations

reflect - is a library in go where you can find helful meta programming / reflection methods

*/
