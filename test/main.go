package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Print(test.PrintIfNot("abcdefz"))
	fmt.Print(test.PrintIfNot("abc"))
	fmt.Print(test.PrintIfNot(""))
	fmt.Print(test.PrintIfNot("14"))
}