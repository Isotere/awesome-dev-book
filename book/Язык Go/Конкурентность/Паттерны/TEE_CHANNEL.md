---
tags:
  - GoLang/SoftwareDesign/Patterns/Concurrency/TeeChannel
---

# Awesome Dev Book / GoLang / Concurrency / Patterns

***
[<<< Back](awesome-dev-book/book/Язык%20Go/Конкурентность/Паттерны/INDEX.md)
***

## Tee-channel

Когда нам нужно данные из одного входного потока направить сразу в несколько других горутин. 

```go
tee := func(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
    out1 := make(chan interface{})
    out2 := make(chan interface{})

    go func() {
        defer close(out1)
        defer close(out2)
		
        for val := range orDone(done, in) {
			// Мы хотим использовать локальные версии каналов поэтому скрываем глобальные
			var out1, out2 = out1, out2
			
            for i := 0; i < 2; i++ {
                select {
				    case <-done: 
                    case out1 <- val:
						// Обнуляем, чтобы в каждый канал записалось только один раз - запись в nil канал блок навсегда
                        out1 = nil
                    case out2 <- val:
                        out2 = nil
                }
            }
        }
    }()
	
    return out1, out2
}

```