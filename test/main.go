package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(test.RetainFirstHalf("A"))
	fmt.Println(test.RetainFirstHalf(""))
	fmt.Println(test.RetainFirstHalf("Hello World"))
}