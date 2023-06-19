package main

import (
	"fmt"
)

func main() {

	var x float32

	fmt.Println("input float:")
	fmt.Scan(&x)

	y := int(x)

	fmt.Print(y)
}
