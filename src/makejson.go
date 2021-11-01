package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string    `json:"name"`
	Address string `json:"address"`
}

func main() {


	var name string
	var address string
	fmt.Println("Terminate program to exit loop (Ctrl+C)")

	for {
		fmt.Println("Enter a name...")
		num1, _ := fmt.Scan(&name)
		_ = num1
		fmt.Println("Enter an address...")
		num2, _ := fmt.Scan(&address)
		_ = num2

		person := Person{Name: name, Address: address}
		barr, _ := json.Marshal(person)
		fmt.Println(string(barr))
	}
}