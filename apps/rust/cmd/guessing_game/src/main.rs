use std::io;
use std::cmp::Ordering;
use rand::Rng;

fn guess_num(){
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    loop{
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin().read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse(){
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }

    println!("The secret number is: {}", secret_number);
}

fn do_compound(){
    let x: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = x.0;

    let six_point_four = x.1;

    let one = x.2;
    println!("five_hundred: {}, six_point_four:{}, other:{}", five_hundred, six_point_four, one);

    let a: [i32; 5] = [1, 2, 3, 4, 5];
    println!(" Array element :{}", a[0]);
}

fn do_float(){
    let x = 2.0; // f64

    let y: f32 = 3.0; // f32
    println!("x:{}, y:{} ", x, y);

    do_compound();

    //expression
    println!("zero number ; {}", zero_plus(23));

    let a = [10, 20];

    for element in a.iter() {
        println!("the value is: {}", element);
    }

    for number in (1..4).rev() {
        print!("{}!", number);
    }
}

fn zero_plus(i: i32) -> i32 {
     0 + i
}

fn main() {
    //mut and default immutable
    let mut i = 0;
    println!("init i :{}", i);
    i = 100;
    println!("change i: {}", i);

    //shadowing
    let x = 5;
    let x = x + 1;
    let x = x * 2;
    println!("The value of x is: {}", x);

    let spaces = "   ";
    let spaces = spaces.len();
    println!("space number :{}", spaces);

    // floating-point numbers
    do_float();

    //guess_num()
}
