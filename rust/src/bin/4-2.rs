use std::collections::HashSet;
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

    let mut total_overlap = 0;

    for line in input.lines() {
        let assignments = line // start with the input line.
            .split(",") // split the input on , to get both elves assignments
            .map(|a| expand_assignment(a)) // expand each assignment
            .collect::<Vec<Vec<i32>>>(); // collect both assignments back into vector of string.

        // convert both elfs assignments to hash sets of i32.
        let first_elf: HashSet<i32> = assignments[0].iter().map(|x: &i32| *x).collect();
        let second_elf: HashSet<i32> = assignments[1].iter().map(|x: &i32| *x).collect();

        // get the intersections between both elfs assignments
        let intersection: HashSet<i32> = first_elf
            .intersection(&second_elf)
            .map(|x: &i32| *x)
            .collect();

        // if there are any intersecting values, increment the overlap counter.
        if intersection.len() > 0 {
            total_overlap += 1;
        }
    }

    println!("{}", total_overlap);
}

// expand_assignment converts and assignment from format "2-5"
// to vector of integers format [2,3,4,5,]
fn expand_assignment(assignment: &str) -> Vec<i32> {
    // convert the start and end sections to integers.
    let secs = assignment
        .split("-")
        .map(|s| s.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    // initialise a new vector
    let mut full_assignment: Vec<i32> = Vec::new();

    // for in loops  ill exclude the final value unless = is added.
    // loop from the start section to the end section
    for i in secs[0]..=secs[1] {
        // save the integer.
        full_assignment.push(i);
    }
    full_assignment
}
