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

func formatMap(robots []RobotDetails, map_bounds Point2D, map_out *[][]rune) {
	// clear the map
	for x := 0; x < map_bounds.x; x++ {
		for y := 0; y < map_bounds.y; y++ {
			(*map_out)[y][x] = '.'
		}
	}

	// update the map with each robot
	for _, r := range robots {
		x := r.current_position.x
		y := r.current_position.y

		if (*map_out)[y][x] == '.' {
			(*map_out)[y][x] = '1'
		} else if (*map_out)[y][x] >= '9' {
			(*map_out)[y][x] = (*map_out)[y][x] - '9'
		} else {
			(*map_out)[y][x] += 1
		}
	}
}

func hasSuspiciousBlock(map_bounds Point2D, formatted_map *[][]rune) bool {
	// vertical line
	contiguous_y := 0
	longest_y := 0
	for x := 0; x < map_bounds.x; x++ {
		for y := 0; y < map_bounds.y; y++ {
			if (*formatted_map)[y][x] != '.' {
				contiguous_y++
			} else {
				if contiguous_y > longest_y {
					longest_y = contiguous_y
				}
				contiguous_y = 0
			}
		}
	}

	// horizontal line
	contiguous_x := 0
	longest_x := 0
	for y := 0; y < map_bounds.y; y++ {
		for x := 0; x < map_bounds.x; x++ {
			if (*formatted_map)[y][x] != '.' {
				contiguous_x++
			} else {
				if contiguous_x > longest_x {
					longest_x = contiguous_x
				}
				contiguous_x = 0
			}
		}
	}

	if longest_y >= 7 && longest_x >= 7 {
		return true
	} else {
		return false
	}

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

	robots = simulateRobotMovement(robots, 7856, map_bounds)
	for i := 0; i < 10; i++ {
		robots = simulateRobotMovement(robots, 1, map_bounds)
		fmt.Println("quadrant sum :", sumQuadreants(robots, map_bounds))
		printMap(robots, map_bounds, false)
	}
	return

	// runes := [][]rune{
	// 	{'1', '.', '.', '.', '.'},
	// 	{'1', '.', '.', '.', '.'},
	// 	{'1', '.', '.', '.', '.'},
	// 	{'1', '.', '.', '.', '.'},
	// 	{'1', '.', '.', '.', '.'},
	// }
	// if (!hasSuspiciousBlock(Point2D{x: 5, y: 5}, &runes)) {
	// 	log.Fatal("error")
	// } else {
	// 	fmt.Println("yippee")
	// }

	tmp := libaoc.MakeRunes(map_bounds.y, map_bounds.x, '.')
	for i := 0; true; i++ {
		robots = simulateRobotMovement(robots, 1, map_bounds)
		formatMap(robots, map_bounds, &tmp)
		// libaoc.PrintRunes(tmp)
		if hasSuspiciousBlock(map_bounds, &tmp) {
			fmt.Println("suspicious vertical block on simulation ", i)
		}

		if i%10000 == 0 {
			fmt.Println("searched ", i)
		}
	}
	/*
		2024
			suspicious vertical block on simulation  181	--
		suspicious vertical block on simulation  441	--
		suspicious vertical block on simulation  484 --
		suspicious vertical block on simulation  585 --
		suspicious vertical block on simulation  632 --
		suspicious vertical block on simulation  686 --
		suspicious vertical block on simulation  738 --
		suspicious vertical block on simulation  781 --
		suspicious vertical block on simulation  882 -- too low
		suspicious vertical block on simulation  888 --
		suspicious vertical block on simulation  1204 --
		suspicious vertical block on simulation  1226 --
		suspicious vertical block on simulation  1436 --
		suspicious vertical block on simulation  1549
		suspicious vertical block on simulation  1696
		suspicious vertical block on simulation  1797
		suspicious vertical block on simulation  1898
		suspicious vertical block on simulation  2295
		suspicious vertical block on simulation  2394
		suspicious vertical block on simulation  2504
		suspicious vertical block on simulation  2604
		suspicious vertical block on simulation  2605
		suspicious vertical block on simulation  2707
		suspicious vertical block on simulation  3009
		suspicious vertical block on simulation  3110
		suspicious vertical block on simulation  3312
		suspicious vertical block on simulation  3413
		suspicious vertical block on simulation  3890
		suspicious vertical block on simulation  3943
		suspicious vertical block on simulation  4124
		suspicious vertical block on simulation  4221
		suspicious vertical block on simulation  4248
		suspicious vertical block on simulation  4252
		suspicious vertical block on simulation  4423
		suspicious vertical block on simulation  4529
		suspicious vertical block on simulation  4726 --
		suspicious vertical block on simulation  4835
		suspicious vertical block on simulation  4838
		suspicious vertical block on simulation  5332
		suspicious vertical block on simulation  5694
		suspicious vertical block on simulation  5864
		suspicious vertical block on simulation  6645
		suspicious vertical block on simulation  7251
		suspicious vertical block on simulation  7342
		suspicious vertical block on simulation  7453
		suspicious vertical block on simulation  7580
		suspicious vertical block on simulation  7599
		suspicious vertical block on simulation  7857
		suspicious vertical block on simulation  7866
		suspicious vertical block on simulation  8023
		suspicious vertical block on simulation  8059
		suspicious vertical block on simulation  8238
		suspicious vertical block on simulation  8475
		suspicious vertical block on simulation  8766
		suspicious vertical block on simulation  9483
		suspicious vertical block on simulation  9574
		suspicious vertical block on simulation  9675
		suspicious vertical block on simulation  9852
		searched  10000
		suspicious vertical block on simulation  10584
		suspicious vertical block on simulation  10844
		suspicious vertical block on simulation  10887
		suspicious vertical block on simulation  10988
		suspicious vertical block on simulation  11035
		suspicious vertical block on simulation  11089
		suspicious vertical block on simulation  11141
		suspicious vertical block on simulation  11184
		suspicious vertical block on simulation  11285
		suspicious vertical block on simulation  11291
		suspicious vertical block on simulation  11607
		suspicious vertical block on simulation  11629
		suspicious vertical block on simulation  11839
		suspicious vertical block on simulation  11952
		suspicious vertical block on simulation  12099
		suspicious vertical block on simulation  12200
		suspicious vertical block on simulation  12301
		suspicious vertical block on simulation  12698
		suspicious vertical block on simulation  12797
		suspicious vertical block on simulation  12907
		suspicious vertical block on simulation  13007
		suspicious vertical block on simulation  13008
		suspicious vertical block on simulation  13110
		suspicious vertical block on simulation  13412
		suspicious vertical block on simulation  13513
		suspicious vertical block on simulation  13715
		suspicious vertical block on simulation  13816
		suspicious vertical block on simulation  14293
		suspicious vertical block on simulation  14346
		suspicious vertical block on simulation  14527
		suspicious vertical block on simulation  14624
		suspicious vertical block on simulation  14651
		suspicious vertical block on simulation  14655
		suspicious vertical block on simulation  14826
		suspicious vertical block on simulation  14932
		suspicious vertical block on simulation  15129
		suspicious vertical block on simulation  15238
		suspicious vertical block on simulation  15241
		suspicious vertical block on simulation  15735
		suspicious vertical block on simulation  16097
		suspicious vertical block on simulation  16267
		suspicious vertical block on simulation  17048
		suspicious vertical block on simulation  17654
		suspicious vertical block on simulation  17745
		suspicious vertical block on simulation  17856
		suspicious vertical block on simulation  17983
		suspicious vertical block on simulation  18002
		suspicious vertical block on simulation  18260
		suspicious vertical block on simulation  18269
		suspicious vertical block on simulation  18426
		suspicious vertical block on simulation  18462
		suspicious vertical block on simulation  18641
		suspicious vertical block on simulation  18878
		suspicious vertical block on simulation  19169
		suspicious vertical block on simulation  19886
		suspicious vertical block on simulation  19977
		searched  20000
		suspicious vertical block on simulation  20078
		suspicious vertical block on simulation  20255
		suspicious vertical block on simulation  20987
		suspicious vertical block on simulation  21247
		suspicious vertical block on simulation  21290
		suspicious vertical block on simulation  21391
		suspicious vertical block on simulation  21438
		suspicious vertical block on simulation  21492
		suspicious vertical block on simulation  21544
		suspicious vertical block on simulation  21587
		suspicious vertical block on simulation  21688
		suspicious vertical block on simulation  21694
		suspicious vertical block on simulation  22010
		suspicious vertical block on simulation  22032
		suspicious vertical block on simulation  22242
		suspicious vertical block on simulation  22355
		suspicious vertical block on simulation  22502
		suspicious vertical block on simulation  22603
		suspicious vertical block on simulation  22704
		suspicious vertical block on simulation  23101
		suspicious vertical block on simulation  23200
		suspicious vertical block on simulation  23310
		suspicious vertical block on simulation  23410
		suspicious vertical block on simulation  23411
		suspicious vertical block on simulation  23513

	*/

	// fmt.Println(getRobotQuadrant(robots[0], map_bounds))
	// printMap(robots, map_bounds, false)
}
