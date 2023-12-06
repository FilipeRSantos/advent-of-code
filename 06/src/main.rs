use std::str::{FromStr, Lines};

struct Race {
    duration: usize,
    current_record: usize,
}

impl Race {

    fn get_distance_with(&self, hold_time: usize) -> usize {
        (self.duration - hold_time) * hold_time
    }

    fn get_margin_of_error(&self) -> Vec<usize> {
        (0..=self.duration).filter(|n| self.get_distance_with(*n) > self.current_record).collect::<Vec<_>>()
    }
}

impl FromStr for Race {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {

        let mut iterator = s.lines();

        let duration = parse_data(&mut iterator);
        let current_record = parse_data(&mut iterator);

        Ok(Race {
            duration,
            current_record,
        })
    }
}

fn parse_data(lines: &mut Lines) -> usize {
    lines.next().expect("Should not be empty")[9..].split(':')
        .last()
        .expect("Should not be empty:")
        .trim()
        .replace(" ", "")
        .parse::<usize>().expect("Should be a valid number")
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");
    let race = input.parse::<Race>().expect("Should be a valid input");
    let ways_to_beat_record = race.get_margin_of_error().iter().count();

    println!("{:?}", ways_to_beat_record);

    Ok(())
}