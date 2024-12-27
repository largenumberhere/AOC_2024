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

func updateStoneConcurrent(stone *int, id int) UpdateResult {
	var out1 int
	var out2 int
	used_out2 := updateStone(*stone, &out1, &out2)

	if used_out2 {
		return UpdateResult2(out1, out2, id)
	} else {
		return UpdateResult1(out1, id)
	}
}

func UpdateResult1(stone1 int, id int) UpdateResult {
	return UpdateResult{
		stone1:     stone1,
		id:         id,
		has_stone2: false,
	}
}

func UpdateResult2(stone1 int, stone2 int, id int) UpdateResult {
	return UpdateResult{
		stone1:     stone1,
		new_stone2: stone2,
		id:         id,
		has_stone2: true,
	}
}

type UpdateResult struct {
	stone1     int
	new_stone2 int
	id         int
	has_stone2 bool
}

func createUpdateStoneRoutine(stone *int, id int, output chan UpdateResult) {
	output <- updateStoneConcurrent(stone, id)
}

func processStones(stones *[]int, startIndex, nStonesToProcess int, results *[]UpdateResult, wg *sync.WaitGroup) {
	for i := 0; i < nStonesToProcess; i++ {
		index := startIndex + i

		var out1 int
		var out2 int
		used_out2 := updateStone((*stones)[index], &out1, &out2)
		if used_out2 {
			(*results)[index] = UpdateResult2(out1, out2, index)
		} else {
			(*results)[index] = UpdateResult1(out1, index)
		}
	}

	wg.Done()
}

var numCores = runtime.NumCPU()

func updateStones(stones *[]int) {
	outputs := make(chan UpdateResult, len(*stones))
	results := make([]UpdateResult, len(*stones))

	// If we would only start 1 goroutine per core anyway, use the old method
	if len(*stones) < numCores {
		for i := len(*stones) - 1; i >= 0; i-- {
			go createUpdateStoneRoutine(&(*stones)[i], i, outputs)
		}

		for i := 0; i < len(*stones); i++ {
			out := <-outputs
			results[out.id] = out
		}
	} else {
		// Otherwise, split the work up into N goroutines where N is the core count of your CPU
		var wg sync.WaitGroup

		nStonesToProcessPerGoroutine := len(*stones) / numCores
		startIndex := 0
		for i := 0; i < numCores; i++ {
			nStonesToProcess := nStonesToProcessPerGoroutine
			if i == numCores-1 { // At the last iteration, process all remaining stones
				nStonesToProcess = len(*stones) - nStonesToProcessPerGoroutine*(numCores-1)
			}

			wg.Add(1)
			go processStones(stones, startIndex, nStonesToProcess, &results, &wg)
			startIndex += nStonesToProcess
		}

		wg.Wait()
	}

	// why is this not needed???
	// slices.SortFunc(results, func(a UpdateResult, b UpdateResult) int { return a.id - b.id })

	for i := len(*stones) - 1; i >= 0; i-- {
		(*stones)[i] = results[i].stone1
		if results[i].has_stone2 {
			*stones = append(*stones, results[i].new_stone2)
		}
	}
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

	fmt.Println(stone_ints)
	for i := 0; i < 45; i++ {
		updateStones(&stone_ints)
		fmt.Println("iteration ", i)
	}

	// fmt.Println("stones: ", stones)
	fmt.Println("stones count: ", len(stone_ints))
}
