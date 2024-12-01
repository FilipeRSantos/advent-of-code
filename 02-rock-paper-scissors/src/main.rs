#[repr(u8)]
#[derive(Debug, PartialEq, Clone)]
enum RockPaperScissorMoves {
    Rock = 1,
    Paper = 2,
    Scissor = 3,
}

#[derive(PartialEq)]
enum Strategy {
    Lose,
    Draw,
    Win,
}

#[derive(Debug)]
struct Game {
    my_move: RockPaperScissorMoves,
    opponent_move: RockPaperScissorMoves,
}

trait RockPaperScissorMatch {
    fn calculate_score(&self) -> u32;
    fn pick_move(&mut self, strategy: Strategy) -> u32;
}

impl RockPaperScissorMatch for Game {
    fn calculate_score(&self) -> u32 {
        let mut score = match self.my_move {
            RockPaperScissorMoves::Rock => 1,
            RockPaperScissorMoves::Paper => 2,
            RockPaperScissorMoves::Scissor => 3,
        };

        if self.my_move == self.opponent_move {
            score += 3;
        }

        if (self.my_move == RockPaperScissorMoves::Rock
            && self.opponent_move == RockPaperScissorMoves::Scissor)
            || (self.my_move == RockPaperScissorMoves::Scissor
                && self.opponent_move == RockPaperScissorMoves::Paper)
            || (self.my_move == RockPaperScissorMoves::Paper
                && self.opponent_move == RockPaperScissorMoves::Rock)
        {
            score += 6;
        }

        score
    }

    fn pick_move(&mut self, strategy: Strategy) -> u32 {
        if strategy == Strategy::Draw {
            self.my_move = self.opponent_move.clone();
            return self.calculate_score();
        }

        if (strategy == Strategy::Lose && self.opponent_move == RockPaperScissorMoves::Scissor)
            || (strategy == Strategy::Win && self.opponent_move == RockPaperScissorMoves::Rock)
        {
            self.my_move = RockPaperScissorMoves::Paper;
            return self.calculate_score();
        }

        if (strategy == Strategy::Lose && self.opponent_move == RockPaperScissorMoves::Rock)
            || (strategy == Strategy::Win && self.opponent_move == RockPaperScissorMoves::Paper)
        {
            self.my_move = RockPaperScissorMoves::Scissor;
            return self.calculate_score();
        }

        self.my_move = RockPaperScissorMoves::Rock;
        return self.calculate_score();
    }
}

fn get_moves_by_ascii(movement: u32) -> RockPaperScissorMoves {
    match movement {
        1 => RockPaperScissorMoves::Rock,
        2 => RockPaperScissorMoves::Paper,
        3 => RockPaperScissorMoves::Scissor,
        _ => panic!("Enum outside scope"),
    }
}

fn get_strategy_by_char(strategy: char) -> Strategy {
    match strategy {
        'X' => Strategy::Lose,
        'Y' => Strategy::Draw,
        'Z' => Strategy::Win,
        _ => panic!("Enum outside scope"),
    }
}

fn main() {
    let mut score_v1: u32 = 0;
    let mut score_v2: u32 = 0;

    for line in include_str!("input.txt").lines() {
        let mut moves = line.chars();
        let first_char = moves.next().unwrap();
        let last_char = moves.last().unwrap();
        let mut game = Game {
            opponent_move: get_moves_by_ascii(first_char as u32 - 64),
            my_move: get_moves_by_ascii(last_char as u32 - 87),
        };
        score_v1 += game.calculate_score();
        score_v2 += game.pick_move(get_strategy_by_char(last_char));
    }

    println!("The original score was: {}", score_v1);
    println!("The correct score is: {}", score_v2);
}
