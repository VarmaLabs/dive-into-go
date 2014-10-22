package main

import "fmt"

func main() {

	// And then first and second elements of array is initialised
	var DomesticAnimals [2]string
	DomesticAnimals[0] = "Cow"
	DomesticAnimals[1] = "Goat"
	fmt.Println("First Element of DomesticAnimals:", DomesticAnimals[0])
	fmt.Println("Second Element of DomesticAnimals:", DomesticAnimals[1])
	fmt.Println("DomesticAnimals Array:", DomesticAnimals)

	var x [5]int
	x[4] = 100
	fmt.Println("x: ", x)

	WildAnimals := []string{"Lion", "Tiger", "Wolf"}
	fmt.Println(WildAnimals)
}

/*

Output
------
First Element of DomesticAnimals: Cow
Second Element of DomesticAnimals: Goat
DomesticAnimals Array: [Cow Goat]
x:  [0 0 0 0 100]
[Lion Tiger Wolf]

Notes
-----

len is a method to get the length of the array
":= []string{el1, el2, el3, ...}" will assign i

*/
