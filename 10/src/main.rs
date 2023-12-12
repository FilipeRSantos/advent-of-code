use std::thread;

#[derive(Debug)]
enum CheckSide {
    Bottom,
    Top,
    Left,
    Right,
}

fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let width = input.lines().next().unwrap().len();
    let height = input.lines().count();
    let mut starting_width = usize::MAX;
    let mut starting_height = usize::MAX;

    let mut steps = vec![vec![-1; width]; height];
    let mut directions = vec![vec!['.'; width]; height];

    for (h, line) in input.lines().enumerate() {
        for (w, char) in line.chars().enumerate() {
            if char == 'S' {
                starting_width = w;
                starting_height = h;
            }
            directions[h][w] = char;
        }
    }

    let num: u64 = 100_000;
    thread::Builder::new().stack_size(num as usize * 0xFF).spawn(move || {
        calculate_position(starting_height, starting_width, &mut steps, &directions, (0, 0), 0);

        let mut max_value = 0;
        for (_, columns) in steps.iter().enumerate() {
            for (_, value) in columns.iter().enumerate() {
                if max_value < *value {
                    max_value = *value;
                }
            }
        }

        let mut loop_steps = vec![];
        for (row, columns) in steps.iter().enumerate() {
            for (column, value) in columns.iter().enumerate() {
                if *value == 0 {
                    loop_steps.push((row, column));
                }
            }
        }

        fill_loop(&steps, &mut directions, (starting_height, starting_width), (0,0), CheckSide::Right);

        let acc = directions.iter().fold(0, |mut acc, current| {
            acc += current.iter().filter(|char| **char == '=').count();
            acc
        });

        //TODO: See bug in point (99,71)(97,104), which was marked as outside

        for row in 0..height {
            for column in 0..width {
                let is_loop = steps[row][column] >= 0;
                let char = match directions[row][column] {
                        'F' => if is_loop { '┏' } else { ' ' },
                        'L' => if is_loop { '┗' } else { ' ' },
                        '7' => if is_loop { '┓' } else { ' ' },
                        'J' => if is_loop { '┛' } else { ' ' },
                        '-' => if is_loop { '━' } else { ' ' },
                        '|'|'S' => if is_loop { '┃' } else { ' ' },
                        '=' => '.',
                        _ => ' ',
                    };

                print!("{}", char)
            }
            println!()
        }

        println!("{}", acc);
    }).unwrap().join().unwrap();

    Ok(())
}

fn check_inside_position(steps: &Vec<Vec<i32>>, directions: &mut Vec<Vec<char>>, (y, x): (usize, usize)) {
    let is_loop = steps[y][x] >= 0;
    if is_loop {
        return;
    }

    if directions[y][x] == '=' {
        return;
    }

    directions[y][x] = '=';

    if y + 1 < directions.len() {
        check_inside_position(steps, directions, (y+1, x));
    }

    if x + 1 < directions.first().unwrap().len(){
        check_inside_position(steps, directions, (y, x+1));
    }

    if y > 0 {
        check_inside_position(steps, directions, (y-1, x));
    }

    if x > 0 {
        check_inside_position(steps, directions, (y, x-1));
    }
}

