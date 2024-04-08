package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go boring("boring", ch)

	for i := 0; i <= 5; i++ {
		fmt.Printf("You say: %q\n", <-ch)
	}

	fmt.Println("You're boring. I'm quit!")
}

func boring(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
