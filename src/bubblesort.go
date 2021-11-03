package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter comma-seperated list of 10 integers..")
	var numStr string
	var arr [10]int

	for {
		num, _ := fmt.Scan(&numStr)
		_ = num

		split := strings.Split(numStr, ",")
		for i, s := range split {
			toInt, _ := strconv.Atoi(s)
			arr[i] = toInt
		}

		BubbleSort(arr[0:10])
		fmt.Printf("%v\n", arr)
	}
}

func BubbleSort(sli []int) {
	sorted := false
	for !sorted {
		sorted = true

		for i, val := range sli[0:len(sli) - 1] {
			if val > sli[i + 1] {
				Swap(sli, i, i + 1)
				sorted = false
			}
		}
	}
}

func Swap(sli []int, l, r int) {
	temp := sli[l]
	sli[l] = sli[r]
	sli[r] = temp
}