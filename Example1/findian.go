package main

import (
	"fmt"
	"strings"
)

func main() {
	var test_string string

	fmt.Println("input text:")
	fmt.Scan(&test_string)

	split_string := strings.Split(test_string, "")

	if split_string[0] == "i" && split_string[len(split_string)-1] == "n" {
		res1 := strings.ContainsAny(test_string, "i|a|n")

		switch res1 {
		case true:
			fmt.Println("Found!")
		case false:
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")

	}

}
