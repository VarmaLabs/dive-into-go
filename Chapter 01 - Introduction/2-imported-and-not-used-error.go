package main

import "fmt"
import "math"

// this is a comment

func main() {
	fmt.Println("Hello World")
}

/*

Output
------
# command-line-arguments
Chapter 1 - Introduction/2-imported-and-not-used-error.go:4: imported and not used: "math"

Notes
-----

If you don't use a package which is imported, it will throw a compiler error "imported and not used: "<PackageName>""

*/
