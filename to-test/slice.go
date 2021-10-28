package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func user_input() (val int, quit bool) {
	var valS string
	var err error

	quit = false
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter integer (X - for exit):")
		valS, err = stdin.ReadString('\n')
		if err == nil {
			valS = strings.ToLower(strings.TrimSpace(valS))
		}

		switch {
		case err != nil:
			fmt.Println("Incorrect input. Error - %s\n", err)
		case valS == "x":
			quit = true
			return
		default:
			//			stdin.ReadString('\n')	// clear input buffer
			val, err = strconv.Atoi(valS)
			if err != nil {
				fmt.Println("Incorrect integer conversion.\n")
			} else {
				return
			}

		}
	}
}

func print_ints(ints []int) {
	var b strings.Builder

	b.WriteString("Sorting ints: ")
	for _, v := range ints {
		fmt.Fprintf(&b, "%d ", v)
	}
	b.WriteString("\n")
	fmt.Println(b.String())
}

func main() {
	var ints []int
	var val int
	var quit bool

	for {
		val, quit = user_input()
		if quit {
			print_ints(ints)
			break
		} else {
			ints = append(ints, val)
			sort.Ints(ints)
			print_ints(ints)
		}

	}
}
