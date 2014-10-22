package main

import (
	"fmt"
)

func Capital(CountryName string) string {
	if CountryName == "India" {
		return "New Delhi"
	} else if CountryName == "Pakisthan" {
		return "Islamabad"
	} else {
		return "Invalid Country"
	}
}

func main() {
	fmt.Println(Capital("India"))
	fmt.Println(Capital("Pakisthan"))
	fmt.Println(Capital("Sri Lanka"))
}

/*

Output
------
New Delhi
Islamabad
Invalid Country

Notes
-----

The if syntax doens't require the paranthesis() and flower brackets {} are mandatory

*/
