package main

import (
	"fmt"
	"sync"
	"time"
)

func MultipleNumber(number int, wait *sync.WaitGroup) {
	defer wait.Done()

	for i := number; i <= number*10; i += number {
		time.Sleep(3 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	var wait sync.WaitGroup

	wait.Add(1)
	go MultipleNumber(3, &wait)

	wait.Wait()
	fmt.Println("DONE")
}
