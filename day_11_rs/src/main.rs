use std::{include_str, usize};
use std::vec::Vec;

type Grid = Vec<Vec<char>>;

#[derive(Clone, Copy)]
struct Pos {
    x: isize,
    y: isize,
}

impl Pos {
    fn new_usize(x: usize, y: usize) -> Pos {
        Pos { x: x as _, y: y as _ }
    }

    fn new_isize(x: isize, y: isize) -> Pos {
        Pos { x, y }
    }

    fn x_usize(self: Self) -> usize {
        self.x as _
    }

    fn y_usize(self: Self) -> usize {
        self.y as _
    }

}

fn read_file() -> Option<Grid> {
    let mut grid = Vec::new();

    let infile = include_str!("../sample_input.txt");
    for line in infile.lines(){
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


#[derive(PartialEq)]
enum JollyFactor {
    Possibly,
    NotJolly,
    MaximallyJolly
}

struct JollyMachine {
    jolly_cursor: usize,
    phrase: Vec<char>,
    grid: Grid,
}


impl JollyMachine {
    fn new(grid: Grid) -> JollyMachine {
        JollyMachine { 
            jolly_cursor: 0, 
            phrase: "XMAX".chars().collect(),
            grid: grid
        }
    }

    fn start_section(self: &mut JollyMachine) -> () {
        self.jolly_cursor = 0;
    }

    fn end_section(self: &mut Self) -> JollyFactor {
        if self.jolly_cursor >= self.phrase.len()-1 {
            return JollyFactor::MaximallyJolly;
        } else {
            return  JollyFactor::NotJolly;
        }
    }
    
    fn is_jolly(self: &mut Self, pos: Pos) -> JollyFactor {
        if self.jolly_cursor >= self.phrase.len()-1 {
            self.jolly_cursor = 0;
            return JollyFactor::MaximallyJolly;
        }

        let current = self.phrase[self.jolly_cursor];
        let item = self.grid[pos.y_usize()][pos.x_usize()];

        if item != current {
            self.jolly_cursor = 0;
            return JollyFactor::NotJolly;
        }
        
        self.jolly_cursor += 1;
        return  JollyFactor::Possibly;
        
    }
}




fn main() {
    let grid = read_file()
        .unwrap();
    debug_grid(&grid);

    let minx  = 0 as isize;
    let miny = 0 as isize;
    let county = grid.len() as isize;
    let countx = grid[0].len() as isize;
    let lastx = countx - 1 as isize;
    let lasty = county - 1 as isize;

    let mut xmax_tally = 0;
    
    /*  Compass
        NW N NE
        W  x E
        SW S SE

        The cardinal directions are actually inverted from the memory layout. North is an increase in y, east is an increase in x, etc.
     */

    // N
    {
        let mut posp = Pos::new_isize(minx, miny);
        while posp.y < county {
            let mut machine = JollyMachine::new(grid.clone());
            machine.start_section();
            let mut pos2 = posp.clone();
            while pos2.x < countx  {
                let jolly = machine.is_jolly(pos2);
                match jolly {
                    JollyFactor::Possibly => {},
                    JollyFactor::NotJolly => {
                        machine.end_section();
                    },
                    JollyFactor::MaximallyJolly => {
                        xmax_tally += 1;
                        machine.end_section();
                        println!("jolly at {} {}", pos2.x, pos2.y);
                    },
                };

                println!("{} {}", pos2.x, pos2.y);
                pos2.x += 1;
            }

            let jolly = machine.end_section();
            if jolly == JollyFactor::MaximallyJolly {
                xmax_tally += 1;
                println!("jolly at {} {}", pos2.x, pos2.y);
            }


            posp.y += 1;
        }

        println!("tally after section 1 {}", xmax_tally);
    }
    {
        // S
        let mut posp = Pos::new_isize(minx, miny);
        while posp.y < county {
            let mut machine = JollyMachine::new(grid.clone());
            machine.start_section();

            let mut pos2 = Pos::new_isize(lastx, posp.y);
            while pos2.x >= 0 {
                let jolly = machine.is_jolly(pos2);
                match jolly {
                    JollyFactor::Possibly => {},
                    JollyFactor::NotJolly => {
                        machine.end_section();
                    }
                    JollyFactor::MaximallyJolly => {
                        xmax_tally += 1;
                        machine.end_section();
                        println!("jolly at {} {}", pos2.x, pos2.y);
                    }
                };

                println!("{} {}", pos2.x, pos2.y);


                pos2.x -= 1;
            }

            let jolly = machine.end_section();
            if jolly == JollyFactor::MaximallyJolly {
                xmax_tally += 1;
                println!("jolly at {} {}", pos2.x, pos2.y);
            }

            posp.y += 1;
        }
    
        println!("tally after section 2 {}", xmax_tally);
    }

    


    // W
    // E

    // NE
    // SE
    // SW
    // NW

    
}
