package main

import (
	"fmt"
	"sync"
)

func MultipleOf3(channel chan<- int, wait *sync.WaitGroup) {
	defer wait.Done()

	for i := 3; i <= 30; i += 3 {
		channel <- i
	}
	close(channel)
}

func PrintChannel(channel <-chan int, wait *sync.WaitGroup) {
	defer wait.Done()

	for number := range channel {
		fmt.Println(number)
	}
}

func main() {
	var channel = make(chan int)
	var wait sync.WaitGroup

	wait.Add(2)
	go MultipleOf3(channel, &wait)
	go PrintChannel(channel, &wait)

	wait.Wait()
	fmt.Println("\nDONE")
}
