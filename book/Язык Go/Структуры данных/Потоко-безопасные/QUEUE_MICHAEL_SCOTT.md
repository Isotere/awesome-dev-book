# Awesome Dev Book / GoLang / Data Structs

***
[<<<Back](awesome-dev-book/book/Язык%20Go/INDEX.md)
***

## lock-free очередь Майкла Скотта 

Потоко-безопасная очередь (реализация Майлка Скотта)

[Ссылка на код](awesome-dev-book/code/go_lang/data_structs/queue_scott/main.go)

Одно из правил: 

> Если мы что-то делаем больше, чем одним CAS-ом, мы должны заставить какую-то другую горутину что-то сделать

