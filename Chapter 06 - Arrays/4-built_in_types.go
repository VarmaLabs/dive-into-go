/* Arrays */
/*
1. An array is a numbered sequence of elements of a single type with a fixed length.  */

/* Slices */
/*
1. A slice is a segment of an array.
2. Arrays slices are indexable and have a length.
3. Unlike arrays this length is allowed to change.
4. The only difference between this and an array is the missing length between the brackets. In this case x has been created with a length of 0.
5. To create a slice you should use the built-in make function.
6. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. */




package main

import "fmt"

func main() {
    // var sl []float64
    sl := make([]float64, 5)
    fmt.Println(sl)
    var arr[5] int
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50
    total := 0
    for i:= 0;i < 5 ; i++ {
      total = total + arr[i]
    }
  fmt.Println(total)
}