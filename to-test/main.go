package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func main() {

	var numInts int
	var userNum int
	var waitGroup sync.WaitGroup
	sliceOfInts := make([]int, 0, 3)

	fmt.Println("Please enter a series of integers to place into the array.")
	fmt.Println("How many numbers would you like to add to the array?")
	fmt.Scan(&numInts)

	for i := 0; i < numInts; i++ {

		fmt.Println("Please Enter a number to fill the array")
		_, err := fmt.Scan(&userNum)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Invalid user input")
		}

		sliceOfInts = append(sliceOfInts, userNum)

	}

	/*Partition the array/slice into 4 broken parts*/
	sliceSize := numInts / 4
	slice1 := sliceOfInts[:sliceSize]
	slice2 := sliceOfInts[sliceSize : 2*(sliceSize)]
	slice3 := sliceOfInts[2*(sliceSize) : 3*(sliceSize)]
	slice4 := sliceOfInts[3*(sliceSize):]

	fmt.Println("Here are your arrays when partitioned", slice1, slice2, slice3, slice4)

	waitGroup.Add(4)
	go sortList(slice1)
	waitGroup.Done()
	go sortList(slice2)
	waitGroup.Done()
	go sortList(slice3)
	waitGroup.Done()
	go sortList(slice4)
	waitGroup.Done()
	waitGroup.Wait()

	/*Merge the Slices into one slice and sort them in main goroutine*/
	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)

	//Print the new slice in the goroutine
	fmt.Println("Your Slice merged, sorted and printed is as follows:", newSlice)

}

func sortList(unsortedSlice []int) []int {
	sort.Ints(unsortedSlice)
	return unsortedSlice

}

func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
	newSlice := []int{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Ints(newSlice)
	return newSlice

}
