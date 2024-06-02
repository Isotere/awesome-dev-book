# Awesome Dev Book / GoLang / Concurrency / Patterns

***
[<<< Back](awesome-dev-book/book/Язык%20Go/Конкурентность/Паттерны/INDEX.md)
***

## Bridge-channel

Последовательность передачи данных в каналах, типа

```go
<-chan <-chan interface{}
```

Отличается от fan-in fan-out, так как обеспечивает упорядоченную передачу данных, даже из разных источников. 

Деструктуризация канала каналов в простой канал - упростит чтение для консьюмера. 

```go
bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
    // Канал, который возвращает все значения из bridge
    valStream := make(chan interface{}) 
	
	go func() {
        defer close(valStream) 
		
        // Обрабатываем все каналы из chanStream
        for {
            var stream <-chan interface{}
            select {
                case maybeStream, ok := <-chanStream:
                    if ok == false {
                        return
                    }
                    stream = maybeStream
                case <-done:
                    return
            }
        }
		
        // Вычитываем значения из текущего канала и отдаем их в мост
        for val := range orDone(done, stream) {
            select {
                case valStream <- val:
                case <-done:
            }
        } 
    }()
	
    return valStream 
}
```

Теперь используем наш bridge для создания фасада над каналом каналов

```go
genVals := func() <-chan <-chan interface{} {
    chanStream := make(chan (<-chan interface{})) 
	
    go func() {
        defer close(chanStream) 
		
        for i := 0; i < 10; i++ {
            // на каждой итерации создаем новый канал
            stream := make(chan interface{}, 1) 
			
            stream <- i
            close(stream)
			
            chanStream <- stream
        }
	}()
	
    return chanStream 
}

for v := range bridge(nil, genVals()) { 
	fmt.Printf("%v ", v)
}
```

