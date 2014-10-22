package main

import "fmt"

func main() {

	// Example 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Example 2
	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// Example 3
	NationalHeros := []string{"Bhagat Singh", "Shivaram Rajguru", "Sukhdev Thapar", "Chandrashekhar Azad"}
	for i := 0; i < len(NationalHeros); i++ {
		fmt.Println(NationalHeros[i])
	}

}

/*

Output
------
45
16
Bhagat Singh
Shivaram Rajguru
Sukhdev Thapar
Chandrashekhar Azad

Notes
-----

Unlike other programming languages, Go has only one construct for looping - for loop
The only change from the C syntax is that the paranthesis is omitted . It is not even optional.
As in C or Java, you can leave the pre and post statements empty.
You can look at the "Chapter x - Arrays" for learining about array declaration and usages

*/
