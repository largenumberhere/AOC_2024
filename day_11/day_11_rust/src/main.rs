use hashbag::HashBag;

fn digits(mut stone: i64) -> usize {
    let mut digits = 0;
    while  stone > 0 {
        digits += 1;
        stone /= 10;
    }

    return digits
}

fn stone_left(mut stone: i64) -> i64 {
    let stone_length = digits(stone);
    for _ in 0..stone_length/2 {
        stone /=10;
    }

    let mut to = 0;
    let mut mul = 0;

    for _ in 0..stone_length/2 {
        if stone == 0  {
            break
        }

        let digit = stone % 10;
        stone /= 10;
        to = to + (digit * i64::pow(10, mul));
        mul +=1;
    }

    return to;
}

fn stone_right(mut stone: i64) -> i64 {
    let mut out = 0;
    let mut power = 0;

    for _ in 0.. digits(stone)/2 {
        let digit = stone % 10;
        stone /= 10;

        out = out + (digit * i64::pow(10, power));
        power+=1;
    }

    return out;
}

fn evaluate_stone(stone: i64) -> (i64, Option<i64>) {
    if stone == 0 {
        return (1, None);

    } else if digits(stone) % 2 == 0 {
        let left = stone_left(stone);
        let right = stone_right(stone);

        return (left, Some(right))
    }

    return (stone * 2024, None)
}


fn iterate_stone(stones: &mut HashBag<i64>, tmp: &mut HashBag<i64>) {
    tmp.clone_from(stones);

    for (stone, count) in tmp.set_iter() {
        let stone = *stone;

        let result = evaluate_stone(stone);

        // replace stone if needed
        if result.0 != stone {
            let item = stones.get(&stone)
                .expect("trying to remove from 0 items");
            if item.1 < count {
                panic!("trying to remove too many items")
            }

            stones.remove_up_to(&stone, count);
            stones.insert_many(result.0, count);
        }

        if let Some(v) = result.1 {
            stones.insert_many(v, count);
        }
    }

}

fn count_stones(bag: &HashBag<i64>) -> u64 {
    let mut tally = 0;
    for (_, v) in bag.set_iter() {
        tally += v;
    }

    return tally as u64;
}

#[allow(unused)]    // this is useful in debugging scenarios
fn print_stones(bag: &HashBag<i64>) {
    println!("Bag {{");
    for (k,v) in bag.set_iter() {
        println!("    {k}: {v}");
    }
    println!("}}");
}

fn main() {
    let mut bag: HashBag<i64> = HashBag::new();

    let stones: [i64; 2] = [125, 17];


    for stone in stones.into_iter() {
        bag.insert(stone);
    }

    let mut tmp = HashBag::new();
    let iterations = 75;
    for _ in 0 .. iterations {
        iterate_stone(&mut bag, &mut tmp);
    }

    println!("iteration {} has count {}", iterations, count_stones(&bag));
}

