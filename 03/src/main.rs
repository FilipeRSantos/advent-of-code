use std::str::FromStr;
use crate::Section::{Numeric, Symbol};

#[derive(Debug, PartialEq)]
enum Section {
    Numeric { value: u32, start_at: usize, end_at: usize },
    Symbol { position: usize },
}

impl Section {
    fn get_part_number(&self, symbols: &Vec<usize>) -> u32 {
        match *self {
            Numeric{ value, start_at, end_at } =>
                if symbols.iter().any(|s| *s >= if start_at == 0 { 0 } else {start_at - 1} && *s <= end_at) {
                    value
                } else {
                    0
                }
            ,
            _ => 0,
        }
    }
}

#[derive(Debug)]
struct Line {
    sections: Vec<Section>
}

impl Line {
    fn get_symbol_positions(&self) -> Vec<usize> {
        self.sections
            .iter()
            .filter_map(|e| match e {
                Symbol { position} => Some(*position),
                _ => None,
            })
            .collect()
    }
}

fn check_ongoing_number(start_position_current_number: Option<usize>, sections: &mut Vec::<Section>, index: usize, s: &str) {
    if start_position_current_number.is_some() {
        let temp = &s[start_position_current_number.unwrap()..index];
        let even_more_temp = temp.parse::<u32>().expect("");

        println!("{:?} - {:?}", temp, even_more_temp);

        sections.push(Numeric { value: even_more_temp, start_at: start_position_current_number.unwrap(), end_at: index })
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

                if c != '.' {
                    sections.push(Section::Symbol { position: index });
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
            let mut symbols = line.get_symbol_positions();

            if index > 0 {
                symbols.append(&mut lines.get(index - 1).unwrap().get_symbol_positions());
            }

            if index+1 < lines.len() {
                symbols.append(&mut lines.get(index + 1).unwrap().get_symbol_positions());
            }

            index += 1;

            line.sections
                .iter()
                .filter_map(|e| match e {
                    Numeric { .. } => Some(e.get_part_number(&symbols)),
                    _ => None,
                })
                .sum::<u32>()
        })
        .sum();

    println!("{:?}", sum);

    Ok(())
}