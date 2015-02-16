/* Number Datatype
1. Go has several different types to represent numbers. Generally we split numbers into two different kinds: integers and floating-point numbers. */

/* Integers */
/*
1. Numbers without a decimal component.
2. Go's integer types are: uint8, uint16, uint32, uint64, int8, int16, int32 and int64. 8, 16, 32 and 64 tell us how many bits each of the types use.
3. uint means “unsigned integer” while int means “signed integer”.
4. Unsigned integers only contain positive numbers (or zero).
5. Two alias types: byte which is the same as uint8 and rune which is the same as int32.
6. There are also 3 machine dependent integer types: uint, int and uintptr.
7. They are machine dependent because their size depends on the type of architecture you are using. */

/* Floating Point Numbers */
/*
1. Floating point numbers are numbers that contain a decimal component (real numbers).
2. Floating point numbers are inexact. Occasionally it is not possible to represent a number.
3. Like integers floating point numbers have a certain size (32 bit or 64 bit). Using a larger sized floating point number increases it's precision.
4. n addition to numbers there are several other values which can be represented: “not a number” (NaN, for things like 0/0) and positive and negative infinity. (+∞ and −∞)
5. Go has two floating point types: float32 and float64.
6. two additional types for representing complex numbers (numbers with imaginary parts): complex64 and complex128. */

/* Strings */
/*
1. A space is also considered a character.
2. Strings are “indexed” starting at 0 not 1.
3. Concatenation uses the same symbol as addition.
4. Since both sides of the + are strings the compiler assumes you mean concatenation and not addition. */

/* Boolean */
/*
1. A boolean value is a special 1 bit integer type used to represent true and false.
2. Three logical operators are used with boolean values
3. AND - &&
4. OR - ||
5. NOT - ! */

package main

import "fmt"

func main() {
    fmt.Println("1 + 1 =", 1 + 1)
    fmt.Println("1 + 1 =", 1.0 + 1.0)
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[2])
    fmt.Println("Hello " + "World")
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}