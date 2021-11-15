/*
EXPLAINATION of race condition:

So in the below program, we have to goroutines one is called runSimpleReader,
th other is runSimpleWriter, they are tryin both to access the variable sharedInt concurrently,
which may lead to the race condition and different execution interleaving with each run of the program


*/
package main

import (
	"fmt"
	"time"
)

var sharedInt int = 0
var unusedValue int = 0
var sli = make([]int, 0)

func runSimpleReader() {
	for {
		//var val int = sharedInt

		//fmt.Println(val)
		if sharedInt%10 == 0 {
			unusedValue = unusedValue + 1
		}

		if sharedInt > 0 {
			fmt.Println(sharedInt)
		}
	}
}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
		if sharedInt%1000 == 0 {
			fmt.Println(sharedInt)
		}
	}
}

func startSimpleReadWrite() {
	go runSimpleWriter()
	go runSimpleReader()
	time.Sleep(20 * time.Millisecond)
}

func main() {
	startSimpleReadWrite()
	fmt.Println(sharedInt)
	fmt.Println(unusedValue)

}
