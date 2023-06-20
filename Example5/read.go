package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fileDescriptor, err := os.Open("name.txt")
	//bAr<r := make([]byte, 10)

	if err == nil {

		people := []Person{}

		scanner := bufio.NewScanner(fileDescriptor)

		for scanner.Scan() {
			fullname := scanner.Text()
			words := strings.Fields(fullname)

			person := Person{
				fname: words[0],
				lname: words[1],
			}

			people = append(people, person)
		}

		fileDescriptor.Close()

		for _, person := range people {
			fmt.Print("First Name:" + person.fname + " ")
			fmt.Print("Last Name:" + person.lname + "\n")

		}
	}

}
