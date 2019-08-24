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

