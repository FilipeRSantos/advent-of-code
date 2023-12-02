use std::iter::Sum;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
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
        Ok(extract_values(s))
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

struct SpelledValue {
    description: &'static str,
    target: usize,
    last_correct_index: usize,
    current_streak: usize,
    value: u32,
}

impl SpelledValue {
    fn new (description: &'static str, value: u32) -> Self {
       SpelledValue {
           description,
           target: description.len(),
           last_correct_index: 0,
           current_streak: 0,
           value,
       }
    }

    fn check(&mut self, char: &char, index: usize) -> Option<u32> {
        if index != self.last_correct_index + 1 {
            self.current_streak = 0;
        }

        let keep_streak_going = self.description.chars().nth(self.current_streak).expect("should not overflow") == *char;

        //tthree would break here
        let start_again = !keep_streak_going && self.description.chars().nth(0).expect("should not be empty") == *char;

        if keep_streak_going || start_again {
            if start_again {
                self.current_streak = 0;
            }

            self.current_streak += 1;
            self.last_correct_index = index;

            if self.current_streak == self.target {
                self.current_streak = 0;
                return Some(self.value);
            }
        }

        None
    }
}

fn extract_values(line: &str) -> CalibrationValue {
    let mut first: Option<u32> = None;
    let mut last: Option<u32> = None;
    let mut index = 0;
    let mut options: [SpelledValue; 9] = [
        SpelledValue::new("one", 1),
        SpelledValue::new("two", 2),
        SpelledValue::new("three", 3),
        SpelledValue::new("four", 4),
        SpelledValue::new("five", 5),
        SpelledValue::new("six", 6),
        SpelledValue::new("seven", 7),
        SpelledValue::new("eight", 8),
        SpelledValue::new("nine", 9)];

    for c in line.chars() {
        let mut current: Option<u32> = None;

        if c.is_digit(10) {
            current = c.to_digit(10);
        } else {
            for x in &mut options {
                let placeholder = x.check(&c, index);

                //sevenine would break here
                if current.is_none() && placeholder.is_some() {
                    current = placeholder
                }
            }
        }

        if current.is_some() {
            last = current;

            if first.is_none() {
                first = current;
            }
        }

        index += 1;
    };

    CalibrationValue {
        first_score: first.expect("Cannot be empty"),
        last_score: last.expect("Cannot be empty"),
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn sample_tests() {
        let inputs: [&str; 7] = [
            "two1nine",
            "eightwothree",
            "abcone2threexyz",
            "xtwone3four",
            "4nineeightseven2",
            "zoneight234",
            "7pqrstsixteen",
        ];

        let outputs = inputs.into_iter().map(|x| extract_values(x)).collect::<Vec<_>>();
        assert_eq!(outputs.into_iter().sum::<u32>(), 281);
    }

    #[test]
    fn edge_cases() {
        assert_eq!(extract_values("one"),
                   CalibrationValue{ first_score: 1, last_score: 1});

        assert_eq!(extract_values("1"),
                   CalibrationValue{ first_score: 1, last_score: 1});

        assert_eq!(extract_values("9fourcsjph86shfqjrxlfourninev"),
                   CalibrationValue{ first_score: 9, last_score: 9});

        assert_eq!(extract_values("bjcrvvglvjn1"),
                   CalibrationValue{ first_score: 1, last_score: 1});

        assert_eq!(extract_values("321fivefour"),
                   CalibrationValue{ first_score: 3, last_score: 4});

        assert_eq!(extract_values("13five5"),
                   CalibrationValue{ first_score: 1, last_score: 5});

        assert_eq!(extract_values("four352"),
                   CalibrationValue{ first_score: 4, last_score: 2});

        assert_eq!(extract_values("7ninesevennine"),
                   CalibrationValue{ first_score: 7, last_score: 9});

        assert_eq!(extract_values("113six"),
                   CalibrationValue{ first_score: 1, last_score: 6});

        assert_eq!(extract_values("seven8onertbqhthreefourctdbsrcvcsjlvcxneight"),
                   CalibrationValue{ first_score: 7, last_score: 8});

        assert_eq!(extract_values("nine6eight"),
                   CalibrationValue{ first_score: 9, last_score: 8});

        assert_eq!(extract_values("fivesevenczmt22nfvnxhbvgjtvjmdzhfqhxtthree"),
                   CalibrationValue{ first_score: 5, last_score: 3});

        assert_eq!(extract_values("tthree"),
                   CalibrationValue{ first_score: 3, last_score: 3});

        assert_eq!(extract_values("9sevenklrvhclkrfourtwo96four"),
                   CalibrationValue{ first_score: 9, last_score: 4});

        assert_eq!(extract_values("onetwofivedblgtrxzpvmrhhsj2jhrcbbsseven3"),
                   CalibrationValue{ first_score: 1, last_score: 3});

        assert_eq!(extract_values("862twonec"),
                   CalibrationValue{ first_score: 8, last_score: 1});

        assert_eq!(extract_values("twone"),
                   CalibrationValue{ first_score: 2, last_score: 1});

        assert_eq!(extract_values("sixninesixoneeighttwo96"),
                   CalibrationValue{ first_score: 6, last_score: 6});

        assert_eq!(extract_values("threeseven1sixsix21five6"),
                   CalibrationValue{ first_score: 3, last_score: 6});

        assert_eq!(extract_values("peightwothree7pdmktbcfouroneninekrf6two"),
                   CalibrationValue{ first_score: 8, last_score: 2});

        assert_eq!(extract_values("eighthree"),
                   CalibrationValue{ first_score: 8, last_score: 3});

        assert_eq!(extract_values("sevenine"),
                   CalibrationValue{ first_score: 7, last_score: 9});
    }

}
