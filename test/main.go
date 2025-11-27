package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.CountAlpha("Hello world"))
	fmt.Println(test.CountAlpha("H e l l o"))
	fmt.Println(test.CountAlpha("H1e2l3l4o"))
}