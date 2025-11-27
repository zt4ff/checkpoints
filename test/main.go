package main

import (
	"fmt"
	"test"
)

func main () {
	fmt.Println(test.CamelToSnakeCase("HelloWorld"))
	fmt.Println(test.CamelToSnakeCase("helloWorld"))
	fmt.Println(test.CamelToSnakeCase("camelCase"))
	fmt.Println(test.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(test.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(test.CamelToSnakeCase("hey2"))
}