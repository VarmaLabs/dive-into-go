package main

import "fmt"

func main() {

	// Array DomesticAnimals is declared first as a string array with 2 elements
	// And then first and second elements of array is initialised
	var DomesticAnimals [2]string
	fmt.Println("First Element of DomesticAnimals:", DomesticAnimals[0])
	fmt.Println("Second Element of DomesticAnimals:", DomesticAnimals[1])
	fmt.Println("DomesticAnimals Array:", DomesticAnimals)

	// Arrays gets initialised to default values (Zero Value)
	var x [5]int
	fmt.Println("x: ", x)
}

/*

Output
------
First Element of DomesticAnimals:
Second Element of DomesticAnimals:
DomesticAnimals Array: [ ]
x:  [0 0 0 0 0]

Notes
-----

DomesticAnimals is an array composed of 2 strings
You can see that the values of DomesticAnimals soon after the declaration is empty string "".
This is because Arrays will initialise with the Zero Value of the type of Array

*/
