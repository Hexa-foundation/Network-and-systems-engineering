//include the rand crate in the Cargo.toml then update, build & run
extern crate rand;
use rand::Rng;
use std::cmp::Ordering;

fn main(){
println!("Guess the number");



println!("Please input your guess");

let mut guess = String::new();

std::io::stdin().read_line(&mut guess)
                .expect("Failed to read line");
                
//generates a random number                
let secret_number = rand::thread_rng().gen_range(1, 101);

//converts the String that it reads as input to a real number
//we shadow variable guess
let guess: u32 = guess.trim().parse()
                 .expect("Please type a number");
println!("The secret number is: {}", secret_number);
println!("You guessed:{}",guess);

//compares cmp to secret number
match guess.cmp(&secret_number)
{
Ordering::Less => println!("too small!"),
Ordering::Greater => println!("Too large!"),
Ordering::Equal => println!("You win!"),
}
}
