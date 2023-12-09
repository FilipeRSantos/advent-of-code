fn check_current_line(values: &Vec<i32>) -> Option<Vec<i32>> {
    let mut previous_value = values.first().unwrap();
    let mut should_return = false;

    let new_values = values.iter().skip(1).map(|value| {
        if *value != 0 {
            should_return = true;
        }

        let diff = value - previous_value;

        previous_value = value;

        diff
    }).collect::<Vec<_>>();

    if should_return {
        Some(new_values)
    } else {
        None
    }

}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let extrapolated_values = input.lines().map(|line| {
        let values = line.split_ascii_whitespace()
                            .map(|value| value.parse::<i32>().expect("Should be a valid number"))
                            .collect::<Vec<_>>();

        let mut last_values = vec![];
        last_values.push(*values.last().unwrap());

        let mut current_values = values;
        while let Some(value) = check_current_line(&current_values) {
            last_values.push(*value.last().unwrap());
            current_values = value;
        }

        last_values.iter().rev().fold(0, |mut acc, current| {
            acc = acc + current;
            acc
        })
    }).sum::<i32>();

    println!("{:?}", extrapolated_values);

    Ok(())
}