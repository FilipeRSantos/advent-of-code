use std::io;

fn main() {
    let mut stacks = get_stacks().expect("Erro ao processar arquivo");

    stacks.sort_by(|a, b| b.cmp(a));
    let highest_stack = stacks.first().unwrap();
    let calories_amount = &stacks[..3];

    println!("The highest stack is: {}", highest_stack);
    println!(
        "The top 3 stacks combined are: {}",
        calories_amount.iter().sum::<u32>()
    );
}

fn get_stacks() -> Result<Vec<u32>, io::Error> {
    let lines = include_str!("input.txt");
    let mut stacks = vec![0];

    for current in lines.lines() {
        if current.is_empty() {
            stacks.push(0);
            continue;
        }

        let current_value: u32 = current.trim().parse().expect("Invalid calorie count");

        *stacks.last_mut().unwrap() += current_value;
    }

    Ok(stacks)
}
