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

use std::collections::HashMap;

fn do_map(){
    let mut map = HashMap::new();
    map.insert(1, 2);
    println!("map :{:?}", map);

    let teams  = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];

    let mut scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
    println!("scores map :{:?}", scores);

    for (key, value) in &scores {
        println!("key:{}: value: {}", key, value);
    }
    let team_name = String::from("Blue");

    println!{"team name : {:?}", scores.get(&team_name)};


    let text = "hello world wonderful world";

    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(10);
        //println!("word: {}", word);
        *count += 1;
    }

    println!("{:?}", map);
    //
    let mut s = String::from("你好");
    s.push_str(", Bruce Li!");
    s.push('耶');
    println!("{}", s);

    let s1 = String::from("Rust, ");
    let s2 = String::from("faster!");
    //// note s1 has been moved here and can no longer be used
    let s3 = s1 + &s2;

    println!("s3：{}", s3);
    do_string();
}

fn do_string(){
    let s1 = String::from("tic");
    let s2 = String::from("tac");
    let s3 = String::from("toe");
    let s = s1 + "-" + &s2 + "-" + &s3;
    println!("s: {}", s);

    let s4 = String::from("suffix!");
    let  s = format!("{}-{}-{}", s2, s3, s4);
    println!("s: {}", s);
    //.bytes()   //raw number
//    for c in s.chars() {
//        println!("{}", c);
//    }
}

fn do_err(){
    use std::fs::File;
    //other way: let f = File::open("hello.txt").unwrap();
    //let f = File::open("hello.txt").expect("Failed to open hello.txt");
    let f = File::open("readme.md");

    let f = match f {
        Ok(file) => file,
        Err(error) => {
            panic!("Problem opening the file: {:?}", error)
        },
    };

    //A Shortcut for Propagating Errors: the ? Operator
}

fn largest(list: &[i32]) -> i32 {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

//Another way we could implement largest is for the function to
// return a reference to a T value in the slice. I
fn get_gt<T: PartialOrd + Copy >(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

struct Point<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point<T, U> {
    fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

fn do_trait(){
    let number_list = vec![34, 50, 25, 100, 65];

    let result = get_gt(&number_list);
    println!("The largest number is {}", result);

    let char_list = vec!['y', 'm', 'a', 'q'];

    let result = get_gt(&char_list);
    println!("The largest char is {}", result);
}

fn do_generic(){
    let number_list = vec![34, 50, 25, 100, 65];

    let result = largest(&number_list);
    println!("The largest number is {}", result);

    let number_list = vec![102, 34, 6000, 89, 54, 2, 43, 8];

    let result = largest(&number_list);
    println!("The largest number is {}", result);

    let p1 = Point { x: 5, y: 10.4 };
    let p2 = Point { x: "Hello", y: 'c'};

    let p3 = p1.mixup(p2);

    println!("p3.x = {}, p3.y = {}", p3.x, p3.y);
    do_trait()
}

fn do_closure(){
    let v1 = vec![1, 2, 3];

    let v1_iter = v1.iter();

    let total: i32 = v1_iter.sum();

    assert_eq!(total, 6);

    let v1: Vec<i32> = vec![1, 2, 3];

    let v2: Vec<_> = v1.iter().map(|x| x + 1).collect();
    assert_eq!(v2, vec![2, 3, 4]);

    guessing_number::run_shoes_test();

    guessing_number::calling_next_directly();
}

fn do_smart_p(){
    let x = 5;
    let y = &x;

    assert_eq!(5, x);
    assert_eq!(5, *y);

    let x1 = 5;
    let y1 = Box::new(x);

    assert_eq!(5, x1);
    assert_eq!(5, *y1);
}

fn do_concurrency(){
    use std::thread;
    use std::time::Duration;

    let handle =thread::spawn(|| {
        for i in 1..6 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));
        }
    });

    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }

    handle.join().unwrap();
    do_concurrency1();
}

