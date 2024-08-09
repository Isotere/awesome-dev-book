---
tags:
  - Rust/Essentials
---
# Установка Rust
```make
install-rust:
	curl --proto '=https' --tlsv1.3 https://sh.rustup.rs -sSf | sh
	rustup component add rust-analyzer
	cargo install cargo-get
```

[cargo-get](https://crates.io/crates/cargo-get)

# Создание проекта cargo
```sh
cargo new <project_name>
```

# Сборка проекта
```sh
cargo build
```


# Сборка и запуск
```sh
cargo run
```

