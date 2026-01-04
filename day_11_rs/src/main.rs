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


fn main() {
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
