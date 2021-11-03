package main

import "fmt"

func main() {
	var accel float64
	fmt.Print("Acceleration:")
	fmt.Scanln(&accel)
	var velo float64
	fmt.Print("Initial velocity:")
	fmt.Scanln(&velo)
	var displace float64
	fmt.Print("Initial displacement:")
	fmt.Scanln(&displace)
	var time float64
	fmt.Print("Time:")
	fmt.Scanln(&time)
	fn := GenDisplaceFn(accel, velo, displace)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(accel, velo, displace float64) func(float64) float64 {
	return func(time float64) float64 {
		return 1/2*accel*time*time + velo*time + displace
	}
}
