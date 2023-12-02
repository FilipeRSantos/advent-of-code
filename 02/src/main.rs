use std::iter::Sum;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
enum Color {
    Red,
    Blue,
    Green
}

#[derive(Debug)]
struct Set {
    color: Color,
    count: u32,
}

#[derive(Debug)]
struct Round {
    draws: Vec<Set>,
}

#[derive(Debug)]
struct Game {
    id: u32,
    rounds: Vec<Round>,
}

impl FromStr for Color {
    type Err = ();

    fn from_str(input: &str) -> Result<Color, Self::Err> {
        match input {
            "blue"  => Ok(Color::Blue),
            "green"  => Ok(Color::Green),
            "red"  => Ok(Color::Red),
            _      => Err(()),
        }
    }
}

impl FromStr for Set {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut values = s.trim().split_ascii_whitespace();

        let count = *&values.next()
            .expect("Should not be empty").parse::<u32>()
            .expect("Should be a valid number");

        Ok(Set{
            count,
            color: (values.next().expect("Should contain one whitespace").parse::<Color>().expect("Should follow doc expecifications"))
        })
    }
}

impl FromStr for Round {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let draws = s.split(',')
            .map(|draw| draw.parse::<Set>().expect("Failed to parse draws"))
            .collect::<Vec<_>>();

        Ok(Round {
            draws
        })
    }
}

impl Game {
    fn get_max_draw(&self, color: Color) -> u32 {
        self.rounds.iter()
            .flat_map(|round| round.draws.iter())
            .filter(|set| set.color == color)
            .max_by_key(|set| set.count)
            .unwrap_or(&Set { count: 0, color })
            .count
    }
}

impl FromStr for Game {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut sections = s.split(':');

        let id = *&sections.next()
                        .expect("Should not be empty")[5..].parse::<u32>()
                        .expect("Should be a valid number");

        let rounds = sections.next()
                                    .expect("Should contain :")
                                    .split(';')
                                    .map(|section| section.parse::<Round>().expect("Failed to parse rounds"))
                                    .collect::<Vec<_>>();

        Ok(Game {
            id,
            rounds,
        })
    }
}

impl Sum<Game> for u32 {
    fn sum<I: Iterator<Item=Game>>(iter: I) -> Self {
        iter.fold(
            0,
            |mut total, game| {
                total += game.id;
                return total
            }
        )
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let values = input
        .lines()
        .map(|line| line.parse::<Game>().expect("Call the engineers"))
        .map(|game| {
            let blue = game.get_max_draw(Color::Blue);
            let green = game.get_max_draw(Color::Green);
            let red = game.get_max_draw(Color::Red);

            return blue * green * red
        })
        .collect::<Vec<_>>();

    let total: u32 = values.into_iter().sum();
    println!("{total}");

    Ok(())
}