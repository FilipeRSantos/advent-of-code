use std::str::{FromStr, Lines};

struct Race {
    duration: usize,
    current_record: usize,
}

struct Races {
    records: Vec<Race>,
}

impl Race {

    fn get_distance_with(&self, hold_time: usize) -> usize {
        (self.duration - hold_time) * hold_time
    }

    fn get_margin_of_error(&self) -> Vec<usize> {
        (0..=self.duration).filter(|n| self.get_distance_with(*n) > self.current_record).collect::<Vec<_>>()
    }
}

impl FromStr for Races {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {

        let mut iterator = s.lines();

        let times = parse_data(&mut iterator).into_iter();
        let mut distances = parse_data(&mut iterator).into_iter();

        Ok(Races {
            records: times.map(|time| {
                                Race {
                                    duration: time,
                                    current_record: distances.next().expect("Should have same number of elements"),
                                }
                            })
                            .collect::<Vec<_>>(),
        })
    }
}

fn parse_data(lines: &mut Lines) -> Vec<usize> {
    lines.next().expect("Should not be empty")[9..].split(':')
        .last()
        .expect("Should not be empty:")
        .trim()
        .split_ascii_whitespace()
        .map(|value| value.parse().expect("Should be a valid number")).collect::<Vec<_>>()
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let races = input.parse::<Races>().expect("Should be a valid input");
    let mut product: usize = 1;

    races.records.iter().for_each(|race| {
        let ways_to_beat_record = race.get_margin_of_error().iter().count();

        if ways_to_beat_record > 0 {
            product *= ways_to_beat_record;
        }
    });

    println!("{:?}", product);

    Ok(())
}