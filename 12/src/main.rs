#[derive(Debug)]
struct Record<'a> {
    line: &'a str,
    groups: Vec<usize>,
}

impl<'a> Record<'a> {
    fn new(s: &'a str) -> Self {
        let mut values = s.split_ascii_whitespace();

        Record {
            line: values.next().unwrap(),
            groups: values.next().unwrap().split(',').map(|value| value.parse::<usize>().unwrap()).collect()
        }
    }

    fn get_valid_permutations_count(&self) -> usize {
        let unknowns = self.line.chars().filter(|char| *char == '?').count();
        let permutations: usize = 2_usize.pow(unknowns as u32);

        let mut correct_figures = 0;

        for i in 0..permutations {
            if self.check_variation(format!("{:0width$b}", i, width = unknowns)) {
                correct_figures += 1;
            }
        }

        correct_figures
    }

    fn check_variation(&self, pattern: String) -> bool {
        let mut unkowns = pattern.chars();
        let variation = self.line.chars().map(|char| {
            if char == '?' {
                match unkowns.next() {
                    None => panic!(""),
                    Some(char) =>
                        if char == '0' {
                            '.'
                        } else {
                            '#'
                        }
                }
            } else {
                char
            }
        }).collect::<String>();

        let sequence: Vec<usize> = variation
            .split('.')
            .filter_map(|current| {
                match current {
                    "." => None,
                    _ => Some(current)
                }
            })
            .map(|sequence| sequence.len())
            .filter(|sequence| *sequence > 0)
            .collect();

        sequence == self.groups
    }
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let values = input
        .lines()
        .map(|line| Record::new(line))
        .collect::<Vec<_>>();

    let acc = values.iter().fold(0, |mut acc, current| {
        acc += current.get_valid_permutations_count();

        acc
    });

    println!("{}", acc);

    Ok(())
}


#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn sample_tests() {
        let inputs: [&str; 6] = [
            "???.### 1,1,3",
            ".??..??...?##. 1,1,3",
            "?#?#?#?#?#?#?#? 1,3,1,6",
            "????.#...#... 4,1,1",
            "????.######..#####. 1,6,5",
            "?###???????? 3,2,1",
        ];

        let arrangements =
            inputs.iter()
                .map(|item| Record::new(item))
                .map(|record| record.get_valid_permutations_count())
                .sum::<usize>();

        assert_eq!(arrangements, 21);
    }
}
