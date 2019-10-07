//! #My Crate
//!
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.
//!
//!
#[derive(PartialEq, Debug)]
struct Shoe {
    size: u32,
    style: String,
}

fn shoes_in_my_size(shoes: Vec<Shoe>, shoe_size: u32) -> Vec<Shoe> {
    shoes.into_iter()
        .filter(|s| s.size == shoe_size)
        .collect()
}


pub fn run_shoes_test(){
    let shoes = vec![
        Shoe { size: 10, style: String::from("sneaker") },
        Shoe { size: 13, style: String::from("sandal") },
        Shoe { size: 10, style: String::from("boot") },
    ];

    let in_my_size = shoes_in_my_size(shoes, 10);

    assert_eq!(
        in_my_size,
        vec![
            Shoe { size: 10, style: String::from("sneaker") },
            Shoe { size: 10, style: String::from("boot") },
        ]
    );
}

struct Counter {
    count: u32,
}

impl Counter {
    fn new() -> Counter {
        Counter { count: 0 }
    }
}

impl Iterator for Counter {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        self.count += 1;

        if self.count < 6 {
            Some(self.count)
        } else {
            None
        }
    }
}

pub fn calling_next_directly() {
    let mut counter = Counter::new();

    assert_eq!(counter.next(), Some(1));
    assert_eq!(counter.next(), Some(2));
    assert_eq!(counter.next(), Some(3));
    assert_eq!(counter.next(), Some(4));
    assert_eq!(counter.next(), Some(5));
    assert_eq!(counter.next(), None);

}



/// Adds one to the number given.
///
/// # Examples
///
/// ```
/// let arg = 5;
/// let answer = guessing_number::add_one(arg);
///
/// assert_eq!(6, answer);
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}


#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn do_cmd() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn do_holder(){
        let larger = Rectangle { width: 8, height: 7 };
        let smaller = Rectangle { width: 5, height: 1 };

        assert!(larger.can_hold(&smaller));
    }

    #[test]
    fn test_add(){
        assert_eq!(4, add_two(2));

        assert_ne!(56, add_two(2));
    }

    #[test]
    fn it_works() -> Result<(), String> {
        if 2 + 2 == 4 {
            Ok(())
        } else {
            Err(String::from("two plus two does not equal four"))
        }
    }
    //1. If we want to see printed values for passing tests as well,
    //we can disable the output capture behavior by using the --nocapture flag:

    //2. he tests are running in parallel, as we talked about in the previous section.
    // Try using the --test-threads=1

    //3. check the results of the ignored tests and you have time to wait for the
    // results, you can run cargo test -- --ignored instead.

    #[test]
    fn internal() {
        assert_eq!(4, internal_adder(2, 2));
    }
}

fn internal_adder(a: i32, b: i32) -> i32 {
    a + b
}


pub fn add_two(a: i32) -> i32 {
    a + 2
}


#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

