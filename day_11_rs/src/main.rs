use std::{include_str, iter, usize};
use std::vec::Vec;

type Grid = Vec<Vec<char>>;

fn read_file() -> Option<Grid> {
    let mut grid = Vec::new();

    let infile = include_str!("../input.txt");
    for line in infile.lines(){
        if line.is_empty() {
            continue;
        }
        grid.push(Vec::new());
        for chr in line.chars() {
            grid.last_mut()?.push(chr);
        }
    }

    return Some(grid);
}


fn debug_grid(grid: &Grid) {
    for row in grid {
        for chr in row {
            print!("{}", chr);
        }
        println!("");
    }

}


fn checkmas(grid: &Grid, x: usize, y: usize) -> usize {
    let x = x as isize;
    let y= y as isize;
    let mut tally = 0;
    let xcount = grid[0].len() as isize;
    let ycount = grid.len() as isize;

    let mas = "XMAS";
    for dx in -1..=1 {
        for dy in -1..=1 {
            for peek in 0..=3 {
                if dx == 0 && dy == 0 {
                    continue;
                }

                let newx = x + (dx * peek);
                let newy = y + (dy * peek);

                if newy >= ycount  || newx >= xcount || newy < 0 || newx < 0 {
                    continue;
                }
                
                let value = grid[newy as usize][newx as usize];
                if value != (mas.as_bytes()[peek as usize] as char) {
                    break;  // try next coordinate
                }

                if peek ==3 {
                    tally += 1;
                }
            }
        }
    }


    return  tally;
}

#[allow(dead_code)]
fn part1() {
    let grid = read_file()
        .unwrap();
    debug_grid(&grid);

    let mut tally = 0;
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == 'X' {
                let news = checkmas(&grid, x, y);
                tally += news;
            }

        }
    }    

    println!("total = {}", tally);
}


fn main() {
    part2();    
}

fn try_check_direction(vector: (isize, isize), origin: (isize, isize), grid: &Grid) -> bool {
    let ycount = grid.len() as isize;
    let xcount = grid[0].len() as isize;

    let message = "MAS";

    let mut count = 0;
    for lookahead in -1..=1 {
        let newx = origin.0 + (vector.0 * lookahead);
        let newy = origin.1 + (vector.1 * lookahead);

        if newx < 0 || newy < 0 || newy >= ycount || newx >= xcount  {
            // skip out of bounds
            continue;
        } 

        let character = grid[newy as usize][newx as usize];
        let message_n = message.as_bytes()[(lookahead + 1) as usize] as char;
        if character == message_n {
            count += 1;
        }

    }

    if count >= 3 {
        return  true;
    }

    return false;
}


fn checkmas2(grid: &Grid, x: usize, y: usize) -> usize {
    let x = x as isize;
    let y = y as isize;

    let mut line_count = 0;
    
    let values = [-1, 1];
    let mut cur1 = -1 as isize;
    let mut cur2 = 0;
    let iter_diagonal_vectors = iter::from_fn(|| {
        cur1 += 1;
        
        if (cur1 as usize) >= values.len() {
            cur1 = 0;
            cur2 += 1;
        }

        if cur2 >= values.len() {
            return None;
        }

        let left = values[cur1 as usize];
        let right = values[cur2];
        return Some((left, right));

    });


    for (delta_x , delta_y) in iter_diagonal_vectors {
        let has_line = try_check_direction((delta_x, delta_y), (x, y), grid);   
        if has_line {
            line_count += 1;
        }

        if line_count >= 2 {
            return  1;
        }
    }

    return  0;
}

fn part2() {
    let grid = read_file()
        .unwrap();
    debug_grid(&grid);

    let mut iterx = -1 as isize;
    let mut itery = 0 as isize;
    let grid_iter = std::iter::from_fn(|| {
        let ycount = grid.len();
        let xcount = grid[0].len();

        iterx += 1;
        if (iterx as usize)  >= xcount {
            iterx = 0;
            itery += 1;
        }

        if (itery as usize) >= ycount {
            return  None;
        }

        return Some((iterx as usize, itery as usize));
    });

    let mut tally = 0;
    for (x, y) in grid_iter {
        if grid[y][x] == 'A' {
            let news = checkmas2(&grid, x, y);
            tally += news;
        }
    }


    println!("total = {}", tally);
}
