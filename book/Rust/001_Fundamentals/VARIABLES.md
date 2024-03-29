# Awesome Dev Book / Rust / Fundamentals

***
[<<<Back](./INDEX.md)
***

## Переменные

Объявляются с помощью ключевого слова _let_:

```rust
fn main() {
    let a = 32;
    let b = "some string";
}
```

По умолчанию все переменные Immutable (неизменяемые). Для того, чтобы переменные можно было изменять нужно добавить ключевое слово _mut_:

```rust
fn main() {
    let mut a = 32;
    
    a += 1;
}
```

## Скрытие (shadowing) переменных и области видимости

Можно повторно объявить переменную с таким же названием в том же блоке кода, при этом она скроет объявление предыдущей переменной до конца области видимости данного блока кода.

```rust
    // Mutable
    let mut m = 32;
    m += 1; // Works
    println!("Mutable m {m}");

    // Shadowing
    let m = m + 4;
    println!("Shadowed m {m}");
```

Можно создать внутренний блок, скрыть переменную, но после выхода из блока она вернется к предыдущему значению.

```rust
    let m = 4;
    println!("Original m {m}");

    {
        // Shadowing in custom block
        let m = 5;
        println!("Shadowed in block m {m}");
    }
    println!("Original m {m}");
```

## Unit type

в Rust есть тип Unit - который означает буквально "ничего". Ближайший аналог тип void из С/С++. Т.е. тот случай, когда выражение ничего не возвращает.

```rust
    // Unit type
    let u: () = {
        let a = 12;
    };
```

