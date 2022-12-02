use std::fs;

// get_input returns raw string input read directly from the input file.
// Docs for std library String are at: https://doc.rust-lang.org/std/string/struct.String.html
fn get_input() -> String {
    // expect does the same thing as unwrap, except it makes panic nicer by showing a message.
    let input = fs::read_to_string("input/2.txt").expect("unable to read input");
    input
}

#[derive(PartialEq)]
enum Shape {
    Rock,
    Paper,
    Scissors,
    Unknown,
}

struct Round {
    opponent: Shape,
    player: Shape,
}

fn main() {
    let input = get_input();
    let mut score: i32 = 0;

    for line in input.lines() {
        // First characters in the line is the opponents input.
        let opponent_input = line.chars().next().expect("unable to read opponent input");

        // Third character in the line is the players input.
        let player_input = line.chars().nth(2).expect("unable to read player input");

        let mut this_round = Round {
            opponent: Shape::Unknown,
            player: Shape::Unknown,
        };

        match opponent_input {
            'A' => this_round.opponent = Shape::Rock,
            'B' => this_round.opponent = Shape::Paper,
            'C' => this_round.opponent = Shape::Scissors,
            _ => this_round.opponent = Shape::Unknown,
        }

        match player_input {
            'X' => this_round.player = Shape::Rock,
            'Y' => this_round.player = Shape::Paper,
            'Z' => this_round.player = Shape::Scissors,
            _ => this_round.player = Shape::Unknown,
        }

        let points = score_round(this_round);

        score += points;
    }
    println!("{}", score);
}

fn score_round(round: Round) -> i32 {
    let mut points: i32 = 0;
    // calculate the points for the players shape.
    match round.player {
        Shape::Rock => points += 1,
        Shape::Paper => points += 2,
        Shape::Scissors => points += 3,
        _ => {}
    }

    // add three points if the round is a draw.
    if round.opponent == round.player {
        points += 3;
    }

    // add six points if the round is a win
    match round.opponent {
        Shape::Rock => {
            if round.player == Shape::Paper {
                points += 6
            }
        }
        Shape::Paper => {
            if round.player == Shape::Scissors {
                points += 6
            }
        }
        Shape::Scissors => {
            if round.player == Shape::Rock {
                points += 6
            }
        }
        _ => {}
    }

    points
}

#[cfg(test)]
mod test {
    use super::*;

    // test scores when player loses.
    #[test]
    fn test_score_round_loss_with_rock() {
        let round = Round {
            opponent: Shape::Paper,
            player: Shape::Rock,
        };
        assert_eq!(score_round(round), 1);
    }
    #[test]
    fn test_score_round_loss_with_paper() {
        let round = Round {
            opponent: Shape::Scissors,
            player: Shape::Paper,
        };
        assert_eq!(score_round(round), 2);
    }
    #[test]
    fn test_score_round_loss_with_scissors() {
        let round = Round {
            opponent: Shape::Rock,
            player: Shape::Scissors,
        };
        assert_eq!(score_round(round), 3);
    }

    // test scores when round is a draw
    #[test]
    fn test_score_round_draw_with_rock() {
        let round = Round {
            opponent: Shape::Rock,
            player: Shape::Rock,
        };
        assert_eq!(score_round(round), 4);
    }
    #[test]
    fn test_score_round_draw_with_paper() {
        let round = Round {
            opponent: Shape::Paper,
            player: Shape::Paper,
        };
        assert_eq!(score_round(round), 5);
    }
    #[test]
    fn test_score_round_draw_with_scissors() {
        let round = Round {
            opponent: Shape::Scissors,
            player: Shape::Scissors,
        };
        assert_eq!(score_round(round), 6);
    }

    // test scores when player wins
    #[test]
    fn test_score_round_win_with_rock() {
        let round = Round {
            opponent: Shape::Scissors,
            player: Shape::Rock,
        };
        assert_eq!(score_round(round), 7);
    }
    #[test]
    fn test_score_round_win_with_paper() {
        let round = Round {
            opponent: Shape::Rock,
            player: Shape::Paper,
        };
        assert_eq!(score_round(round), 8);
    }
    #[test]
    fn test_score_round_win_with_scissors() {
        let round = Round {
            opponent: Shape::Paper,
            player: Shape::Scissors,
        };
        assert_eq!(score_round(round), 9);
    }
}
