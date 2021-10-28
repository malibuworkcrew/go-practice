package main

import "fmt"
import "strings"

func main() {
	var name string
	fmt.Println("Terminate program to exit loop (Ctrl+C)")

	for {
		fmt.Println("Enter a string...")
		num, _ := fmt.Scan(&name)
		lowName := strings.ToLower(name)
		_ = num

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