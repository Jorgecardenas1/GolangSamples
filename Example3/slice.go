package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var x string
	var counter = 0

	my_slice := make([]int, 3)

	for x != "X" {

		fmt.Println("input integer:")
		fmt.Scan(&x)

		if x != "X" {
			new_int, err := strconv.Atoi(x)

			if err == nil {
				if counter < 3 {
					my_slice[counter] = new_int
				} else {

					my_slice = append(my_slice, new_int)
					sort.Slice(my_slice, func(i, j int) bool {
						return my_slice[j] > my_slice[i]
					})
				}

				fmt.Print(my_slice)
				fmt.Printf("\n")
			} else {
				fmt.Print(err)
			}

			counter = counter + 1
		}

	}

}
