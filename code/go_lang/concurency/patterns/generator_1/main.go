package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := makeGenerator(done, &wg)

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("value: ", v)
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}

func makeGenerator(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int, 1)
	var i int

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				fmt.Println("Done")
				return
			default:
				time.Sleep(time.Millisecond + 200)
				ch <- i
				i++

			}
		}
	}()

	return ch
}
