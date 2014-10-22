package main

import (
	"fmt"
	"strings"
)

func splitAddress(address string) (street, city, state string) {
	var list = strings.Split(address, ",")
	street, city, state = list[0], list[1], list[2]
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
Mahathma Gandhi Road
Bangalore
Karnataka

Notes
-----

Look at the return statement of the function splitAdress.
it is not "return street, city, state" but just "return"
If the results are named, the return statement without arguments will return the current values of the named results
This is working as we have declared and named the return types - (street, city, state string) while declareing the function
In this case if will regturn the current values of street, city and state

*/