fn fill_loop(steps: &Vec<Vec<i32>>, directions: &mut Vec<Vec<char>>, (y, x): (usize, usize), (prev_y, prev_x): (usize, usize), check: CheckSide) {
    let current_symbol = directions[y][x];

    let from_top = prev_y < y;
    let from_left = prev_x < x;
    let from_bottom = prev_y > y;

    if current_symbol == 'S' && prev_x != 0 && prev_y != 0 {
        return;
    }

    let target =
        match current_symbol {
            'F' => if from_bottom { (y, x+1) } else { (y+1, x) },
            'L' => if from_top { (y, x+1) } else { (y-1, x) },
            'J' => if from_top { (y, x-1) } else { (y-1, x) },
            '7' => if from_bottom { (y, x-1) } else { (y+1, x) },
            '-' => if from_left { (y, x+1) } else { (y, x-1) },
            '|' => if from_bottom { (y-1, x) } else { (y+1, x) },
            'S' => (y+1, x),
            _ => panic!("{}", current_symbol),
        };

    match check {
        CheckSide::Bottom => {
            if y + 1 < directions.len() {
                check_inside_position(steps, directions, (y+1, x));
            }
        },
        CheckSide::Top => {
            if y > 0 {
                check_inside_position(steps, directions, (y-1, x));
            }
        }
        CheckSide::Left => {
            if x > 0 {
                check_inside_position(steps, directions, (y, x-1));
            }
        }
        CheckSide::Right => {
            if x + 1 < directions.first().unwrap().len() {
                check_inside_position(steps, directions, (y, x+1));
            }
        }
    }

    let check_in =
        match current_symbol {
            'F' | 'J' =>
                match check {
                    CheckSide::Left => CheckSide::Top,
                    CheckSide::Top => CheckSide::Left,
                    CheckSide::Bottom => CheckSide::Right,
                    CheckSide::Right => CheckSide::Bottom,
                },
            'L' | '7' =>
                match check {
                    CheckSide::Left => CheckSide::Bottom,
                    CheckSide::Bottom => CheckSide::Left,
                    CheckSide::Right => CheckSide::Top,
                    CheckSide::Top => CheckSide::Right,
                },
            '-' | '|' | 'S' => check,
            _ => panic!("{}", current_symbol),
        };

    fill_loop(steps, directions, target, (y, x), check_in);
}

fn calculate_position(h: usize, w: usize, steps: &mut Vec<Vec<i32>>, directions: &Vec<Vec<char>>, (prev_h, prev_w): (usize, usize), current_step: i32) {
    if h >= steps.len() || w >= steps.first().unwrap().len() {
        return;
    }

    let current_type = directions[h][w];

    let mut keep_looking =
        match current_type {
            '|' => prev_w == w && prev_h != h,
            '-' => prev_w != w && prev_h == h,
            'L' => (prev_h < h && prev_w == w) || (prev_h == h && prev_w > w),
            'J' => (prev_h < h && prev_w == w) || (prev_h == h && prev_w < w),
            'F' => (prev_h > h && prev_w == w) || (prev_h == h && prev_w > w),
            '7' => (prev_h > h && prev_w == w) || (prev_h == h && prev_w < w),
            'S' => true,
            '.' => false,
            _ => panic!("Should not be here"),
        };

    keep_looking = keep_looking && (steps[h][w] > current_step || steps[h][w] == -1);
    if !keep_looking {
        return;
    }

    steps[h][w] = current_step;

    let paths_to_check = match current_type {
        '|' => if h == 0 {vec![(h+1, w)]} else {vec![(h-1, w),(h+1, w)]},
        '-' => if w == 0 {vec![(h, w+1)]} else {vec![(h, w-1),(h, w+1)]},
        'L' => if h == 0 {vec![(h, w+1)]} else {vec![(h-1, w),(h, w+1)]},
        'J' => if h == 0 {vec![(h, w-1)]} else {vec![(h-1, w),(h, w-1)]},
        'F' => vec![(h+1, w),(h, w+1)],
        '7' => if w == 0 {vec![(h+1, w)]} else {vec![(h+1, w),(h, w-1)]},
        'S' => {
            let mut t = vec![(h+1, w), (h, w+1)];
            if w > 0 {
                t.push((h, w-1));
            }
            if h > 0 {
                t.push((h-1, w));
            }
            t
        }

        _ => panic!("Should not be here"),
    };

    for (target_h, target_w) in paths_to_check {
        if target_h == prev_h && target_w == prev_w {
            continue;
        }

        calculate_position(target_h, target_w, steps, &directions, (h, w), current_step + 1);
    }
}
