package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.DigitLen(100, 10))
	fmt.Println(test.DigitLen(100, 2))
	fmt.Println(test.DigitLen(-100, 16))
	fmt.Println(test.DigitLen(100, -1))
}