fn do_concurrency1(){
    use std::thread;
    use std::sync::mpsc;
    use std::time::Duration;
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let vals = vec![
            String::from("你好！"),
            String::from("你去做什么？"),
            String::from("Why？"),
            String::from("那很好呀！"),
        ];

        for val in vals {
            tx.send(val).unwrap();
//            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {}", received);
    }
    do_concurrency2();
    do_concurrency3();
    do_match()
}

fn do_match_p(){
    println!("one");
}


fn do_match(){
    let x = 1;

    match x {
        1 => do_match_p(),
        2 => println!("two"),
        3 => println!("three"),
        _ => println!("anything"),
    }
    //Matching Named Variables
    let x = Some(5);
    let y = 10;

    match x {
        Some(50) => println!("Got 50"),
        Some(y) => println!("Matched, y = {:?}", y),
        _ => println!("Default case, x = {:?}", x),
    }

    println!("at the end: x = {:?}, y = {:?}", x, y);

    let x = 1;

    match x {
        1 | 2 => println!("one or two"),
        3 => println!("three"),
        _ => println!("anything"),
    }

    let x = 2;

    match x {
        1...5 => println!("one through five"),
        _ => println!("something else"),
    }

    let x = 'A';

    match x {
        'a'...'j' => println!("early ASCII letter"),
        'k'...'z' => println!("late ASCII letter"),
        'A'...'Z' => println!("UP ASCII letter"),
        _ => println!("something else"),
    }

    //Destructuring to Break Apart Values
    let p = Point { x: 0, y: 7 };

    match p {
        Point { x, y: 0 } => println!("On the x axis at {}", x),
        Point { x: 0, y } => println!("On the y axis at {}", y),
        Point { x, y } => println!("On neither axis: ({}, {})", x, y),
    }

    let msg = Message::ChangeColor(Color::Hsv(0, 160, 255));

    match msg {
        Message::ChangeColor(Color::Rgb(r, g, b)) => {
            println!(
                "Change the color to red {}, green {}, and blue {}",
                r,
                g,
                b
            )
        },
        Message::ChangeColor(Color::Hsv(h, s, v)) => {
            println!(
                "Change the color to hue {}, saturation {}, and value {}",
                h,
                s,
                v
            )
        }
        _ => ()
    }

    //bind
    do_match1();
    //Rust's unsafe code
    do_unsafe();
}

//Rust unsafe code demo
fn do_unsafe(){
    //doesn’t enforce these memory safety guarantees.
    //Gaining extra superpowers.
    //You can take four actions in unsafe Rust
        //Dereference a raw pointer
        //Call an unsafe function or method
        //Access or modify a mutable static variable
        //Implement an unsafe trait

}

fn do_match1(){
    let msg = MessageNum::Hello { id: 5 };

    match msg {
        MessageNum::Hello { id: id_variable @ 3...7 } => {
            println!("Found an id in range: {}", id_variable)
        },
        MessageNum::Hello { id: 10...12 } => {
            println!("Found an id in another range")
        },
        MessageNum::Hello { id } => {
            println!("Found some other id: {}", id)
        },
    }
}
enum MessageNum {
    Hello { id: i32 },
}
enum Color {
    Rgb(i32, i32, i32),
    Hsv(i32, i32, i32),
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(Color),
}
//Similarities Between RefCell<T>/Rc<T> and Mutex<T>/Arc<T>
fn do_concurrency2(){
    use std::thread;
    use std::sync::mpsc;
    use std::time::Duration;
    let (tx, rx) = mpsc::channel();

    let tx1 = mpsc::Sender::clone(&tx);
    thread::spawn(move || {
        let vals = vec![
            String::from("1:你好！"),
            String::from("1:你去做什么？"),
            String::from("1:Why？"),
            String::from("1:那很好呀！"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("2:你好！"),
            String::from("2:你去做什么？"),
            String::from("2:Why？"),
            String::from("2:那很好呀！"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {}", received);
    }
}

fn do_concurrency3(){
    use std::sync::{Mutex, Arc};
    use std::thread;

    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        println!("Result: {}", *counter.lock().unwrap());
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
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
    do_map();
    do_err();
    do_generic();
    do_closure();
    do_smart_p();
    do_concurrency();
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
