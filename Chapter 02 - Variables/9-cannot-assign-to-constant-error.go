package main

import "fmt"

const Quote = "जननी जन्मभूमिश्च स्वर्गादपि गरीयसी"

func main() {
	Quote = "सत्यमेव जयते"
	fmt.Println(Quote)
}

/*

Output
------

# command-line-arguments
./9-change-constant-error.go:8: cannot assign to Quote

Notes
-----

You cannot change the value of a constant

*/
