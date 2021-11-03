package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter acceleration, initial velocity, and initial displacement (comma-seperated)...")
	var numStr string

	num, _ := fmt.Scan(&numStr)
	_ = num

	split := strings.Split(numStr, ",")
	acc, _ := strconv.ParseFloat(split[0], 64)
	vo, _ := strconv.ParseFloat(split[1], 64)
	so, _ := strconv.ParseFloat(split[2], 64)

	dispFunc := GenDisplaceFn(acc, vo, so)

	for {
		fmt.Println("Enter time to get displacement")
		num, _ := fmt.Scan(&numStr)
		_ = num

		time, _ := strconv.ParseFloat(numStr, 64)
		fmt.Printf("Time: %f\n", dispFunc(time))
	}
}

func GenDisplaceFn(a, vo, so float64) func (float64) float64 {
	return func (t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
	}
}