package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var newStr string
	var arr = make([]int, 0, 3)

	for {
		fmt.Println("Enter an int (or 'X' to exit)...")
		num, _ := fmt.Scan(&newStr)
		_ = num

		if newStr == "X" {
			fmt.Println("Saw 'X', exiting program")
			break
		} else {
			newNum, _ := strconv.Atoi(newStr)
			arr = append(arr, newNum)
			sort.Ints(arr)
			fmt.Printf("%v\n", arr)
		}
	}
}