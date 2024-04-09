//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// Ограничение - один раз в секунду
const rateLimit = time.Second

func Crawl(url string, depth int, wg *sync.WaitGroup, throttle <-chan time.Time) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Читаем из тикера - будет приходить с периодичностью rateLimit, пока ничего нет - блокируемся в этом месте
	<-throttle
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, wg, throttle)
	}
	return
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg, time.Tick(rateLimit))
	wg.Wait()
}
