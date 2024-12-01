use std::str::FromStr;

#[derive(Debug)]
struct Ruststack {
    first_compartment: String,
    second_compartment: String,
}

#[derive(Debug)]
struct Group {
    first: String,
    second: String,
    third: String,
}

trait GetPriority {
    fn get_priority(&self) -> u32;
}

impl GetPriority for char {
    fn get_priority(&self) -> u32 {
        if self.is_uppercase() {
            return *self as u32 - 38;
        }

        *self as u32 - 96
    }
}

impl GetPriority for Ruststack {
    fn get_priority(&self) -> u32 {
        let mut dupped_items = vec![];
        for item in self.first_compartment.chars() {
            if self.second_compartment.chars().any(|l| l == item) {
                if !dupped_items.contains(&item) {
                    dupped_items.push(item);
                }
            }
        }
        dupped_items.iter().map(|&item| item.get_priority()).sum()
    }
}

impl GetPriority for Group {
    fn get_priority(&self) -> u32 {
        let mut dupped_items = vec![];
        for item in self.first.chars() {
            if self.second.chars().any(|l| l == item) && self.third.chars().any(|l| l == item) {
                if !dupped_items.contains(&item) {
                    dupped_items.push(item);
                }
            }
        }
        dupped_items.iter().map(|&item| item.get_priority()).sum()
    }
}

impl FromStr for Ruststack {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let stack_size = s.len() / 2;
        Ok(Self {
            first_compartment: (&s[..stack_size]).parse().unwrap(),
            second_compartment: (&s[stack_size..]).parse().unwrap(),
        })
    }
}

fn main() {
    let lines: u32 = include_str!("input.txt")
        .lines()
        .map(|line| line.parse::<Ruststack>())
        .map(|ruststack| ruststack.unwrap().get_priority())
        .sum();
    dbg!(lines);

    let mut groups = include_str!("input.txt").lines();
    let mut group_priority = 0;
    while let (Some(a), Some(b), Some(c)) = (groups.next(), groups.next(), groups.next()) {
        let group = Group {
            first: a.parse().unwrap(),
            second: b.parse().unwrap(),
            third: c.parse().unwrap(),
        };
        group_priority += group.get_priority();
    }
    dbg!(group_priority);
}
