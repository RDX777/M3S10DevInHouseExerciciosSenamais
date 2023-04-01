package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
}
