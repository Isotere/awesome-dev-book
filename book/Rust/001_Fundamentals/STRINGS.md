# Awesome Dev Book / Rust / Fundamentals

***
[<<<Back](./INDEX.md)
***

## Строки

Строки и в Rust  строки - храним текстовую информацию.

Есть два типа строк:

String - owned
&str     - borrowed String slice

Для использования в структуре - String
Для передачи в функцию - &str

```rust
fn print_it(s &str) {
    println!("{:?}", s);
}

let owned_string = "owned_string".to_owned();
let another_owned = String::from("another");

print_it("string_slice");
print_it(&owned_string);


struct Employee {
    name: String,
}
```

### Чтение с консоли

```rust
fn read_line() -> String {
    let mut input = String::new();
    std::io::stdin().read_line(&mut input).expect("Stdin does not working");

    // избавляемся от завершающего перевода строки и возвращаем
    input.trim().to_string()
}
```