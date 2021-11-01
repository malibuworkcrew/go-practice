package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string
	Address string 
}

func user_input(psn *Person) (err error) {
	var psnT Person

	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter user name:")
	psnT.Name, err = stdin.ReadString('\n')
	if err == nil {
		psnT.Name = strings.TrimSpace(psnT.Name)
		fmt.Printf("Enter user address:")
		psnT.Address, err = stdin.ReadString('\n')
		if err == nil {
			psnT.Address = strings.TrimSpace(psnT.Address)
			*psn = psnT
		}
	}
	return err
}

func main() {
	var persons []Person
	var psn Person
	var personsUn []Person

	err := user_input(&psn)
	if err == nil {
		persons = append(persons, psn)
		js, err := json.Marshal(persons)
		if err != nil {
			fmt.Println("Error marshalling: %s)", err)
		} else {
			fmt.Println("Source map: ", persons)
			fmt.Println("JSON data: ", js)
			fmt.Println("Marshalling is Valid? - ", json.Valid(js))
			json.Unmarshal(js, &personsUn)
			fmt.Println("Unmarshall: ", personsUn)
		}
	} else {
		fmt.Println("Input error: %s", err)
	}

}
