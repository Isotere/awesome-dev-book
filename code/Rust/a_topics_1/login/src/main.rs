use authentication::{login, read_line};

fn main() {
    println!("Input username: ");
    let username = read_line();
    println!("Input password: ");
    let password = read_line();

    if login(&username, &password) {
        println!("You are in!");
    } else {
        println!("Username or password is incorrect!");
    }
}
