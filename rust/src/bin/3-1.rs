use std::collections::HashSet;
use std::fs;

// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/3.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    let mut items: Vec<char> = Vec::new();

    for line in input.lines() {
        // &str is a string literal and can't be indexed with an integer.
        // so need to convert to type String, apparently to_owned does this.
        let line_string = line.to_owned();

        let mid = line.len() / 2;

        // split the string in the middle.
        let first_rucksack = &line_string[..mid];
        let second_rucksack = &line_string[mid..];

        // a hash set is like a hash map but it only cares about keys.
        // a hash set is guaranteed to have no duplicate elements.
        // converting both rucksacks to hash sets of char will remove any duplicate chars.
        let first_set: HashSet<char> = first_rucksack.chars().collect();
        let second_set: HashSet<char> = second_rucksack.chars().collect();

        // hash sets support an operation called intersection.
        // intersection will get all the elements which are in both hash sets.
        let intersection = first_set.intersection(&second_set).collect::<Vec<&char>>();

        // We know there is only one item which is in both rucksacks.
        // so it will be the first of the intersection items.
        items.push(*intersection[0]);
    }

    // start a counter to add up all the priorities.
    let mut priority_total: i32 = 0;

    // calculate the priority for each item
    let items_iter = items.iter();
    for item in items_iter {
        let priority = calc_priority(*item);
        priority_total += priority as i32;
    }

    println!("{}", priority_total);
}

fn calc_priority(c: char) -> u8 {
    let byte = c as u8;
    if byte > 96 {
        // lowercase bytes are 97 - 122
        // to transform these values to 1 - 26, substract 96
        return byte - 96;
    } else {
        // uppercase bytes are 65 - 90
        // to transform these values to 27 - 52, subtract 38.
        return byte - 38;
    }
}
