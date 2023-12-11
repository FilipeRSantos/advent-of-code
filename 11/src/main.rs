use std::cmp::{max, min};

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");
    let mut galaxies = vec![];

    for (y, line) in input.lines().enumerate() {
        for (w, char) in line.chars().enumerate() {
            if char == '#' {
                galaxies.push((y, w));
            }
        }
    }

    let empty_rows =
        (0..input.lines().count())
            .filter(|row| !galaxies.iter().any(|galaxy| galaxy.0 == *row))
            .collect::<Vec<_>>();

    let empty_columns =
        (0..input.lines().next().unwrap().len())
            .filter(|column| !galaxies.iter().any(|galaxy| galaxy.1 == *column))
            .collect::<Vec<_>>();

    let mut acc = 0;
    const MULTIPLIER: usize = 1000000 - 1;
    for (index, current) in galaxies.iter().enumerate() {
        let other_galaxies = galaxies.iter().skip(index+1);
        for (_, target) in other_galaxies.into_iter().enumerate() {
            let start_row = min(current.0, target.0);
            let start_column = min(current.1, target.1);
            let end_row = max(current.0, target.0);
            let end_column = max(current.1, target.1);

            let offset_columns = empty_columns.iter()
                .filter(|column| start_column < **column && **column < end_column)
                .count() * MULTIPLIER;

            let offset_rows = empty_rows.iter()
                .filter(|row| start_row < **row && **row < end_row)
                .count() * MULTIPLIER;

            let distance = (end_row - start_row) + (end_column - start_column) + offset_columns + offset_rows;
            acc += distance;
        }
    }

    println!("{:?}", acc);

    Ok(())
}
