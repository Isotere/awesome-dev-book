# Awesome Dev Book / GoLang / Concurrency / Patterns

***
[<<< Back](./PATTERNS.md)
***

## Pipeline

Для работы со стримами или пачками данных. 

По-простому - это набор функций, которые получают данные, обрабатывают их и отправляют дальше по цепочке. 

[Простая реализация](../../../../code/go_lang/concurency/patterns/pipeline_1/main.go)
[Еще одна реализация](../../../../code/go_lang/concurency/patterns/pipeline_2/main.go)

Мы используем в параметрах interface{} - это чтобы можно было работать с любым типом, но вполне можно использовать конкретные типы, либо generic. 

### Генераторы

Генератор - это любая функция, которая конвертирует набор дискретных значений в поток значений в канале. 

```go
repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
    valueStream := make(chan interface{}) 

    go func() {
        defer close(valueStream) 
	
        for {
            for _, v := range values {
                select {
                    case <-done: return
                    case valueStream <- v:
                }
            }
        }
    }()

    return valueStream
}

```

