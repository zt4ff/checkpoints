package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Print(test.PrintIf("abcdefz"))
	fmt.Print(test.PrintIf("abc"))
	fmt.Print(test.PrintIf(""))
	fmt.Print(test.PrintIf("14"))
}