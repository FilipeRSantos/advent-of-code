use std::collections::HashMap;
use std::str::{FromStr};
use crate::HandType::{FiveOfAKind, FourOfAKind, FullHouse, HighCard, OnePair, ThreeOfAKind, TwoPair};

#[derive(Debug, PartialEq)]
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
        let joker_amount = self.cards.iter().filter(|card| **card == 1).count();

        self.cards.iter().filter(|card| **card != 1).for_each(|card| {
           *values.entry(*card as usize).or_insert(0) += 1;
        });

        let mut current_figure = HandType::HighCard;
        values.iter().for_each(|(_, size)| {
            if *size != 1 {
                current_figure = match size {
                    5 => FiveOfAKind,
                    4 => FourOfAKind,
                    3 => if current_figure == OnePair { FullHouse } else { ThreeOfAKind },
                    2 => if current_figure == ThreeOfAKind { FullHouse }
                            else if current_figure == OnePair { TwoPair }
                            else { OnePair },
                    _ => panic!("Should not be here")
                };
            }

        });

        for _ in 0..joker_amount {
            current_figure = match current_figure {
                HighCard => OnePair,
                OnePair => ThreeOfAKind,
                TwoPair => FullHouse,
                ThreeOfAKind => FourOfAKind,
                FullHouse => FourOfAKind,
                FourOfAKind => FiveOfAKind,
                FiveOfAKind => FiveOfAKind,
            };
        }

        current_figure
    }

    fn get_hand_weight(&self) -> usize {
        let mut multiplier = 100000000;
        let current_figure = self.get_figure();

        self.cards.iter().fold(current_figure as usize * multiplier * 100, |mut state, x| {
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
                    'T' => 10,
                    '2'..='9' => char.to_digit(10).expect("Should be a valid number") as usize,
                    'J' => 1,
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