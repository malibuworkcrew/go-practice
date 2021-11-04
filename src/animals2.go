package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal2 interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	sound string
}

type Bird struct {
	food string
	locomotion string
	sound string
}

type Snake struct {
	food string
	locomotion string
	sound string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}
func (b Bird) Speak() {
	fmt.Println(b.sound)
}
func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func doCommand2(animal Animal2, command string) {
	switch command {
	case "eat": animal.Eat()
	case "move": animal.Move()
	case "speak": animal.Speak()
	}
}


func main() {
	fmt.Println("Enter command in either format 'newanimal [name] [cow,snake,bird]' " +
		"OR 'query [name] [eat,move,speak]'")

	anaMap := map[string]Animal2{
		"cow": Cow{"grass", "walk", "moo"},
		"bird": Bird{"worms", "fly", "peep"},
		"snake": Snake{"mice", "slither", "hsss"},
	}

	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		split := strings.Split(line, " ")
		command := split[0]

		switch command {
		case "newanimal":
			anaMap[split[1]] = anaMap[split[2]]
			fmt.Println("Created it!")
		case "query":
			doCommand2(anaMap[split[1]], split[2])
		}
	}
}