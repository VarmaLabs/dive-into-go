package main

import (
	"fmt"
	"strings"
)

func splitAddress(address string) (street, city, state string) {
	var list = strings.Split(address, ",")
	var street, city, state = list[0], list[1], list[2]
	return
}

func main() {
	var street, city, state = splitAddress("Mahathma Gandhi Road,Bangalore,Karnataka")
	fmt.Println(street)
	fmt.Println(city)
	fmt.Println(state)
}

/*

Output
------
# command-line-arguments
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:10: street redeclared in this block
  previous declaration at Chapter 2 - Functions/6-named-results-and-redeclared-error.go:8
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:10: city redeclared in this block
  previous declaration at Chapter 2 - Functions/6-named-results-and-redeclared-error.go:8
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:10: state redeclared in this block
  previous declaration at Chapter 2 - Functions/6-named-results-and-redeclared-error.go:8
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:11: street is shadowed during return
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:11: city is shadowed during return
Chapter 2 - Functions/6-named-results-and-redeclared-error.go:11: state is shadowed during return

Notes
-----

When you use named results (street, city, state string), you are already declaring the variables.
Hence you dont need to redeclare it inside the fucntion which is what we were trying to do in this example
var street, city, state = list[0], list[1], list[2]
this will result in a "redeclared error"

*/
