package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib"
)

// func removeLeadingZeros(integer *string) {
// 	zeros := 0
// 	for i := 0; i < len(*integer); i++ {
// 		if (*integer)[i] != '0' {
// 			break
// 		} else {
// 			zeros += 1
// 		}
// 	}

// 	if zeros > 0 {
// 		*integer = (*integer)[zeros:]
// 	}

// 	if *integer == "" {
// 		*integer = "0"
// 	}
// }

func stoneLen(stone int) int {
	size := 0
	for stone != 0 {
		size++
		stone /= 10
	}

	return size
}

func intPow(base int, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))

	// if exponent < 1 {
	// 	panic("negative exponent not supported")
	// }

	// if exponent == 0 {
	// 	return 1
	// }

	// out := base
	// for i := 0; i < exponent; i++ {
	// 	out += base
	// }

	// return out
}

/*
	static STONE_T left_half(STONE_T value, int length) {
	   STONE_T vin = value;
	    // discard first half
	    for (int i = 0; i < (length/2); i++) {
	        value /=10;
	    }


	    STONE_T to = 0;
	    STONE_T mul = 0;
	    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
	        // pop off last digit
	        STONE_T digit = value % 10;
	        value /= 10;
	        to = to + (digit * (size_t)pow(10, mul));
	        mul ++;

	    }

	    return to;
	}
*/
func stoneLeft(stone int, stone_length int) int {
	input := stone
	for i := 0; i < stone_length/2; i++ {
		input /= 10
	}

	to := 0
	mul := 0
	for i := 0; (i < stone_length/2) && (input != 0); i++ {
		digit := input % 10
		input /= 10
		to = to + (digit * intPow(10, mul))
		mul += 1
	}

	return to
}

/*
	static STONE_T right_half(STONE_T value, int length) {
	    STONE_T to = 0;
	    STONE_T mul = 0;
	    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
	        // pop off last digit
	        STONE_T digit = value % 10;
	        value /= 10;
	        to = to + (digit * pow(10, mul));
	        mul ++;

	    }

	    return to;
	}
*/
func stoneRight(stone int, stone_length int) int {
	input := stone

	to := 0
	mul := 0
	for i := 0; (i < stone_length/2) && (input != 0); i++ {
		digit := input % 10
		input /= 10
		to = to + (digit * intPow(10, mul))
		mul += 1
	}

	return to
}

func updateStone(stone int, out1 *int, out2 *int) bool {
	if stone == 0 {
		*out1 = 1
		return false
	} else if stoneLen(stone)%2 == 0 {
		// midpoint := stoneLen(*stone) / 2

		stone_length := stoneLen(stone)
		left := stoneLeft(stone, stone_length)
		right := stoneRight(stone, stone_length)

		*out1 = left
		*out2 = right

		return true
	} else {
		stone_val := (stone) * 2024
		*out1 = stone_val

		return false
	}
}

func updateStoneConcurrent(stone int) UpdateResult {
	var out1 int
	var out2 int
	used_out2 := updateStone(stone, &out1, &out2)

	if used_out2 {
		return UpdateResult2(stone, out1, out2)
	} else {
		return UpdateResult1(stone, out1)
	}
}

func UpdateResult1(stone_in int, stone1 int) UpdateResult {
	return UpdateResult{
		is_valid:   true,
		stone_in:   stone_in,
		stone1:     stone1,
		has_stone2: false,
	}
}

func UpdateResult2(stone_in int, stone1 int, stone2 int) UpdateResult {
	return UpdateResult{
		is_valid:   true,
		stone_in:   stone_in,
		stone1:     stone1,
		new_stone2: stone2,
		has_stone2: true,
	}
}

type UpdateResult struct {
	is_valid   bool
	stone_in   int
	stone1     int
	new_stone2 int
	id         int
	has_stone2 bool
	count      int
}

// func processStones(stones *[]int, startIndex, nStonesToProcess int, results *[]UpdateResult, wg *sync.WaitGroup) {
// 	for i := 0; i < nStonesToProcess; i++ {
// 		index := startIndex + i

// 		var out1 int
// 		var out2 int
// 		used_out2 := updateStone((*stones)[index], &out1, &out2)
// 		if used_out2 {
// 			(*results)[index] = UpdateResult2(out1, out2, index)
// 		} else {
// 			(*results)[index] = UpdateResult1(out1, index)
// 		}
// 	}

// 	wg.Done()
// }

var numCores = runtime.NumCPU()

func updateStones(stones *map[int]int) {

	if countStones(*stones) <= numCores {
		waits := sync.WaitGroup{}
		output := make(chan UpdateResult)
		for key, keys_count := range *stones {
			for i := 0; i < keys_count; i++ {

				waits.Add(1)
				go func(output chan UpdateResult, key int, wg *sync.WaitGroup) {
					defer wg.Done()
					output <- updateStoneConcurrent(key)

				}(output, key, &waits)

			}
		}

		go func() {
			waits.Wait()
			close(output)
		}()

		for result := range output {
			// fmt.Println(result)
			// if stone has been replaced, remove input stone
			if result.stone1 != result.stone_in {
				(*stones)[result.stone_in] -= 1
			}

			// initialize entry in map
			_, key_exists := (*stones)[result.stone1]
			if !key_exists {
				(*stones)[result.stone1] = 0
			}

			(*stones)[result.stone1] += 1

			// add new stone if any
			if result.has_stone2 {
				_, key_exists := (*stones)[result.new_stone2]
				if !key_exists {
					(*stones)[result.new_stone2] = 0
				}

				(*stones)[result.new_stone2] += 1
			}
		}
	} else {
		// otherwise do one job for each key in the map
		channel := make(chan UpdateResult)
		waits := sync.WaitGroup{}
		for key, count := range *stones {
			waits.Add(1)
			go func(stone int, count int) {
				result := updateStoneConcurrent(stone)
				result.count = count
				channel <- result
				waits.Done()
			}(key, count)
		}

		go func() {
			waits.Wait()
			close(channel)
		}()

		// add the results
		for result := range channel {
			for i := 0; i < result.count; i++ {
				if result.stone1 != result.stone_in {
					(*stones)[result.stone_in] -= 1
				}

				// initialize entry in map
				_, key_exists := (*stones)[result.stone1]
				if !key_exists {
					(*stones)[result.stone1] = 0
				}

				(*stones)[result.stone1] += 1

				// add new stone if any
				if result.has_stone2 {
					_, key_exists := (*stones)[result.new_stone2]
					if !key_exists {
						(*stones)[result.new_stone2] = 0
					}

					(*stones)[result.new_stone2] += 1
				}
			}
		}
	}

}

func countStones(hashmap map[int]int) int {
	total := 0
	for _, count := range hashmap {
		total += count
	}

	return total
}

func main() {
	// defer profile.Start().Stop()
	input, err := libaoc.GrabInput("sample_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// convert to strings
	stone_strings := strings.Split(input, " ")
	stone_ints := make([]int, 0, len(stone_strings))

	// convert to ints
	for _, stone_str := range stone_strings {
		stone_int, err := strconv.Atoi(stone_str)
		if err != nil {
			log.Fatal(err)
		}
		stone_ints = append(stone_ints, stone_int)
	}

	// convert to hashmap
	stone_map := make(map[int]int, 0)
	for _, stone_int := range stone_ints {
		stone_map[stone_int] = 1
	}

	fmt.Println(stone_map)
	for i := 0; i < 75; i++ {
		updateStones(&stone_map)
		fmt.Println("iteration ", i)
	}

	fmt.Println("stones: ", stone_map)
	fmt.Println("stones count: ", countStones(stone_map))
}
