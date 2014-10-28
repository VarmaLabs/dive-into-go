// Map literals are like struct literals, but the keys are required. 


package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

func main() {
    fmt.Println(m)
}



/******OUTPUT******
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}] */


//If the top-level type is just a type name, you can omit it from the elements of the literal. 

package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)
}


/******OUTPUT******
map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}] */