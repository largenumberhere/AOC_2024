package main

import (
	"bytes"
	"fmt"
	"iter" // requires go 1.23
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib"
	"github.com/pkg/profile"
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
		return UpdateResult2(stone, out1, out2, 1)
	} else {
		return UpdateResult1(stone, out1, 1)
	}
}

func UpdateResult1(stone_in int, stone1 int, count int) UpdateResult {
	return UpdateResult{
		stone_in:       stone_in,
		stone_out1:     stone1,
		has_stone_out2: false,
		count:          count,
	}
}

func UpdateResult2(stone_in int, stone1 int, stone2 int, count int) UpdateResult {
	return UpdateResult{
		stone_in:       stone_in,
		stone_out1:     stone1,
		stone_out2:     stone2,
		has_stone_out2: true,
		count:          count,
	}
}

type UpdateResult struct {
	stone_in       int
	stone_out1     int
	stone_out2     int
	has_stone_out2 bool
	count          int
}

type Stones Bag[int]

func (stones Stones) String() string {
	bag := (Bag[int])(stones)

	bag_str := bag.Format(func(i int) string { return strconv.Itoa(i) })
	return bag_str
}

func NewStones() *Stones {
	bag := NewBag[int]()
	return (*Stones)(bag)
}

func (stones *Stones) Insert(item int) {
	bag := (*Bag[int])(stones)

	bag.Insert(item)
}

func (stones *Stones) InsertCount(item int, count int) {
	bag := (*Bag[int])(stones)

	bag.InsertCount(item, count)
}

func (stones *Stones) Remove(item int) (int, bool) {
	bag := (*Bag[int])(stones)
	return bag.Remove(item)
}

func (stones *Stones) RemoveCount(item int, count int) (int, int) {

	bag := (*Bag[int])(stones)
	item, given := bag.RemoveCount(item, count)

	if given != count {
		panic("some removals failed")
	}

	return item, count
}

func (stones *Stones) Replace(item int, with int) bool {
	bag := (*Bag[int])(stones)
	return bag.Replace(item, with)
}

func (stones Stones) Items() iter.Seq[int] {
	bag := (Bag[int])(stones)

	return bag.Items()
}

func (stones Stones) UniqueItems() iter.Seq2[int, int] {
	bag := (Bag[int])(stones)

	return bag.UniqueItems()
}

func (stones *Stones) Count() int {
	bag := (*Bag[int])(stones)
	return bag.Count()
}

type Bag[T1 comparable] struct {
	inner map[T1]int
}

// Insert one of the given item into the bag
func (bag *Bag[T]) Insert(item T) {
	_, ok := bag.inner[item]
	if !ok {
		bag.inner[item] = 0
	}

	bag.inner[item] += 1
}

func (bag *Bag[T]) InsertCount(item T, count int) {
	_, ok := bag.inner[item]
	if !ok {
		bag.inner[item] = 0
	}

	bag.inner[item] += count
}

// Remove one of the given item from the bag
func (bag *Bag[T]) Remove(item T) (T, bool) {
	_, ok := bag.inner[item]
	if !ok {
		var defaultT T
		return defaultT, false
	}

	bag.inner[item] -= 1

	if bag.inner[item] == 0 {
		delete(bag.inner, item)
	}

	return item, true
}

func (bag *Bag[T]) RemoveCount(item T, count int) (T, int) {
	count_existing, ok := bag.inner[item]

	if !ok {
		var defaultT T
		return defaultT, 0
	}

	can_remove := min(count, count_existing)
	count_existing -= can_remove
	bag.inner[item] = count_existing

	return item, can_remove
}

func (bag *Bag[T]) Replace(item T, with T) bool {
	// no-op
	if item == with {
		return true
	}

	_, ok := bag.Remove(item)
	if !ok {
		return false
	}

	bag.Insert(with)
	return true
}

func (bag *Bag[T]) Count() int {
	tally := 0
	for _, count := range bag.inner {
		tally += count
	}

	return tally
}

func NewBag[T comparable]() *Bag[T] {
	hashmap := make(map[T]int)
	bag := Bag[T]{
		inner: hashmap,
	}

	return &bag
}

func (bag *Bag[T]) Items() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item, count := range bag.inner {
			for i := 0; i < count; i++ {
				if !yield(item) {
					return
				}
			}
		}
	}
}

func (bag *Bag[T]) UniqueItems() iter.Seq2[T, int] {
	return func(yield func(T, int) bool) {
		for item, count := range bag.inner {
			if !yield(item, count) {
				return

			}
		}
	}
}

func (bag *Bag[T]) Format(formatFunc func(T) string) string {
	buff := bytes.Buffer{}

	buff.WriteString("Bag {")
	for item, count := range (*bag).UniqueItems() {
		buff.WriteString("\n    ")
		buff.WriteString(formatFunc(item))
		buff.WriteString(": ")
		buff.WriteString(strconv.Itoa(count))

	}
	buff.WriteString("\n}\n")

	return buff.String()
}

var numCores = runtime.NumCPU()

type Tuple struct {
	a int
	b int
}

func updateStones(stones *Stones) {
	if stones.Count() <= numCores {
		waits := sync.WaitGroup{}
		output := make(chan UpdateResult, 10)
		for range stones.Items() {
			waits.Add(1)
		}

		for value := range stones.Items() {
			go func(output chan UpdateResult, key int, wg *sync.WaitGroup) {
				defer wg.Done()
				output <- updateStoneConcurrent(key)

			}(output, value, &waits)
		}

		go func() {
			waits.Wait()
			close(output)
		}()

		for result := range output {
			// handle stone replacement
			if result.stone_out1 != result.stone_in {
				stones.Replace(result.stone_in, result.stone_out1)
			}

			// add new stone if any
			if result.has_stone_out2 {
				stones.Insert(result.stone_out2)
			}
		}
	} else {
		wg := sync.WaitGroup{}

		inputs := make(chan Tuple, 100000)
		outputs := make(chan UpdateResult, 10000)

		// spin up channels
		for range numCores {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for input_tuple := range inputs {
					result := updateStoneConcurrent(input_tuple.a)
					result.count = input_tuple.b

					outputs <- result
				}
			}()
		}

		// dispatch work
		for key, count := range stones.UniqueItems() {
			inputs <- Tuple{a: key, b: count}
		}

		close(inputs)

		go func() {
			wg.Wait()
			close(outputs)
		}()

		// save the results
		for result := range outputs {

			if result.stone_in != result.stone_out1 {
				stones.RemoveCount(result.stone_in, result.count)
				stones.InsertCount(result.stone_out1, result.count)
			}

			//stones.Replace(result.stone_in, result.stone_out1)
			if result.has_stone_out2 {
				stones.InsertCount(result.stone_out2, result.count)
			}
		}

	}
}

func main() {
	fmt.Println("cores detected: ", numCores)

	defer profile.Start().Stop()
	input, err := libaoc.GrabInput("input.txt")

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

	// convert to stones
	stones := NewStones()
	for _, stone_int := range stone_ints {
		stones.Insert(stone_int)
	}

	fmt.Println(stones)
	previous_time := time.Now()
	for i := 0; i < 75; i++ {
		updateStones(stones)
		duration := time.Since(previous_time)

		fmt.Println("iteration ", i+1, " took", duration)

		previous_time = time.Now()
	}

	fmt.Println("stones: ", stones)
	fmt.Println("stones count: ", stones.Count())
}
