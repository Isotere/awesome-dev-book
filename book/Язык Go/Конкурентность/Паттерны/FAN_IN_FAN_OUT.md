# Awesome Dev Book / GoLang / Concurrency / Patterns

---

[<<< Back](PATTERNS.md)

---

## FAN-IN / FAN-OUT

### FAN-OUT

При реализации паттерна pipeline иногда возникает ситуация, когда один из стейджов будет требовать много ресурсов и выполняться долго. И таким
образом увеличится время обработки всех данных из генератора в потоке.

Можно в ресурсо-емких местах распараллелить работу.

Для этого результат должен быть порядко-независимый. Как и сами стейджи.

Например так:

```go
numFinders := runtime.NumCPU()

finders := make([]<-chan int, numFinders)

for i := 0; i < numFinders; i++ {
    finders[i] = primeFinder(done, randIntStream)
}
```

### FAN-IN

После применения FAN-OUT появляется другая проблема - в итоге нам нужен один канал, откуда мы будем читать.

Например так

```go
fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
    var wg sync.WaitGroup
    multiplexedStream := make(chan interface{})

    multiplex := func(c <-chan interface{}) {
        defer wg.Done()

        for i := range c {
            select {
                case <-done: return
                case multiplexedStream <- i:
            }
        }
    }

    // Select from all the channels
    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }

    // Wait for all the reads to complete
    go func() {
        wg.Wait()
        close(multiplexedStream)
    }()

    return multiplexedStream
}
```

