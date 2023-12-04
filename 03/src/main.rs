use std::str::FromStr;
use crate::Section::{Numeric, Gear};

#[derive(Debug, PartialEq)]
enum Section {
    Numeric { value: u32, start_at: usize, end_at: usize },
    Gear { position: usize },
}

impl Section {
    fn get_gear_ratio(&self, parts: Vec<&Section>) -> u32 {
        match *self {
            Gear{ position } => {
                let mut close_parts = parts
                    .iter()
                    .filter_map(|e| match e {
                        Numeric { value, start_at, end_at } => {
                            if position >= if *start_at == 0 { 0 } else {start_at - 1} && position <= *end_at { Some(*value) } else { None }
                        },
                        _ => None,
                    });

                if close_parts.clone().count() == 2 {
                    close_parts.next().unwrap() * close_parts.next().unwrap()
                } else { 0 }
            },
            _ => 0,
        }
    }
}

#[derive(Debug)]
struct Line {
    sections: Vec<Section>
}

impl Line {
    fn get_parts(&self) -> Vec<&Section> {
        self.sections
            .iter()
            .filter_map(|e| match e {
                Numeric { .. } => Some(e),
                _ => None,
            })
            .collect()
    }
}

fn check_ongoing_number(start_position_current_number: Option<usize>, sections: &mut Vec::<Section>, index: usize, s: &str) {
    if start_position_current_number.is_some() {
        sections.push(Numeric { value: s[start_position_current_number.unwrap()..index].parse::<u32>().expect(""), start_at: start_position_current_number.unwrap(), end_at: index })
    }
}

impl FromStr for Line {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {

        let mut sections: Vec::<Section> = vec!();
        let mut start_position_current_number: Option<usize> = None;
        let mut index: usize = 0;

        s.chars().for_each(|c| {

            if c.is_numeric()  {
                if start_position_current_number.is_none() {
                    start_position_current_number = Some(index);
                }
            } else {
                check_ongoing_number(start_position_current_number, &mut sections, index, s);

                if c == '*' {
                    sections.push(Section::Gear { position: index });
                }

                start_position_current_number = None;
            }

            index += 1;
        });

        check_ongoing_number(start_position_current_number, &mut sections, index, s);

        Ok(Line{
            sections
        })
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let lines = input
        .lines()
        .map(|line| line.parse::<Line>().expect("Call the engineers"))
        .collect::<Vec<_>>();

    let mut index = 0;
    let sum: u32 = lines.iter()
        .map(|line| {
            let mut parts = line.get_parts();

            if index > 0 {
                parts.append(&mut lines.get(index - 1).unwrap().get_parts());
            }

            if index+1 < lines.len() {
                parts.append(&mut lines.get(index + 1).unwrap().get_parts());
            }

            index += 1;

            line.sections
                .iter()
                .filter_map(|e| match e {
                    Gear { .. } => Some(e.get_gear_ratio(parts.clone())),
                    _ => None,
                })
                .sum::<u32>()
        })
        .sum();

    println!("{:?}", sum);

    Ok(())
}