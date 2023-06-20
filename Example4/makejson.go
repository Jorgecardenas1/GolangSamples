package main

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	Name    string
	Address string
}

func main() {

	var name string
	var address string

	fmt.Println("input name:")
	fmt.Scan(&name)
	fmt.Println("input address:")
	fmt.Scanln(&address)

	client := Client{Name: name, Address: address}

	barr, err := json.Marshal(client)

	if err == nil {
		fmt.Println(string(barr))
	} else {
		fmt.Println("Error")
	}
}
