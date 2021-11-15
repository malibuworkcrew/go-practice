package main

import (
	"fmt"
	"strconv"
	"time"
)

func writeNums(num int) {
	fmt.Println("Next number = " + strconv.Itoa(num))
}

func main() {
	// These calls should print 1 through 10 in different ways each time we run this process
	for i := 1; i <= 10; i++ {
		go writeNums(i)
	}
	time.Sleep(3 * time.Second)
}
