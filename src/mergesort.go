package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var numStr string

	for {
		fmt.Println("Enter comma-seperated list of integers..")
		num, _ := fmt.Scan(&numStr)
		arr := make([]int, 0, 1)
		_ = num

		split := strings.Split(numStr, ",")
		for _, s := range split {
			newNum, _ := strconv.Atoi(s)
			arr = append(arr, newNum)
		}

		splitGroups := int(math.Min(4, float64(len(arr))))
		groupSize := len(arr) / splitGroups

		for i := 0; i < splitGroups; i++ {
			wg.Add(1)
			var endDex int
			if i == splitGroups - 1 {
				endDex = len(arr)
			} else {
				endDex = groupSize * i + groupSize
			}
			toSort := arr[(groupSize * i):(endDex)]
			fmt.Printf("To Sort: %v\n", toSort)
			go BubbleSort2(toSort, &wg)
		}

		wg.Wait()
		wg.Add(1)
		BubbleSort2(arr[0:], &wg)
		fmt.Printf("%v\n", arr)
	}
}

func BubbleSort2(sli []int, wg *sync.WaitGroup) {
	sorted := false
	for !sorted {
		sorted = true

		for i, val := range sli[0:len(sli) - 1] {
			if val > sli[i + 1] {
				Swap2(sli, i, i + 1)
				sorted = false
			}
		}
	}
	wg.Done()
}

func Swap2(sli []int, l, r int) {
	temp := sli[l]
	sli[l] = sli[r]
	sli[r] = temp
}