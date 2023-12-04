use std::collections::HashMap;
use std::str::FromStr;
#[derive(Debug)]
struct CardNumber {
    numbers: Vec<u8>
}

#[derive(Debug)]
struct Game {
    id: u32,
    winner: CardNumber,
    bet: CardNumber,
}

impl FromStr for CardNumber {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let numbers = s.trim().split_ascii_whitespace().map(|m| m.parse::<u8>().unwrap()).collect();

        Ok(CardNumber {
            numbers
        })
    }
}


impl FromStr for Game {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut sections = s.split(':');

        let id = *&sections.next()
            .expect("Should not be empty")[5..].trim().parse::<u32>()
            .expect("Should be a valid number");

        let mut numbers = sections.next().unwrap().split('|');
        let winning_numbers = numbers.next().unwrap().parse::<CardNumber>()?;
        let actual_numbers = numbers.next().unwrap().parse::<CardNumber>()?;

        Ok(Game {
            id,
            winner: winning_numbers,
            bet: actual_numbers,
        })
    }
}

impl Game {
    fn get_winning_numbers(&self) -> Vec<&u8> {
        self.bet.numbers.iter()
            .filter(|bet| self.winner.numbers.iter().any(|winner| winner.eq(bet)))
            .collect::<Vec<_>>()
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let mut multiplier: HashMap<u32, u32> = HashMap::new();

    let values = input
        .lines()
        .map(|line| line.parse::<Game>().expect("Call the engineers"))
        .filter_map(|card| {
            let right_picks = card.get_winning_numbers().iter().count() as u32;
            let cards = multiplier.get(&card.id).unwrap_or(&0) + 1;

            (1..=cards).for_each(|_| {
                (card.id+1..=card.id+right_picks).for_each(|n| {
                    *multiplier.entry(n).or_insert(0) += 1;
                });
            });

            Some(cards)
        })
        .collect::<Vec<_>>();

    let total: u32 = values.into_iter().sum();
    println!("{total}");

    Ok(())
}