package main

import (
	"fmt"
	"sync"
)

type Philosopher int
type Chopstick struct {
	mutex sync.Mutex
}

var philOneChannel chan int = make(chan int, 2)
var philTwoChannel chan int = make(chan int, 2)
var philThreeChannel chan int = make(chan int, 2)
var philFourChannel chan int = make(chan int, 2)
var philFiveChannel chan int = make(chan int, 2)
var quit chan int = make(chan int)

func Host(wg *sync.WaitGroup) {
	count := 0
	for {
		select {
		case philOne := <-philOneChannel:
			if philOne == 1 && count < 2 {
				philOneChannel <- 1
				count += 1
			} else if philOne == 0 {
				count -= 1
			}
		case philTwo := <-philTwoChannel:
			if philTwo == 1 && count < 2 {
				philTwoChannel <- 1
				count += 1
			} else if philTwo == 0 {
				count -= 1
			}
		case philThree := <-philThreeChannel:
			if philThree == 1 && count < 2 {
				philThreeChannel <- 1
				count += 1
			} else if philThree == 0 {
				count -= 1
			}
		case philFour := <-philFourChannel:
			if philFour == 1 && count < 2 {
				philFourChannel <- 1
				count += 1
			} else if philFour == 0 {
				count -= 1
			}
		case philFive := <-philFiveChannel:
			if philFive == 1 && count < 2 {
				philFiveChannel <- 1
				count += 1
			} else if philFive == 0 {
				count -= 1
			}
		case <-quit:
			wg.Done()
			return
		}
	}
}
func Eat(philosopherId Philosopher, chopsticks []Chopstick, requests []chan int, wg *sync.WaitGroup) {
	for {
		requests[philosopherId] <- 1
		response := <-requests[philosopherId]
		fmt.Printf("Permission Granted from Host for %d to eat\n", philosopherId+1)
		if response == 1 {
			break
		}
	}
	chopsticks[philosopherId].mutex.Lock()
	chopsticks[(philosopherId+1)%5].mutex.Lock()
	fmt.Println("Starting to eat", philosopherId+1)
	for i := 3; i > 0; i-- {
	}
	fmt.Println("Finished Eating", philosopherId+1)
	chopsticks[philosopherId].mutex.Unlock()
	chopsticks[(philosopherId+1)%5].mutex.Unlock()
	requests[philosopherId] <- 0
	wg.Done()
}
func main() {
	philosophers := []Philosopher{1, 2, 3, 4, 5}
	requestChannels := []chan int{philOneChannel, philTwoChannel, philThreeChannel, philFourChannel, philFiveChannel}
	chopsticks := make([]Chopstick, 5)

	var wg sync.WaitGroup
	wg.Add(1)
	go Host(&wg)
	for _, philosopherId := range philosophers {
		wg.Add(1)
		go Eat(philosopherId-1, chopsticks, requestChannels, &wg)
	}
	quit <- 1
	wg.Wait()
}
