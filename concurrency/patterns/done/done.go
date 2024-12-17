package main

import (
	"fmt"
	"sync"
	"time"
)

func run(done <-chan any, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			select {
			case ch <- i:
				fmt.Println("sent", i)
			case <-done:
			}

			wg.Done()
			fmt.Println("terminated", i)
		}()
	}

	return ch
}

func main() {
	var wg sync.WaitGroup
	done := make(chan any)

	ch := run(done, &wg)

	time.Sleep(time.Second)

	<-ch

	time.Sleep(time.Second)

	<-ch

	close(done)

	wg.Wait() // deadlock occours if don't close done channel
	fmt.Println("done")
}
