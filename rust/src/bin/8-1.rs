use std::fs;
// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/8sample.txt").expect("unable to read input");
    input
}

fn main() {
    let input = get_input();

    let num_lines = input.lines().count();
    let mut visible: usize = 0;

    for (index, line) in input.lines().enumerate() {
        // the first row of trees is an edge, so all trees on that row are visible.
        if index == 0 {
            visible += line.len();
            continue;
        }
        // the last row of trees is an edge, so all tress on that row are visible.
        if index == (num_lines - 1) {
            visible += line.len();
            continue;
        }

        let mut chars = line.chars();

        // remove the first char as that is a visible tree
        chars.next();
        // remove the last char as that is also a visible tree
        chars.next_back();
        // count the visble trees.
        visible += 2;

        // need to start looking at tree sizes now, so parse the chars to int
        let trees: Vec<u32> = chars.map(|c| c.to_digit(10).unwrap()).collect();
        println!("{:?}", trees);

        // need a matrix of trees to check each tree. 
        // Should probably parse all the trees first
        // get them into some kind of forrest data structure.
        // Rather than trying to establish visibility on a line by line basis.
    }

    println!("{}", visible);
}
