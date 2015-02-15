/* Variables */
/*
1. A variable is a storage location, with a specific type and an associated name.
2. Variables in Go are created by first using the var keyword, then specifying the variable name(message), the type (string) and finally assigning a value to the variable (Hello World).
3. ':=' is a short hand operator, The type is not necessary because the Go compiler is able to infer the type based on the literal value you assign the variable. */

/* Constants */
/*
1. Go also has support for constants.
2. Constants are basically variables whose values cannot be changed later.
3. They are created in the same way you create variables but instead of using the var keyword we use the const keyword.
4. Constants are a good way to reuse common values in a program without writing them out each time. */

/* Defining Multiple Variables */
/*
1. Use the keyword var (or const) followed by parentheses with each variable on its own line. */

package main

import "fmt"

func main() {

    const message3 string = "Today is Monday"
    fmt.Println(message3) // Today is Monday

    var message string = "Hello World"
    fmt.Println(message) // Hello World

    var message1 string
    var message2 string

    message1 = "Hello"
    fmt.Println(message1) // Hello
    message2 = "World"
    fmt.Println(message2) // World


    message1 = "Hello"
    fmt.Println(message1) // Hello
    message1 = message1 + "World"
    fmt.Println(message1) // Hello World

    message1 = "Hello"
    message2 = "World"
    fmt.Println(message1 == message2) // false

    message1 = "Hello"
    message2 = "Hello"
    fmt.Println(message1 == message2) // true

    fmt.Println("Enter a Value: ")
    var message4 string
    fmt.Scanf("%s", &message4)

    message5 := message4 + "\t" + "Technologies"
    fmt.Println(message5) // Qwinix Technologies
}