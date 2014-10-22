package main

import "fmt"

func PrintHello() {
	fmt.Println("Hello Gopher!")
}

func Add(x int, y int) int {
	return x + y
}

func AddAnotherWay(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(21, 76))
	fmt.Println(AddAnotherWay(21, 76))
}

/*

Output
------
Hello Gopher!
97
97

Notes
-----
The func is the keyword to declare a function.
func should be followed by the name of the function.
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
i.e: Add(x int, y int) can also be written as Add(x, y int)

*/
