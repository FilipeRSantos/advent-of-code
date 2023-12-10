use std::thread;

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
        check_position(starting_height, starting_width, &mut steps, &directions, (0, 0), 0);
        let max = steps.iter().flat_map(|row| row.iter()).max();
        println!("{:?}", max);
    }).unwrap().join().unwrap();

    Ok(())
}

fn check_position(h: usize, w: usize, steps: &mut Vec<Vec<i32>>, directions: &Vec<Vec<char>>, (prev_h, prev_w): (usize, usize), current_step: i32) {
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

        check_position(target_h, target_w, steps, &directions, (h, w), current_step + 1);
    }
}
