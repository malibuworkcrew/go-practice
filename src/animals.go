package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	sound string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.sound)
}

func doCommand(animal Animal, command string) {
	switch command {
	case "eat": animal.Eat()
	case "move": animal.Move()
	case "speak": animal.Speak()
	}
}

func main() {
	fmt.Println("Enter command in format '[animal] [command]'...")
	var animal, call string

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		split := strings.Split(line, " ")
		animal = split[0]
		call = split[1]

		switch animal {
		case "cow": doCommand(cow, call)
		case "bird": doCommand(bird, call)
		case "snake": doCommand(snake, call)
		}
	}
}