---
tags:
  - GoLang/Testing/Benchmarks
---
Создать файл, как любой тестовый файл с постфиксом \_test.go .

Создать функцию с префиксом Benchmark

```go
package main

import "testing"


func myFunction(){
  for i := 0; i < 1000000; i++ {
    fmt.Println(i)
  }
}

func BenchmarkMyFunction(b *testing.B) {
  for i := 0; i < b.N; i++ {
    myFunction()
  }
}

```

Запуск: 

```sh
go test -bench=. -benchmem -benchtime=10s -count=5
```

- bench - какие именно тесты запускать (точка - все)
- benchmem - считать используемую память 
- benchtime - сколько по времени должен длиться запуск
- count - сколько раз запускать тест

Запуск только бенчей, без тестов: 

```sh
go test -bench=. -run=^$
```

Примеры бенчей: 

```go
package main

import (
  "testing"
)

func BenchmarkMyFunction(b *testing.B) {
  m := []myStruct{
    {name: "Kevin", age: 20},
    {name: "Jane", age: 21},
    {name: "John", age: 22},
  }
  b.Run("Method", func(b *testing.B) {
    for i := 0; i < b.N; i++ {
      m[0].randomFunction()
      m[1].randomFunction()
      m[2].randomFunction()
    }
  })
  b.Run("Function", func(b *testing.B) {
    for i := 0; i < b.N; i++ {
      randomFunction(m[0])
      randomFunction(m[1])
      randomFunction(m[2])
    }
  })

}

func randomFunction(m myStruct) {
  // println(m.name)
}

```

```go
var table = []struct {
    input int
}{
    {input: 100},
    {input: 1000},
    {input: 74382},
    {input: 382399},
}

func BenchmarkPrimeNumbers(b *testing.B) {
    for _, v := range table {
        b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                primeNumbers(v.input)
            }
        })
    }
}
```


