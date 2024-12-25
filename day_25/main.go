package main

import (
	"fmt"
	"log"
	"strings"

	aoc_lib "github.com/largenumberhere/AOC_2024/aoc_lib"
)

type LockOrKeyHeights struct {
	heights [5]int
	is_key  bool
}

func (item *LockOrKeyHeights) getHeight(pos int) int {
	return item.heights[pos]
}

func (item *LockOrKeyHeights) setHeight(pos int, value int) {
	item.heights[pos] = value
}

func (item *LockOrKeyHeights) heightsCount() int {
	return len(item.heights)
}

func (item *LockOrKeyHeights) setLock() {
	item.is_key = false
}

func (item *LockOrKeyHeights) setKey() {
	item.is_key = true
}

func (item *LockOrKeyHeights) isKey() bool {
	return item.is_key
}

func parseSchematic(key_lines []string) LockOrKeyHeights {
	var is_lock bool = false
	if key_lines[0][0] == '#' {
		is_lock = true
	}

	item := LockOrKeyHeights{}
	for x := 0; x < len(key_lines[0]); x++ {
		hashes := 0
		for y := 0; y < len(key_lines); y++ {
			if key_lines[y][x] == '#' {
				hashes += 1
			}
		}
		item.setHeight(x, hashes-1)
	}

	if is_lock {
		item.setLock()
	} else {
		item.setKey()
	}

	fmt.Println("item = ", item)

	return item
}

func keyFits(lock LockOrKeyHeights, key LockOrKeyHeights) bool {
	// assert valid states
	if lock.isKey() {
		panic("key passed as lock")
	} else if !key.isKey() {
		panic("lock passed as key")
	}

	for i := 0; i < key.heightsCount(); i++ {
		lockh := lock.getHeight(i)
		keyh := key.getHeight(i)

		if keyh+lockh > 5 {
			return false
		}
	}

	return true
}

func getSchematics(path string) ([]LockOrKeyHeights, error) {
	file_lines, err := aoc_lib.GrabLines(path)
	if err != nil {
		return nil, err
	}

	key := make([]string, 0)
	items := make([]LockOrKeyHeights, 0)
	for i, line := range file_lines {
		if len(strings.Trim(line, "\r\n ")) == 0 {
			continue
		}

		if i%8 == 0 && i != 0 {
			if len(key) > 1 {
				item := parseSchematic(key)
				items = append(items, item)
				key = make([]string, 0)
			}
		}

		key = append(key, line)
	}

	item := parseSchematic(key)
	items = append(items, item)

	return items, nil

}

func countFits(locksAndKeys []LockOrKeyHeights) int {
	fits := 0

	for _, key := range locksAndKeys {
		if !key.isKey() {
			continue
		}
		for _, lock := range locksAndKeys {
			if lock.isKey() {
				continue
			}

			does_fit := keyFits(lock, key)

			// fmt.Println("testing lock, ", lock, "and key ", key, does_fit)

			if does_fit {
				fits += 1
			}

		}
	}

	return fits
}

func main() {
	fmt.Println("")
	schematics, err := getSchematics("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range schematics {
		fmt.Println(item)
	}

	fmt.Println("fits :", countFits(schematics))
}
