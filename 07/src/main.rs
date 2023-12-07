use std::collections::HashMap;
use std::str::{FromStr};

#[derive(Debug)]
enum HandType {
    FiveOfAKind = 7,
    FourOfAKind = 6,
    FullHouse = 5,
    ThreeOfAKind = 4,
    TwoPair = 3,
    OnePair = 2,
    HighCard = 1,
}

#[derive(Debug)]
struct PokerHand {
    cards: Vec<usize>,
    bid: usize,
}

#[derive(Debug)]
struct WeightedHand {
    weight: usize,
    bid: usize,
}

impl PokerHand {
    fn get_figure(&self) -> HandType {
        let mut values = HashMap::new();
        self.cards.iter().for_each(|card| {
           *values.entry(*card as usize).or_insert(0) += 1;
        });

        match values.len() {
            1 => HandType::FiveOfAKind,
            2 => if values.iter().any(|(_, value)| *value == 3) { HandType::FullHouse } else { HandType::FourOfAKind },
            3 =>
                if values.iter().any(|(_, value)| *value == 3) {
                    if values.iter().any(|(_, value)| *value == 1) { HandType::ThreeOfAKind } else { HandType::FullHouse }
                } else { HandType::TwoPair },
            4 => if values.iter().any(|(_, value)| *value == 2) { HandType::OnePair } else { HandType::HighCard },
            _ => HandType::HighCard
        }
    }

    fn get_hand_weight(&self) -> usize {
        let mut multiplier = 100000000;
        let multiplier_type = self.get_figure() as usize;

        self.cards.iter().fold(multiplier_type * multiplier * 100, |mut state, x| {
            state += x * multiplier;

            multiplier /= 100;

            state
        })
    }
}

impl FromStr for PokerHand {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(PokerHand {
            cards: s[0..5].chars().map(|char|
                match char {
                    'A' => 14,
                    'K' => 13,
                    'Q' => 12,
                    'J' => 11,
                    'T' => 10,
                    '9' => 9,
                    '8' => 8,
                    '7' => 7,
                    '6' => 6,
                    '5' => 5,
                    '4' => 4,
                    '3' => 3,
                    '2' => 2,
                    _ => panic!("{} should not exist", char)
                }
            ).collect::<Vec<_>>(),
            bid: s[6..].parse::<usize>().expect("Should contain valid bid"),
        })
    }
}

#[derive(Debug)]
struct Round {
    hands: Vec<PokerHand>,
}

impl Round {
    fn get_total_earning(&self) -> usize {

        let mut weights = self.hands.iter()
            .map(|hand| {
                WeightedHand {
                    bid: hand.bid,
                    weight: hand.get_hand_weight(),
                }
            })
            .collect::<Vec<_>>();

        weights.sort_by_key(|item| item.weight);

        let mut winnings = 0;
        for (pos, e) in weights.iter().enumerate() {
            winnings += (pos + 1) * e.bid;
        }

        winnings
    }
}

impl FromStr for Round {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Round {
            hands: s.lines().map(|line| line.parse::<PokerHand>().expect("Should be valid camel hand")).collect::<Vec<_>>()
        })
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");
    let round = input.parse::<Round>().expect("Should be a valid input");

    println!("{:?}", round.get_total_earning());

    Ok(())
}