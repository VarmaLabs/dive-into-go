package main

import "fmt"

const (
	C     = 1
	A     = "asdasd"
	Big   = 23 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(C)
	fmt.Println(A)
	fmt.Println(needInt(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
