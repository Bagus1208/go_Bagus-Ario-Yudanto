package main

import (
	"fmt"
	"strconv"
	"sync"
)

func FreqLetter(letter string, channel chan<- string, wait *sync.WaitGroup) {
	var store = make(map[string]int)

	defer close(channel)
	defer wait.Done()

	for i := 0; i < len(letter); i++ {
		if string(letter[i]) == " " {
			continue
		}

		current := string(letter[i])

		if _, found := store[current]; found {
			store[current]++
		} else {
			store[current] = 1
		}
	}

	for key, total := range store {
		channel <- (key + ": " + strconv.Itoa(total))
	}
}

func ShowData(channel <-chan string, wait *sync.WaitGroup) {
	defer wait.Done()

	for data := range channel {
		fmt.Println(data)
	}
}

func main() {
	var letter = "lorem ipsum nlsdfhi nknduh tnjfun kfnakhf l;wnqp nsf nk sdfjhp"
	var channel = make(chan string)
	var wait sync.WaitGroup

	wait.Add(2)
	go FreqLetter(letter, channel, &wait)
	go ShowData(channel, &wait)

	wait.Wait()
}
