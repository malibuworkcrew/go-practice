package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (ani Animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani Animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani Animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]Animal)
	animalMap["cow"] = Animal{"grass", "walk", "moo"}
	animalMap["bird"] = Animal{"worms", "fly", "peep"}
	animalMap["snake"] = Animal{"mice", "slither", "hsss"}

	var mainAnimal animalInterface
	for {

		var command, param1, param2 string
		fmt.Print(">")
		fmt.Scan(&command, &param1, &param2)

		if command == "query" {

			mainAnimal = animalMap[param1]

			switch param2 {
			case "eat":
				mainAnimal.Eat()
			case "move":
				mainAnimal.Move()
			case "speak":
				mainAnimal.Speak()
			}
		}

		if command == "newanimal" {
			animalMap[param1] = animalMap[param2]
			fmt.Println("Created it!")
		}
	}
}
