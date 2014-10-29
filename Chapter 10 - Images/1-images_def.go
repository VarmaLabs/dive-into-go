/*  Package image defines the Image interface:

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}


Also, color.Color and color.Model are interfaces, but we'll ignore that by using the predefined implementations color.
RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package

*/

package main

import (
    "fmt"
    "image"
)

func main() {
    m := image.NewRGBA(image.Rect(0, 0, 100, 100))
    fmt.Println(m.Bounds())
    fmt.Println(m.At(0, 0).RGBA())
}


/******OUTPUT******
(0,0)-(100,100)
0 0 0 0
*/

