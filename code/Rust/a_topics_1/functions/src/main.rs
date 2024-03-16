
fn main() {
    let n = 5;
    println!("{:?}", double(n));
}

fn double(n: i32) -> i32 {
    n * 2
}

fn triple(n:i32) -> i32 {
    return n * 3;
}

fn read_line() -> String {
    let mut input = String::new();
    std::io::stdin().read_line(&mut input).expect("Stdin does not working");

    // избавляемся от завершающего перевода строки и возвращаем
    input.trim().to_string()
}
