---
tags:
  - GoLang/Basics/Errors
  - GoLang/Basics/Panic
  - GoLang/Basics/Recover
---
## Паника

При возникновении паники прекращается выполнения функции и (!) запускаются все defer этой функции. 

## Recover

```go
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
```

