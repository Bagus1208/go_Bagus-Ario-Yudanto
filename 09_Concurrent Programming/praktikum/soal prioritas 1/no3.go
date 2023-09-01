package main

import (
	"fmt"
	"sync"
	"time"
)

func MultipleOf3(channel chan int, wait *sync.WaitGroup) {
	defer close(channel)
	defer wait.Done()

	for i := 3; i <= 30; i += 3 {
		channel <- i
	}
}

func PrintChannel(channel <-chan int, wait *sync.WaitGroup) {
	defer wait.Done()

	for number := range channel {
		fmt.Println(number)
	}
}

func main() {
	var channel = make(chan int, 10)
	var wait sync.WaitGroup

	wait.Add(2)
	go MultipleOf3(channel, &wait)

	fmt.Println("Data telah ditampung di channel")
	fmt.Println("Data akan dicetak...")
	time.Sleep(2 * time.Second)

	go PrintChannel(channel, &wait)

	wait.Wait()
	fmt.Println("\nDONE")
}
