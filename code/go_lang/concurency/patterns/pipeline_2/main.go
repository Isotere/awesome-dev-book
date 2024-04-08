package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// Бесконечный генератор который возвращает канал в который по кругу бесконечно пишет входные значения
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)

			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	// Генератор, которые бесконечно повторяет вызов функции
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)

			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	// Функция, которая берет из канала первые N значений
	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})

		go func() {
			defer close(takeStream)

			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()

		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
	fmt.Println()

	randFn := func() interface{} {
		return rand.Int()
	}

	for num := range take(done, repeatFn(done, randFn), 10) {
		fmt.Println(num)
	}
}
