/*  The zero value of a slice is nil.

    A nil slice has a length and capacity of 0.  */


package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }
}




/******OUTPUT******
[] 0 0
nil!  */    