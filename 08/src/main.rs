use std::collections::HashMap;
use num::Integer;

#[derive(Debug, PartialEq)]
enum Instructions {
    Left,
    Right
}

#[derive(Debug)]
struct Map<'a> {
    directions: Vec<Instructions>,
    coords: HashMap<&'a str, (&'a str, &'a str)>,
}

impl<'a> Map<'a> {

    fn new(s: &str) -> Map {
        let mut lines = s.lines();
        let mut coords = HashMap::new();

        let directions = lines.next().unwrap().chars().map(|m|
            if m == 'R' { Instructions::Right } else { Instructions::Left }
        ).collect::<Vec<_>>();

        lines.next();

        lines.for_each(|line| {
            let start = &line[0..3];
            let left = &line[7..10];
            let right = &line[12..15];

            coords.insert(start, (left, right));
        });

        Map{
            directions,
            coords
        }
    }

    fn get_steps(&self) -> u128 {
        let mut steps: u128 = 0;
        let mut all_arrived = false;

        let mut current_positions = self.coords.iter().filter_map(|(coord, _)| {
            if coord.ends_with('A') { Some(coord) } else { None }
        }).collect::<Vec<_>>();

        let mut destinations = vec![];

        while current_positions.len() > 0 {
            self.directions.iter().for_each(|i| {
                all_arrived = true;
                steps += 1;

                current_positions = current_positions.iter().filter_map(|src| {
                    let coords = self.coords.get(*src).expect("Should not be empty");
                    let dst = if *i == Instructions::Left { &coords.0 } else { &coords.1 };

                    if dst.ends_with('Z') {
                        destinations.push(steps);
                        return None;
                    } else {
                        return Some(dst);
                    }
                }).collect();
            });
        }

        destinations.iter()
            .fold(*destinations.first().unwrap(), |mut acc, current| {
                acc = acc.lcm(current);
                acc
            })
    }


}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");
    let map = Map::new(input);

    println!("{:?}", map.get_steps());

    Ok(())
}