package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FirstWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(piscine.FirstWord(" lorem,ipsum "))
	fmt.Print(piscine.LastWord(" "))
}
