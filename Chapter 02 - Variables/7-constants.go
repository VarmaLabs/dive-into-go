package main

import "fmt"

const Quote = "जननी जन्मभूमिश्च स्वर्गादपि गरीयसी"

func main() {
	const Gopher = "गोफर"
	fmt.Println("Hello", Gopher)
	fmt.Println("Quote from Ramayan:", Quote)
}

/*

Output
------

Hello गोफर
Quote from Ramayan: जननी जन्मभूमिश्च स्वर्गादपि गरीयसी

Notes
-----

You can declare a constant using const keyword
However You cannot use := syntax to declare a constant

*/
