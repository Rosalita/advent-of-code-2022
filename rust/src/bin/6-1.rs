use std::fs;
// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/6.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    let line = input.lines().next().unwrap().as_bytes();

    let mut index: i32 = 0;
    let mut to_check: Vec<u8> = Vec::new();

    // if a value is &T, then it is borrowed, if it is T it is copied or cloned.
    // In earlier solutions I was converting &T to T with iter().map(|x| *x).
    // Apparently idiomatic Rust uses iter().copied() instead of iter().map(|x| *x).
    // There's also iter().cloned() which is the same as iter().map(|x| x.clone())
    // Copy is a bitwise copy that is implicit and inexpensive.
    // Clone is explicit and may be expensive depending on the value being cloned.
    // This looks like there are trade offs to consider, but for now
    // I'm just going to use iter().copied() instead of iter().map(|x| *x).
    for byte in line.iter().copied() {
        to_check.push(byte);

        if to_check.len() < 4 {
            index += 1;
            continue;
        }
        // The borrow checker was complaining here when passing to_check to the function
        // This is because to_check is used in the next iteration of the loop
        // By taking a clone of to_check, to_check remains in existence.
        if is_all_unique(to_check.clone()) {
            println!("{}", index + 1);
            return;
        }

        // remove the first byte as it has now been checked.
        to_check.remove(0);
        index += 1;
    }
}

fn is_all_unique(mut vals: Vec<u8>) -> bool {
    let start_len = vals.len();
    vals.sort();
    vals.dedup();
    return start_len == vals.len();
}
