package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	libaoc "github.com/largenumberhere/AOC_2024/aoc_lib"
)

// func RegexLines(path string, regex_expression_each_line string) ([]string, error) {
// 	input, err := libaoc.GrabInput("input.txt")
// 	if err != nil {
// 		return nil, err
// 	}

// 	regex, err := regexp.Compile(regex_expression_each_line)
// 	if err != nil {
// 		return nil, err
// 	}

// 	regex.

// 	return []string{input}, nil
// }

type Point2D struct {
	x int
	y int
}

func isBottomRightOf(point_a Point2D, point_b Point2D) bool {
	return (point_a.x > point_b.x) && (point_a.y > point_b.y)
}

func isTopLeftOf(point_a Point2D, point_b Point2D) bool {
	return (point_a.x < point_b.x) && (point_a.y < point_b.y)
}

type RobotDetails struct {
	starting_position Point2D
	velocity          Point2D
	current_position  Point2D
}

func parseLine(line string) (RobotDetails, error) {
	var details RobotDetails

	// line p=0,4 v=3,-3

	// items {"p=0,4," "v=3,-3"}
	items := strings.Split(line, " ")

	// left {"p", "0,4,"}
	left_parts := strings.Split(items[0], "=")
	// left_values "0,4,"
	left_values := left_parts[1]

	// left_values_trimmed = {"0", "4", ""}
	left_values_trimmed := strings.Split(left_values, ",")

	// extract the starting_point value
	for i, number := range left_values_trimmed {
		if number == "" {
			continue
		}

		number_as_int, err := strconv.Atoi(number)
		if err != nil {
			return RobotDetails{}, err
		}

		if i == 0 {
			details.starting_position.x = number_as_int
			details.current_position.x = number_as_int
		} else if i == 1 {
			details.starting_position.y = number_as_int
			details.current_position.y = number_as_int
		} else {
			panic("out of bounds")
		}

	}

	// right_parts {"v", "3,-3"}
	right_parts := strings.Split(items[1], "=")
	right_values := right_parts[1]
	right_values_trimmed := strings.Split(right_values, ",")

	for i, number := range right_values_trimmed {
		if number == "" {
			continue
		}

		number_as_int, err := strconv.Atoi(number)
		if err != nil {
			return RobotDetails{}, err
		}

		if i == 0 {
			details.velocity.x = number_as_int
		} else if i == 1 {
			details.velocity.y = number_as_int
		} else {
			panic("out of bounds")
		}

	}

	return details, nil
}

func wrapAround(map_bounds Point2D, pos Point2D) Point2D {
	for pos.x < 0 {
		pos.x += map_bounds.x
	}
	for pos.y < 0 {
		pos.y += map_bounds.y
	}

	for pos.x >= map_bounds.x {
		pos.x -= map_bounds.x
	}
	for pos.y >= map_bounds.y {
		pos.y -= map_bounds.y
	}

	return pos
}

func simulateRobotMovement(robots []RobotDetails, ticks int, map_bounds Point2D) []RobotDetails {
	for i := 0; i < ticks; i++ {
		for r := 0; r < len(robots); r++ {
			// calculate robot destination
			new_x := robots[r].current_position.x + robots[r].velocity.x
			new_y := robots[r].current_position.y + robots[r].velocity.y
			destination := wrapAround(map_bounds, Point2D{new_x, new_y})
			// fmt.Println("deset: ", destination)

			robots[r].current_position = destination
		}
	}

	return robots
}

func getRobotQuadrant(robot RobotDetails, map_bounds Point2D) int {
	pos := robot.current_position
	quad := -1
	if pos.x > (map_bounds.x / 2) {
		//1 or 3
		if pos.y > (map_bounds.y / 2) {
			// 3
			quad = 3
		} else if pos.y == (map_bounds.y / 2) {
			// none
		} else if pos.y < (map_bounds.y / 2) {
			// 1
			quad = 1
		}
	} else if pos.x == (map_bounds.x / 2) {
		// none
	} else if pos.x < (map_bounds.x / 2) {
		// 0 or 2
		if pos.y > (map_bounds.y / 2) {
			// 2
			quad = 2
		} else if pos.y < (map_bounds.y / 2) {
			// 0
			quad = 0
		} else if pos.y == (map_bounds.y / 2) {
			// none
		}

	}

	return quad

	// q0_max := Point2D{x: map_bounds.x / 2, y: map_bounds.y / 2}

	// q1_min := Point2D{x: map_bounds.x / 2, y: 0}
	// q1_max := Point2D{x: map_bounds.x, y: map_bounds.y / 2}

	// q2_min := Point2D{x: 0, y: map_bounds.y / 2}
	// q2_max := Point2D{x: map_bounds.x / 2, y: map_bounds.y}

	// q3_min := Point2D{x: map_bounds.x / 2, y: map_bounds.y / 2}

	// quad := -1
	// if isTopLeftOf(robot.current_position, q0_max) {
	// 	quad = 0

	// 	// bugged
	// } else if isBottomRightOf(robot.current_position, q1_min) && isTopLeftOf(robot.current_position, q1_max) {
	// 	quad = 1
	// } else if isBottomRightOf(robot.current_position, q2_min) && isTopLeftOf(robot.current_position, q2_max) {
	// 	quad = 2
	// } else if isBottomRightOf(robot.current_position, q3_min) {
	// 	quad = 3
	// }
	// return quad
}

func sumQuadreants(robots []RobotDetails, map_bounds Point2D) int64 {
	q0_count := int64(0)
	q1_count := int64(0)
	q2_count := int64(0)
	q3_count := int64(0)
	invalids := int64(0)

	for _, robot := range robots {
		quad := getRobotQuadrant(robot, map_bounds)
		switch quad {
		case 0:
			q0_count++
		case 1:
			q1_count++
		case 2:
			q2_count++
		case 3:
			q3_count++
		default:
			invalids += 1
			// on a boundary
		}
	}

	fmt.Println(q0_count, q1_count, q2_count, q3_count, invalids)

	return (q0_count * q1_count * q2_count * q3_count)
}

func printMap(robots []RobotDetails, map_bounds Point2D, show_quads bool) {
	map_ := libaoc.MakeRunes(map_bounds.y, map_bounds.x, '.')

	if show_quads {
		// create the quads
		for y := 0; y < len(map_); y++ {
			map_[y][map_bounds.x/2] = ' '
		}
		for x := 0; x < len(map_[0]); x++ {
			map_[map_bounds.y/2][x] = ' '
		}
	}

	for _, r := range robots {
		x := r.current_position.x
		y := r.current_position.y

		if map_[y][x] == '.' {
			map_[y][x] = '1'
		} else if map_[y][x] >= '9' {
			map_[y][x] = map_[y][x] - '9'
		} else {
			map_[y][x] += 1
		}
	}

	libaoc.PrintRunes(map_)
	fmt.Println("Quadrant sum: ", sumQuadreants(robots, map_bounds))
	fmt.Println()
}

func main() {
	values, err := libaoc.GrabLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots := make([]RobotDetails, 0, 500)
	for _, line := range values {
		robot, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		robots = append(robots, robot)
	}

	map_bounds := Point2D{x: 101, y: 103}
	robots = simulateRobotMovement(robots, 100, map_bounds)
	// fmt.Println(robots)

	fmt.Println("quadrant sum :", sumQuadreants(robots, map_bounds))
	// fmt.Println(getRobotQuadrant(robots[0], map_bounds))
	// printMap(robots, map_bounds, false)
}
