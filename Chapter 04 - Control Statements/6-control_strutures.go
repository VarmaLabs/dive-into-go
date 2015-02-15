/* Assume we have a program to print numbers from 1 to 5 without any control structures */
/*
1. 'for' structure allows us to repeat a list of statements multiple times.
2. 'if' structure conditionally decide which part of the block gets executed. */

package main
import "fmt"

func main() {
  fmt.Println(`1
    2
    3
    4
5`) // Prints 1 to 5
  fmt.Println("\n")
  fmt.Println("For Loop \n")

  i := 1
  for i <= 50 {
    if i % 2 == 0 {
      fmt.Println("Divisible by 2")
    } else if i % 3 == 0 {
      fmt.Println("Divisble by 3")
    } else if i % 4 == 0 {
      fmt.Println("Divisble by 4")
    }
  fmt.Println(i) // Prints 1 to 5 using loop
  i = i + 1
}
}

