package main

import "fmt"

func main() {

	i := 5

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

}

/*

Output
------
Five

Notes
-----

# FIXME - to be rewritten
Define a series of cases after defining switch with variable(s) and it will excute the right case depending on the value of the variable
It also support a default case when any of the cases doesn't match.

*/
