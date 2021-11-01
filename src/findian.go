package main

import (
	"fmt"
	"os"
)
import "strings"
import "bufio"

func main() {
	fmt.Println("Terminate program to exit loop (Ctrl+C)")

	for {
		fmt.Println("Enter a string...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		lowName := strings.ToLower(name)

		switch {
		case strings.HasPrefix(lowName, "i") &&
			strings.HasSuffix(lowName, "n") &&
			strings.Contains(lowName, "a"):
			fmt.Println("Found!")
		default:
			fmt.Println("Not Found!")
		}
	}
}