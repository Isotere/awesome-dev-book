package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const rateLimit = time.Second / 5

var (
	veryBigSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	retryMax     = 3
	batchSize    = 5
)

type Payload struct {
	Error  error
	Result int
}

func main() {
	result, err := ourReportMaker()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func workerRoutine(doneCh <-chan struct{}, input []int) <-chan Payload {
	ch := make(chan Payload, 1)

	go func() {
		defer close(ch)

		select {
		case <-doneCh:
			fmt.Fprintln(os.Stdout, "Channel closed at start")
			return
		default:
			{
				result, err := retryRemoutFunc(remoutFunc, input)
				if err != nil {
					ch <- Payload{
						Error: err,
					}
					return
				}

				for _, v := range result {
					select {
					case <-doneCh:

						fmt.Fprintln(os.Stdout, "Channel closed at seeding")
						return
					default:
						ch <- Payload{
							Result: v,
						}
					}
				}

				return
			}
		}
	}()

	return ch
}

func ourReportMaker() (map[int]Payload, error) {
	throttle := time.Tick(rateLimit)

	if len(veryBigSlice) == 0 {
		return nil, nil
	}

	rB := 0
	lB := batchSize

	doneCh := make(chan struct{})
	defer close(doneCh)

	workerChannels := make([]<-chan Payload, 0)

	for rB < len(veryBigSlice) {
		var batch []int
		if (rB + lB) < len(veryBigSlice) {
			batch = veryBigSlice[rB:lB]
		} else {
			batch = veryBigSlice[rB:]
		}

		workerChannels = append(workerChannels, workerRoutine(doneCh, batch))
		lB, rB = lB+batchSize, lB

		<-throttle
	}

	resultCh := fanIn(doneCh, workerChannels...)

	result := make(map[int]Payload)
	for v := range resultCh {
		if v.Error != nil {
			return nil, v.Error
		}

		result[v.Result] = v
	}

	return result, nil
}

func fanIn(doneCh <-chan struct{}, channels ...<-chan Payload) <-chan Payload {
	var wg sync.WaitGroup

	multiplexedStream := make(chan Payload, 1)
	multiplex := func(c <-chan Payload) {
		defer wg.Done()
		for i := range c {
			select {
			case <-doneCh:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func retryRemoutFunc(f func(n []int) ([]int, error), args []int) ([]int, error) {
	var err error

	for i := 0; i < retryMax; i++ {
		result, iErr := f(args)
		if iErr != nil {
			err = iErr
			continue
		}

		return result, nil
	}

	return nil, err
}

func remoutFunc(n []int) ([]int, error) {
	lag := rand.Intn(200) + 100

	if lag > 150 {
		return nil, errors.New("Remout error")
	}

	time.Sleep(time.Millisecond * time.Duration(lag))

	result := make([]int, 0, len(n))
	for _, v := range n {
		result = append(result, v*10)
	}

	return result, nil
}
