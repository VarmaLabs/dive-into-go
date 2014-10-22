package main

import "fmt"

func main() {

	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(len(x))

	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println("Total:", total)

	// Length of Arrays
	var DomesticAnimals [2]string
	DomesticAnimals[0] = "Cow"
	DomesticAnimals[1] = "Goat"

	fmt.Println("There are", len(DomesticAnimals), "domestic animals")
	for i := 0; i < len(DomesticAnimals); i++ {
		fmt.Println(DomesticAnimals[i])
	}

}

/*

Output
------
[0 0 0 0 100]
5
Total: 100
There are 2 domestic animals
Cow
Goat

Notes
-----

len is a method to get the length of the array


*/
