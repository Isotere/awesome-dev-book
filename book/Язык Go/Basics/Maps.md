---
tags:
  - GoLang/Basics/Maps
---
Хранение данных и доступ по ключ.  
Ключ на основе hash map и хранение в стуктуре "корзин" (buckets)

```go
type user struct {
    name     string
    username string
}
// Construct a map set to its zero value,
// that can store user values based on a key of type string.
// Trying to use this map will result in a runtime error (panic).
var users map[string]user
// Construct a map initialized using make,
// that can store user values based on a key of type string.
users := make(map[string]user)
// Construct a map initialized using empty literal construction,
// that can store user values based on a key of type string.
users := map[string]user{}
```

Присвоение и доступ - по ключу

```go
func main() {
    users := make(map[string]user)
    users["Roy"] = user{"Rob", "Roy"}
    users["Ford"] = user{"Henry", "Ford"}
    users["Mouse"] = user{"Mickey", "Mouse"}
    users["Jackson"] = user{"Michael", "Jackson"}
    for key, value := range users {
        fmt.Println(key, value)
    }
}
Output:
Roy {Rob Roy}
Ford {Henry Ford}
Mouse {Mickey Mouse}
Jackson {Michael Jackson}
```

**Порядок, в каком будет итерация по карте - не определен.**

При доступе по ключу - возвращается два значения, само значение и признак того, есть элемент или нет в карте. Если элемента не было - значение для него будет "нулевое" для этого типа.

Для удаление элемента из карты есть встроенная функция

```go
delete(users, "Roy")
```

В качестве ключа можно использовать любой тип, которые можно использовать в hash функции (объекты этих типов можно сравнивать). (слайсы, функции - пример типа, который НЕЛЬЗЯ использовать).