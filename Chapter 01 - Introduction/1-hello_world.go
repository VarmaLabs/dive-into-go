package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}


/* Breakdown of how the program is read */

/* Package Declaration */
/*
1. 'package main' - This is known as a “package declaration. ”
2. Packages are Go's way of organizing and reusing code.
3. There are two types of Go programs: executables and libraries.
4. Executable applications are the kinds of programs that we can run directly from the terminal(in Windows they end with .exe)
5. Libraries are collections of code that we package together so that we can use them in other programs. */

/* Import */
/*
1. 'import fmt' - The import keyword is how we include code from other packages to use with our program.
2. The fmt package implements formatting for input and output.
3. The use of double quotes like this is known as a “string literal” which is a type of “expression”.
4. In Go strings represent a sequence of characters (letters, numbers, symbols, …) of a definite length. */

/* Function Declaration */
/*
func main() {
    fmt.Println("Hello World")
}
1. Functions are the building blocks of a Go program.
2. They have inputs, outputs and a series of steps called statements which are executed in order.
3. All functions start with the keyword func followed by the name of the function (main in this case), a list of zero or more “parameters” surrounded by parentheses, an optional return type and a “body” which is surrounded by curly braces.
4. The name main is special because it's the function that gets called when you execute the program. */

/* Final piece */
/*
1. 'fmt.Println("Hello World")' - This statement is made of three components.
2. First we access another function inside of the fmt package called Println.
3. Create a new executable program, which references the fmt library and contains one function called main.
4. The Println function does the real work in this program.
To understand more about Println, type 'godoc fmt Println' in terminal. */

