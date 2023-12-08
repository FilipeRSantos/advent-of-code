use std::collections::HashMap;
use std::str::{FromStr};

#[derive(Debug, PartialEq)]
enum Instructions {
    Left,
    Right
}

#[derive(Debug, Eq, PartialEq, Hash, Copy, Clone)]
struct Coords {
    coord: u32,
}

impl FromStr for Coords {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Coords {
            coord: s.chars().fold(0, |mut state, char| {
                state *= 100;

                state += char as u32;

                state
            })
        })
    }
}

#[derive(Debug)]
struct Map {
    directions: Vec<Instructions>,
    coords: HashMap<Coords, (Coords, Coords)>,
}

impl FromStr for Map {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let mut coords = HashMap::new();

        let directions = lines.next().unwrap().chars().map(|m|
            if m == 'R' { Instructions::Right } else { Instructions::Left }
        ).collect::<Vec<_>>();

        lines.next();

        lines.for_each(|line| {
            let start = line[0..3].parse::<Coords>().unwrap();
            let left = line[7..10].parse::<Coords>().unwrap();
            let right = line[12..15].parse::<Coords>().unwrap();

            coords.insert(start, (left, right));
        });

        Ok(Map{
            directions,
            coords
        })
    }
}

impl Map {
    fn get_steps(&self) -> usize {
        let mut current_position = "AAA".parse::<Coords>().unwrap();
        let target_position = "ZZZ".parse::<Coords>().unwrap();
        let mut steps: usize = 0;

        while current_position != target_position {
            self.directions.iter().for_each(|i| {
                steps += 1;

                let coords = self.coords.get(&current_position).expect("Should not be empty");
                current_position = if *i == Instructions::Left { coords.0 } else { coords.1 };

                if current_position == target_position {
                    return;
                }
            });
        }

        steps
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");
    let map = input.parse::<Map>().expect("Should be a valid input");

    println!("{:?}", map.get_steps());

    Ok(())
}