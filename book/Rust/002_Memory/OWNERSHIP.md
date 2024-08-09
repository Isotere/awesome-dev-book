---
tags:
  - Rust/Memory/Ownership
---
RUST регламентирует управление памятью в виде "Владение".

Память может быть либо "moved" (перемещена, отдана другому владельцу), либо "borrowed" (заимствована).

```rust
enum Light {
    Bright, 
    Dull,
}

fn display_light(light: Light) { // Функция получает владение над памятью, на которую указывает light
    match light {
        Light::Bright => println!("Bright"),
        Light::Dull => println!("Dull"),
    }

    // по выходу из функции вся память принадлежащая функции обнуляется
}

fn display_light2(light: &Light) { // Функция НЕ получает владение над памятью, на которую указывает light, а заимствует ее на время
    match light {
        Light::Bright => println!("Bright"),
        Light::Dull => println!("Dull"),
    }

    // по выходу из функции вся память принадлежащая функции НЕ обнуляется
}

fn main() {
    let dull = Light::Dull; // Создали переменную и владелец у нее функци main

    display_light(dull); // Передали переменную в другую функцию, теперь она владелец (участка памяти),а не main
    display_light(dull); // Ошибка, переменная уже не существует ибо moved

    
    let dull2 = Light::Dull; // Создали переменную и владелец у нее функци main

    display_light2(&dull2); // Передали переменную в другую функцию "в долг", владелец все еще main
    display_light2(&dull2); // Ошибки нет
}
```
