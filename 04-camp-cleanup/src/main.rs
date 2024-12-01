use std::str::FromStr;

#[derive(Debug, Clone, Copy)]
struct Section {
    first: usize,
    last: usize,
}

#[derive(Debug, Copy, Clone)]
struct Group {
    first: Section,
    last: Section,
}

trait Contains {
    fn overlaps(&self, section: &Section) -> bool;
    fn contains(&self, section: &Section) -> bool;
}

impl Contains for Section {
    fn contains(&self, section: &Section) -> bool {
        self.first <= section.first && self.last >= section.last
    }

    fn overlaps(&self, section: &Section) -> bool {
        self.first <= section.last && self.last >= section.first
    }
}

impl FromStr for Section {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut splits = s.split('-');
        Ok(Section {
            first: splits.next().unwrap().parse().unwrap(),
            last: splits.last().unwrap().parse().unwrap(),
        })
    }
}

impl FromStr for Group {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut splits = s.split(',');
        Ok(Group {
            first: splits.next().unwrap().parse().unwrap(),
            last: splits.last().unwrap().parse().unwrap(),
        })
    }
}

fn main() {
    let wrong_groups = include_str!("input.txt")
        .lines()
        .map(|line| line.parse::<Group>())
        .filter(|&group| {
            let current = group.unwrap();
            // current.first.contains(&current.last) || current.last.contains(&current.first)
            current.first.overlaps(&current.last) || current.last.overlaps(&current.first)
        })
        .count();

    dbg!(wrong_groups);
}
