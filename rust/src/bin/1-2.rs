use std::fs;

// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/1.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    // Calories_per_elf is a vector which holds the total calories for each elf.
    let mut calories_per_elf: Vec<i32> = Vec::new();
    // Total is a counter used to add calories together.
    let mut total = 0;

    for line in input.lines() {
        // if the option returned from parsing the line as i32 has Some calories value
        if let Some(calories) = line.parse::<i32>().ok() {
            // add the calories to the total.
            total += calories
        }

        // if the line is empty, this is the end of data for the current elf
        if line.is_empty() {
            // push the total onto the calories_per_elf vector
            calories_per_elf.push(total);
            // reset the total
            total = 0;
        }
    }
    // Push the total for the final elf.
    calories_per_elf.push(total);

    // sort changes the order of elements in vec from smallest to largest.
    calories_per_elf.sort();
    // reverse switches the order of the elements so they are backwards.
    calories_per_elf.reverse();

    // print the sum of the first three elements in the vec.
    println!("{}", calories_per_elf[..3].iter().sum::<i32>());
}
