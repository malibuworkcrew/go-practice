package main

import "fmt"

func main() {
	var float float32
	fmt.Println("Terminate program to exit loop (Ctrl+C)")

	for {
		fmt.Printf("Enter a floating point number...\n")
		num, err := fmt.Scan(&float)
		_ = num
		_ = err
		fmt.Printf("Your float as an int: %d\n", int32(float))
	}
}