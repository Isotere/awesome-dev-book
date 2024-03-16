fn main() {
    // Immutable
    let n = 32;
    // n += 1; // Does not compile
    println!("Immutable n {n}");

    // Mutable
    let mut m = 32;
    m += 1; // Works
    println!("Mutable m {m}");

    // Shadowing
    let m = m + 4;
    println!("Shadowed m {m}");

    {
        // Shadowing in custom block
        let m = 5;
        println!("Shadowed in block m {m}");
    }
    println!("Original m {m}");

    // Unit type
    let u: () = {
        let a = 12;
    };
}
