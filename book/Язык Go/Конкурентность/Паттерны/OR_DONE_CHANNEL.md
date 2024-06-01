# Awesome Dev Book / GoLang / Concurrency / Patterns

***
[<<< Back](PATTERNS.md)
***

## OR-Done Channel

Для случая, когда мы хотим отслеживать, что канал, из которого мы читаем - закрылся.
(иначе из закрытого канала нам будет приходить zero-value, а не факт, что оно нам нужно)

Пример: 

```go
orDone := func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <-v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

for val := range orDone(done, myChan) {
	...
}

```

