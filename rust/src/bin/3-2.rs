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

    let mut line1: &str = "";
    let mut line2: &str = "";
    let mut line3: &str = "";

    // Use enumerate to get the index of the for loop.
    for (index, line) in input.lines().enumerate() {

        // save three lines to make up a group of elves.
        if index % 3 == 0 {
            line1 = line;
            continue;
        }
        if index % 3 == 1 {
            line2 = line;
            continue;
        }
        if index % 3 == 2 {
            line3 = line;
        }

        // Create three hash sets of characters from each line.
        let first_elf: HashSet<char> = line1.chars().collect();
        let second_elf: HashSet<char> = line2.chars().collect();
        let third_elf: HashSet<char> = line3.chars().collect();

        // create the first intersection, we want to generate HashSet<char> so can compare with a third elf
        // To avoid creating HashSet<&char> each char needs to be dereferenced, map can do this.
        let intersection1: HashSet<char> = first_elf.intersection(&second_elf).map(|x: &char| *x).collect();

        let intersection2: Vec<&char> = intersection1.intersection(&third_elf).collect::<Vec<&char>>();

        // We know there is only one item which is carried by all three elves.
        // This item is the first item in the second.
        items.push(*intersection2[0]);
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
