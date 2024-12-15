package main

import (
	"fmt"
	"log"
	"strings"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib" // nb: update to latest with `go get github.com/largenumberhere/AOC_2024/aoc_lib@latest`
)

type Point2D struct {
	x int
	y int
}

func SplitInputs(file_path string) ([][]rune, []Direction, error) {

	lines, err := libaoc.GrabLines(file_path)
	if err != nil {
		return nil, nil, err
	}

	map_lines := [][]rune{}
	directions := []Direction{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			// the map
			map_lines = append(map_lines, []rune(line))
		} else {
			// the instructions
			directions = append(directions, []Direction(line)...)
		}
	}

	return map_lines, directions, nil
}

type Direction rune

const (
	DirectionUp    Direction = '^'
	DirectionRight Direction = '>'
	DirectionDown  Direction = 'v'
	DirectionLeft  Direction = '<'
)

func (direction Direction) String() string {
	return string(rune(direction))
}

func findRobot(field [][]rune) Point2D {
	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field[0]); y++ {
			if field[y][x] == '@' {
				return Point2D{x: x, y: y}
			}
		}
	}

	return Point2D{x: -1, y: -1}
}

func (direction Direction) directionOffset() Point2D {
	robot_pos := Point2D{0, 0}
	switch direction {
	case DirectionUp:
		robot_pos.y -= 1
	case DirectionRight:
		robot_pos.x += 1
	case DirectionDown:
		robot_pos.y += 1
	case DirectionLeft:
		robot_pos.x -= 1
	}

	return robot_pos
}

func (a Point2D) AddComponents(b Point2D) Point2D {
	return Point2D{x: a.x + b.x, y: a.y + b.y}
}

func (a Point2D) isWithin(min Point2D, max Point2D) bool {
	if a.x < min.x {
		return false
	}

	if a.y < min.y {
		return false
	}

	if a.x > max.x {
		return false
	}

	if a.y > max.y {
		return false
	}

	return true
}

func fieldBounds(field [][]rune) (Point2D, Point2D) {
	min := Point2D{x: 0, y: 0}
	max := Point2D{x: len(field), y: len(field[0])}

	return min, max
}

func fieldAt(field [][]rune, pos Point2D) *rune {
	return &(field[pos.y][pos.x])
}

// func isGapInDirection(robot_pos Point2D, field [][]rune, direction Direction) bool {
// 	offset := direction.directionOffset()

// 	// if the position is a box, search it's neigbours, if there is a contiguous line to a wall in direction, false
// 	cursor := robot_pos

// 	for ; cursor.isWithin(fieldBounds(field)); cursor = cursor.AddComponents(offset) {
// 		cell := *fieldAt(field, cursor)
// 		if cell == '.' {
// 			return true
// 		}
// 		if cell == '#' {
// 			return false
// 		}
// 	}

// 	return false
// }

func getFirstGapInDirection(robot_pos Point2D, field [][]rune, direction Direction) (Point2D, bool) {
	offset := direction.directionOffset()

	// if the position is a box, search it's neigbours, if there is a contiguous line to a wall in direction, false
	cursor := robot_pos

	for ; cursor.isWithin(fieldBounds(field)); cursor = cursor.AddComponents(offset) {
		cell := *fieldAt(field, cursor)
		if cell == '.' {
			return cursor, true
		}
		if cell == '#' {
			break
		}
	}

	return Point2D{-1, -1}, false
}

func canRobotMove(field [][]rune, direction Direction, is_gap bool, gap_pos Point2D) bool {
	robot_pos := findRobot(field)

	offset := direction.directionOffset()

	robot_pos = robot_pos.AddComponents(offset)
	// if the position is a wall, false
	if field[robot_pos.y][robot_pos.x] == '#' {
		return false
	}

	// check if there are boxes that can be moved
	if !is_gap {
		return false
	}

	// else the position is empty, true
	return true

}

func tryMoveRobot(starting_robot_pos Point2D, field [][]rune, direction Direction) Point2D {
	direction_offset := direction.directionOffset()
	robot_destination := starting_robot_pos.AddComponents(direction_offset)

	new_robot_pos := Point2D{-1, -1}

	gap_pos, is_gap := getFirstGapInDirection(starting_robot_pos, field, direction)
	if canRobotMove(field, direction, is_gap, gap_pos) {
		// move block into next space if any
		if is_gap == false {
			log.Fatal("unreachable. gap should be populated")
		}
		cell := fieldAt(field, robot_destination)
		if *cell == 'O' {
			*fieldAt(field, gap_pos) = 'O'
		}

		// erase old position
		*fieldAt(field, starting_robot_pos) = '.'

		// write new position
		*fieldAt(field, robot_destination) = '@'
		new_robot_pos = robot_destination
	} else {
		new_robot_pos = starting_robot_pos
	}

	return new_robot_pos
}

func sumBoxes(field [][]rune) int64 {
	var tally int64 = 0

	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field); y++ {
			pos := Point2D{x, y}
			cell := *fieldAt(field, pos)

			if cell == 'O' {
				tally += int64((100 * y) + x)
			}
		}
	}

	return tally

}

func main() {
	if libaoc.VersionCount < 2 {
		log.Fatal("at least version 2 of libaoc is required")
	}

	map_lines, directions, err := SplitInputs("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	libaoc.PrintRunes(map_lines)
	fmt.Println()

	robot_pos := findRobot(map_lines)
	for _, direction := range directions {
		robot_pos = tryMoveRobot(robot_pos, map_lines, direction)
	}

	libaoc.PrintRunes(map_lines)
	fmt.Println(sumBoxes(map_lines))
}
