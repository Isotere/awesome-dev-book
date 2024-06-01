# Awesome Dev Book / GoLang / Concurrency / Patterns

***
[<<<Back](awesome-dev-book/book/Язык%20Go/INDEX.md)
***

## Паттерны конкурентности

- [The for-select loop](FOR_SELECT_LOOP.md)
- [The OR-channel](OR_CHANNEL.md)
- [Pipeline](PIPELINE.md)
- [FAN-IN / FAN-OUT](FAN_IN_FAN_OUT.md)
- [OR-Done Channel](OR_DONE_CHANNEL.md)
- [Tee-channel](TEE_CHANNEL.md)
- [Bridge-channel](BRIDGE_CHANNEL.md)

## Предотвращение утечки горутин (for-select loop)

Горутины очень легковесны и дешевы, но, тем не менее, они требуют ресурсов и не удаляются GC, так что нужно контролировать 
их завершение и очистку, когда больше нет необходимости в них. 

Когда завершаются горутины? 

- когда заканчивают всю свою работу 
- когда они не могут продолжать работу по причине возникновения ошибки
- когда им скажут прекратить свою работу

Первые два - очевидны. Последний надо рассмотреть подробнее. 

Родительская горутина должна иметь возможность сказать производным - останови выполнение. 

Простой пример утечки горутины: 

```go
doWork := func(strings <-chan string) <-chan interface{} {
    completed := make(chan interface{})
	
    go func() {
        defer fmt.Println("doWork exited.")
        defer close(completed)
		
        for s := range strings {
            // Do something interesting
            fmt.Println(s)
        }
    }()
	
    return completed
}

doWork(nil)

fmt.Println("Done.")
```

Мы передали в функцию nil-канал, который передается в горутину. Чтение из nil-канала блокируется навсегда. 

Мы можем разрешить это с помощью сигнала, который основная горутина может послать в производную для указания, чтобы та завершилась. 

```go
doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
    terminated := make(chan interface{}) 
	
    go func() {
        defer fmt.Println("doWork exited.")
        defer close(terminated)
		
        for {
            select {
                case s := <-strings:
                    // Do something interesting
                    fmt.Println(s)
                case <-done:
                    return
            }
        }
    }()
	
    return terminated
}

done := make(chan interface{})

terminated := doWork(done, nil)

go func() {
    // Cancel the operation after 1 second.
    time.Sleep(1 * time.Second)
	fmt.Println("Canceling doWork goroutine...")
	
	close(done)
}()

<-terminated

fmt.Println("Done.")
```