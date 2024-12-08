package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directionMap = map[Direction]string{
	Up:    "Up",
	Down:  "Down",
	Left:  "Left",
	Right: "Right",
}

func (direction Direction) String() string {
	return directionMap[direction]
}

func (direction *Direction) Turn() {
	switch *direction {
	case Up:
		*direction = Right
		break
	case Right:
		*direction = Down
		break
	case Down:
		*direction = Left
		break
	case Left:
		*direction = Up
		break
	}
}

type GuardMoevemnt int

const (
	Obstructed GuardMoevemnt = iota
	OutsideMap
	Moved
)

var guardMoevemntMap = map[GuardMoevemnt]string{
	Moved:      "Moved",
	Obstructed: "Obstructed",
	OutsideMap: "OutsideMap",
}

func (movement GuardMoevemnt) String() string {
	return guardMoevemntMap[movement]
}

func printRunes(runes *[][]rune) {
	for _, i := range *runes {
		for _, j := range i {
			if j != '.' && j != '^' && j != '#' && j != 'x' {
				fmt.Print("?")
			} else {
				fmt.Print(string(j))
			}
		}
		fmt.Println("")
	}

}

type Point2D struct {
	x int
	y int
}

func findGuard(runes *[][]rune) (Point2D, error) {
	// find the guard

	for i, row := range *runes {
		for j, item := range row {
			if item == '^' {
				return Point2D{y: i, x: j}, nil
			}
		}
	}

	return Point2D{y: -1, x: -1}, errors.New("no guard found")
}

func moveGuard(runes *[][]rune, direction Direction) GuardMoevemnt {
	pos, err := findGuard(runes)
	if err != nil {
		log.Fatal(err)
	}

	// calculate new gard position
	pos2 := pos
	switch direction {
	case Up:
		pos2.y -= 1
		break
	case Down:
		pos2.y += 1
		break
	case Left:
		pos2.x -= 1
		break
	case Right:
		pos2.x += 1
		break
	}

	// // handle bounds checking
	maxy := len(*runes)
	maxx := len((*runes)[0])
	out_of_bounds := pos2.x < 0 || pos2.y < 0 || pos2.x >= maxx || pos2.y >= maxy
	if out_of_bounds {
		// replace prev position with X
		(*runes)[pos.y][pos.x] = 'x'
		return OutsideMap
	}

	obstructed := (*runes)[pos2.y][pos2.x] == '#'
	if obstructed {
		// check if obstructed
		return Obstructed
	}

	// create guard at required position
	(*runes)[pos2.y][pos2.x] = '^'

	// replace prev position with X
	(*runes)[pos.y][pos.x] = 'x'

	return Moved
}

func simulateGuard(lines_runes *[][]rune) {
	// printRunes(lines_runes)
	direction := Up
	for {
		res := moveGuard(lines_runes, direction)
		if res == Obstructed {
			// todo update direction
			direction.Turn()
		} else if res == OutsideMap {
			break
		} else {

		}

		// printRunes(lines_runes)
	}

	// printRunes(lines_runes)
}

func countVisited(lines_runes *[][]rune) int {
	var count int

	for _, row := range *lines_runes {
		for _, item := range row {
			if item == 'x' {
				count += 1
			}
		}
	}

	return count
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file_contents := strings.Trim(string(bytes), " \n")
	lines := strings.Split(file_contents, "\n")

	var lines_runes [][]rune
	for _, line := range lines {
		var line_rune []rune
		for _, rune := range line {
			if rune == '\r' {
				continue
			}
			line_rune = append(line_rune, rune)
		}

		lines_runes = append(lines_runes, line_rune)
	}

	simulateGuard(&lines_runes)
	visited := countVisited(&lines_runes)
	printRunes(&lines_runes)
	fmt.Println(visited)

}
