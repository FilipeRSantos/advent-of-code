use std::iter::Sum;
use std::str::FromStr;

struct CalibrationValue {
    first_score: u32,
    last_score: u32,
}

impl Sum<CalibrationValue> for u32
{
    fn sum<I: Iterator<Item=CalibrationValue>>(iter: I) -> Self {
        iter.fold(
            0,
            |mut total, item| {
                total += item.first_score * 10 + item.last_score;
                return total
            }
        )
    }
}


impl FromStr for CalibrationValue {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let reversed = s.chars().rev().collect::<String>();

        Ok(CalibrationValue {
            first_score: get_first_number(s).unwrap_or(0),
            last_score: get_first_number(&reversed).unwrap_or(0),
        })
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let values = input
        .lines()
        .map(|line| line.parse::<CalibrationValue>().expect("Call the engineers"))
        .collect::<Vec<_>>();

    let total: u32 = values.into_iter().sum();
    println!("{total}");

    Ok(())
}

fn get_first_number(line: &str) -> Option<u32> {
    line.chars().filter(|char| char.is_numeric()).next()?.to_digit(10)
}