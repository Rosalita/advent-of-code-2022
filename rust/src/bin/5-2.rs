use std::fs;

// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/5.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    let mut contents_lines: Vec<&str> = Vec::new();
    let mut move_lines: Vec<&str> = Vec::new();
    let mut num_stacks: usize = 0;

    // Parse the input data
    for line in input.lines() {
        // Ignore empty lines.
        if line.is_empty() {
            continue;
        }
        // If the line contains "[" then it is a contents line.
        if line.contains("[") {
            contents_lines.push(&line);
            continue;
        } // If the line contains "move" then it is a move line.
        if line.contains("move") {
            move_lines.push(&line);
            continue;
        }

        // If the line is none of of these things then it is the number of stacks line.
        // As the maximum number of stacks is 9, can just count how many chars in this line.
        // Can use a filter to convert to Vec<char> and remove the whitespace.
        let stacks: Vec<char> = line.chars().filter(|c| !c.is_whitespace()).collect();
        num_stacks = stacks.len();
    }

    // create a vector of vectors of char to hold all the stacks.
    let mut stacks: Vec<Vec<char>> = Vec::new();
    // create a vector of char to represent each stack.
    for _ in 0..num_stacks {
        let stack: Vec<char> = Vec::new();
        // and push it to the vector containing all stacks.
        stacks.push(stack);
    }

    // Use the content lines to populate the stacks.
    // data needs to be built up from the bottom of the stack
    // so the content lines need to be reversed
    contents_lines.reverse();
    for line in contents_lines {
        // for each line of content, loop for each stack.
        for i in 0..num_stacks {
            // calculate the index of the content for this stack.
            let index = 1 + (i * 4);
            // retrieve the content using the index.
            let content = line.chars().nth(index).unwrap();
            // if there is nothing in the crate, just ignore and continue.
            if content.is_whitespace() {
                continue;
            }
            // push the content of the crate to the stack.
            stacks[i].push(content);
        }
    }

    // work through the move lines modifying the stacks.
    for line in move_lines {
        // Split the line on " " and collect to vec of move instructions.
        let move_instructions: Vec<&str> = line.split(" ").collect();
        let n: usize = move_instructions[1].parse::<usize>().unwrap();
        let from: usize = move_instructions[3].parse::<usize>().unwrap();
        let to: usize = move_instructions[5].parse::<usize>().unwrap();

        // Going to insert crates after any pre-existing crates,
        // so need to get the index of the last pre-existing crate.
        let insert_index = stacks[to - 1].len();
        for _ in 0..n {
            // to move the contents of the crate maintaining order.
            // get the contents
            let content = stacks[from - 1].pop().unwrap();
            // Then insert all contents at the same index.
            stacks[to - 1].insert(insert_index, content);
        }
    }

    // Get the crate from the top of each stack
    let mut result = String::new();
    for i in 0..num_stacks {
        let content = stacks[i].pop().unwrap();
        result.push(content);
    }
    println!("{}", result);
}
