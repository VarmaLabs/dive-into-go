/*
1. Idea behind go interfaces is duck typing.(http://en.wikipedia.org/wiki/Duck_typing)
2. Example "When I see a bird that walks like a duck and swims like a duck and quacks like a duck, I call that bird a duck."
3. Meaning that if your object implements all duck's features then there should be no problem using it as a duck. */

package main

import (
    "fmt"
)

type Walker interface {
    Walk() string
}

type Human string
type Dog string

func (human Human) Walk() string { //A human is a walker
    return "I'm a man and I walked!"
}

func (dog Dog) Walk() string { //A dog is a walker
    return "I'm a dog and I walked!"
}

//Make a walker walk
func MakeWalk(w Walker) {
    fmt.Println(w.Walk())
}

func main() {
    var human Human
    var dog Dog
    MakeWalk(human)
    MakeWalk(dog)
}