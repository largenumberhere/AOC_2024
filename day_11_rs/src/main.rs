use std::{include_str, usize};
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

    println!("checking {} {}", x, y);
    let mut line_count = 0;

    for dx in [-1, 1] {
        for dy in [-1, 1] {
            let has_line = try_check_direction((dx, dy), (x, y), grid);   
            if has_line {
                line_count += 1;
            }

            if line_count >= 2 {
                return  1;
            }
        }
    }

    return  0;
}

fn part2() {
    let grid = read_file()
        .unwrap();
    debug_grid(&grid);

    let mut tally = 0;
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == 'A' {
                let news = checkmas2(&grid, x, y);
                tally += news;
                if tally >= 1 {
                    println!("cross at {} {}", x, y);

                }
            }

        }
    }

    println!("total = {}", tally);


}
