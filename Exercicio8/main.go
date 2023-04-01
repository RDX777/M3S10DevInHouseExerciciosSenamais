package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int)
	for i := 1; i <= 10; i++ {
		go printRandomNumber(i, c)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-c)
	}
}

func printRandomNumber(number int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	c <- number
}
