package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.CountChar("Hello World", 'l'))
	fmt.Println(test.CountChar("5  balloons", 5))
	fmt.Println(test.CountChar("   ", ' '))
	fmt.Println(test.CountChar("The 7 deadly sins", '7'))
}