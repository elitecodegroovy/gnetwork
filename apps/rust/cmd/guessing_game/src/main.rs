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

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

fn string_slice(){
    let my_string = String::from("Rust Async");

    // first_word works on slices of `String`s
    let _word = first_word(&my_string[..]);

    let my_string_literal = "Rust Async";

    // first_word works on slices of string literals
    let _word = first_word(&my_string_literal[..]);

    // Because string literals *are* string slices already,
    // this works too, without the slice syntax!
    let _word = first_word(my_string_literal);
    println!(" word: {}", _word)
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

    let s = String::from("The Rust Programming Language");
    let s1 = &s;
    let s2 =&s;
    println!("s1: {}, s2: {}", s1, s2);
    let  s3 = &s;
    println!("s3: {}", s3);

    string_slice();
    do_struct();
}

fn zero_plus(i: i32) -> i32 {
     0 + i
}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

//fn area(r: &Rectangle) -> u32 {
//    r.height * r.width
//}

impl Rectangle {
    fn area(&self) -> u32 {
        self.height * self.width
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }

    fn square(size: u32) -> Rectangle {
        Rectangle { width: size, height: size }
    }
}


fn do_struct(){
    let rect1 = Rectangle { width: 20, height: 50 };
    let rect2 = Rectangle { width: 10, height: 40 };
    let rect3 = Rectangle { width: 60, height: 45 };

    println!("rect1 area: {}", rect1.area());
    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));

    println!("rect1: {:?}", &(Rectangle::square(3)));
//    println!(
//        "The area of the rectangle is {} square pixels.",
//        area(&rect1)
//    );
//    println!("rect1: {:?}", &rect1);
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
