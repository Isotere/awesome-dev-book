---
tags:
  - GoLang/Memory
  - Golang/Optimization
---
## Выделение памяти для структур и ее выравнивание

Допустим у нас есть структура: 

```go
type Example struct {
	flag    bool
	counter int16
	pi      float32
}
```

На первый взгляд нам может показаться, что экземпляр данной структуры займет в памяти 7 байт (1 байт на bool + 2 байта на int16 + 4 байта на float32). 


|                               bool                                | int32                                                                                                      | float32                                                                                                                                                                                      |
| :---------------------------------------------------------------: | ---------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table><tr><td style="border: solid 1px">&nbsp;</td></tr></table> | <table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> | <table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |
|                                                                   |                                                                                                            |                                                                                                                                                                                              |


Но на самом деле это не совсем так. Компилятор будет выполнять __выравнивание__ по памяти. 

И на самом деле картина будет следующей: 

| bool |  | int32 | float32 |
|:----:| ---- | ----- | ------- |
|<table><tr><td style="border: solid 1px">&nbsp;</td></tr></table>|<table><tr><td style="border: solid 1px">&nbsp;</td></tr></table>|<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |

После булева значения добавится еще один байт (Padding - оступ), который выровняет нашу структуру до размера 8 байт. 

Почему это происходит? Потому что компилятор выполняет выравнивание по размеру, кратному  типу. 

Т.е. у нас булево значение, занимает один байт. Мы выделяем для него 1 байт и у нас в слове (8 байт) остается 7 байт. Далее у нас идет двух-байтный тип целого - мы должны выравнивать внутри слова по размеру этого типа (т.е. 0+1 байт, 2+3 и тд). Но у нас один байт занят булевым значением. Т.е. в нулевую позицию мы положить не можем - занято, а в 1 нельзя, поэтому мы сдвигаемся на один байт и начиная со 2-го кладем наше целое. Дальше идет четырех байтной действительное число - оно подчиняется той же логике (0-3, 4-7 байты). Но тут у нас все ОК - в совокупности они занято уже 4 байта, поэтому мы без отступов кладем число в позицию 4. 

В данном случае один байт. А что если у нас вместо двух байтного целого будет 8 байтное? Тогда мы теряем аж 7 байт!  Плюс появляется отступ у действительного 4-х байтного число, ибо чтобы поместилось в слово. Кусок слова другому не смогут отдать.

| bool |  | int64 | float32 | |
|:----:| ---- | ----- | ------- | ---- |
|<table><tr><td style="border: solid 1px">&nbsp;</td></tr></table>|<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table>|<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |<table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table>|

Для того, чтобы немного нивелировать все это, необходимо всегда размещать ОТ большего к меньшему типу. В таком случае "потери" будут только в "хвосте" куска памяти для объекта структуры.

|                                                                                                                                                                              int64                                                                                                                                                                               | float32                                                                                                                                                                                      | bool                                                              |                                                                                                                                                     |
|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| <table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> | <table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> | <table><tr><td style="border: solid 1px">&nbsp;</td></tr></table> | <table><tr><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td><td style="border: solid 1px">&nbsp;</td></tr></table> |

> ==Но необходимо помнить, что все это микро-оптимизации и следить за этим нужно только в том случае, когда это действительно нужно.==

