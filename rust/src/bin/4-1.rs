use std::fs;

// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/4.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    let mut total_overlap: i32 = 0;

    for line in input.lines() {
        let assignments = line // start with the input line.
            .split(",") // split the input on , to get both elves assignments
            .map(|a| expand_assignment(a)) // expand each assignment
            .collect::<Vec<String>>(); // collect both assignments back into vector of string.

        // if the first elfs assignment is larger than the second elfs assignment
        if assignments[0].len() > assignments[1].len() {
            // check if the first elfs assignment contains the second elfs assignment.
            if assignments[0].contains(&assignments[1][..]) {
                // if it does, increment the counter
                total_overlap += 1;
            }
        } else {
            // check if the second elfs assigment contains the first elfs assignment.
            if assignments[1].contains(&assignments[0][..]) {
                // if it does, increment the counter
                total_overlap += 1;
            }
        }
    }

    println!("{}", total_overlap);
}

// expand_assignment converts and assignment from format "2-5"
// to a longer string format "2,3,4,5,"
fn expand_assignment(assignment: &str) -> String {
    // convert the start and end sections to integers.
    let secs = assignment
        .split("-")
        .map(|s| s.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    // initialise a new string to hold the full_assignment
    let mut full_assignment = String::new();

    // for in loops  ill exclude the final value unless = is added.
    // loop from the start section to the end section
    for i in secs[0]..=secs[1] {
        // convert the index to string
        let s = i.to_string();

        // numbers less than 10 need to be prefixed with a 0.
        if i < 10 {
            full_assignment.push_str("0");
        }
        full_assignment.push_str(&s);
        // numbers also need to be separated.
        full_assignment.push_str(",");
    }
    full_assignment
}
