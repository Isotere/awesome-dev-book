---
tags:
  - Rust/Memory/Lifetime
---
Lifetimes позволяют
- использовать "заимствованные" данные в структурах
- возвращать ссылки из функций

Это механизм, позволяющий определять, как долго живут данные.
Проверяется на уровне компилятора.

```rust
struct Name<'a> {
    field: &'a DataType,
}
```

- по договоренности используем 'a, 'b, 'c
- 'static зарезервированно - остаются в памяти пока существует программа

'a указывает программе, что данные должны существовать после того, как структура будет уничтожена.  Говорит, что переменная в поле была создана ранее, чем создан объект структуры.

Структуры ,использующие "заимствованные" данные, должны:
- быть созданы ПОСЛЕ того, как собственник данных создан
- быть уничтожены ДО того, как будет уничтожен собственник данных

```rust
fn name<'a>(arg: &'a DataType) -> &'a DataType {}
```

### Пример Lifetime & Structures

```rust
const MOCK_DATA: &'static str = include_str!("mock-data.csv");

struct Names<'a> {
    // Если указать просто &str будет ошибка компиляции
    inner: Vec<&'a str>,
}

struct Titles<'a> {
    inner: Vec<&'a str>,
}

fn main() {
    let data: Vec<_> = MOCK_DATA.split('\n').skip(1).collect();
    let names: Vec<_> = data
        .iter()
        .filter_map(|line| line.split(',').nth(1))
        .collect();
    let names = Names { inner: names };

    let titles: Vec<_> = data
        .iter()
        .filter_map(|line| line.split(',').nth(4))
        .collect();
    let titles = Titles { inner: titles };

    let data = names.inner.iter().zip(titles.inner.iter());

    for (name, title) in data {
        println!("Name: {}, Title: {}", name, title);
    }
}
```

### Пример Lifetime & Functions

```rust
fn longest<'a>(one: &'a str, two: &'a str) -> &'a str {
    if two > one {
        two
    } else {
        one
    }
}

fn main() {
    let short = "hello";
    let long = "this is a long message";
    println!("{}", longest(short, long))
}
```