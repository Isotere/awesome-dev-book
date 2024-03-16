# Awesome Dev Book / Rust / Create a library

***
[<<<Back](../INDEX.md)
***

## Создание библиотеки (crate)

Создаем следующим образом: 

```Bash
cargo new --lib authentication
```

Создастся библиотека с одним публичным методом 

```Rust
pub fn add(left: usize, right: usize) -> usize {
    left + right
}
```

и юнит тестом: 

```Rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
```

```#[cfg(test)]``` говорит нам, что ниже модуль с юнит тестами

```#[test]``` - что ниже конкретный юнит тест

```use super::*``` - импортирует все функции из выше-лежащего модуля

Запустить тесты можно командой: 

```Bash
cargo test
```

Для того, что использовать библиотеку в приложении нужно: 

1. Добавить зависимость для нее 
    ```Toml
    [dependencies]
    authentication = {path = "../authentication"}
    ```
2. В нужном файле программы добавить импорт
    ```Rust
    use authentication::greet_user;
    ```
3. Для того, чтобы можно было импортировать методы - нужно, чтобы они были объявлены как публичные



