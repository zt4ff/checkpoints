package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.RectPerimeter(10, 2))
	fmt.Println(test.RectPerimeter(434343, 898989))
	fmt.Println(test.RectPerimeter(10, -2))
}