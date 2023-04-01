package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func(n int) {
			defer wg.Done()
			ch <- n
		}(i)
	}

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
	close(ch)
}
