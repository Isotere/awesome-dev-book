# Awesome Dev Book / GoLang / Concurrency / Patterns

---

[<<< Back](./PATTERNS.md)

---

## The for-select loop

Все очень просто:

```go
for { // Either loop infinitely or range over something
    select {
        // Do some work with channels
    }
}
```

Сценарии использования.

### Итерация по коллекции и отправка значений в канал

```go
for _, s := range []string{"a", "b", "c"} {
    select {
        case <-done:
            return
        case stringStream <- s:

    }
}
```

### Бесконечная итерация с ожиданием сигнала завершения

```go
for {
    select {
        case <-done: // до тех пор, пока done не закрыт выполняем все, что ниже
            return
        default:
    }

    // Выполняем какую-то работу (можно поместить в блок default)
}
```
