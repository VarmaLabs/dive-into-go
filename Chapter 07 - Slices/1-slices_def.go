 /*  A slice points to an array of values and also includes a length.

     []T is a slice with elements of type T.  */


package main

import "fmt"

func main() {

    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }

}

/******OUTPUT******
p == [2 3 5 7 11 13]
p[0] == 2
p[1] == 3
p[2] == 5
p[3] == 7
p[4] == 11
p[5] == 13 */