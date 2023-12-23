package day10

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
)

//go:embed input_example2.txt
var input string

type position struct {
	row int
	col int
}

type movement struct {
	pos  position
	step int
}

func parse() ([][]rune, position) {
	lines := strings.Split(input, "\n")
	tiles := make([][]rune, 0)
	var start position
	for i, line := range lines {
		row := make([]rune, 0)
		for j, v := range line {
			row = append(row, v)
			if v == 'S' {
				start = position{row: i, col: j}
			}
		}
		tiles = append(tiles, row)
	}
	return tiles, start
}

func getPipe(tiles [][]rune, pos position) rune {
	if pos.row < 0 || pos.col < 0 || pos.row >= len(tiles) || pos.col >= len(tiles[0]) {
		return '.'
	}
	return tiles[pos.row][pos.col]
}

func getMainLoopSteps(tiles [][]rune, start position) ([][]int, int) {
	queue := make([]movement, 0)

	// Handle start position
	topPos := position{row: start.row - 1, col: start.col}
	rightPos := position{row: start.row, col: start.col + 1}
	bottomPos := position{row: start.row + 1, col: start.col}
	leftPos := position{row: start.row, col: start.col - 1}
	if slices.Contains([]rune{'|', '7', 'F'}, getPipe(tiles, topPos)) {
		queue = append(queue, movement{pos: topPos, step: 1})
	}
	if slices.Contains([]rune{'-', 'J', '7'}, getPipe(tiles, rightPos)) {
		queue = append(queue, movement{pos: rightPos, step: 1})
	}
	if slices.Contains([]rune{'|', 'J', 'L'}, getPipe(tiles, bottomPos)) {
		queue = append(queue, movement{pos: bottomPos, step: 1})
	}
	if slices.Contains([]rune{'-', 'F', 'L'}, getPipe(tiles, leftPos)) {
		queue = append(queue, movement{pos: leftPos, step: 1})
	}

	steps := make([][]int, len(tiles))
	for i := range steps {
		steps[i] = make([]int, len(tiles[0]))
	}

	steps[start.row][start.col] = 0

	maxStep := 0
	for len(queue) > 0 {
		currMovement := queue[0]
		queue = queue[1:]

		if steps[currMovement.pos.row][currMovement.pos.col] > 0 {
			continue
		}
		steps[currMovement.pos.row][currMovement.pos.col] = currMovement.step
		if currMovement.step > maxStep {
			maxStep = currMovement.step
		}

		// move to the next tile
		currPipe := tiles[currMovement.pos.row][currMovement.pos.col]
		topPos := position{row: currMovement.pos.row - 1, col: currMovement.pos.col}
		rightPos := position{row: currMovement.pos.row, col: currMovement.pos.col + 1}
		bottomPos := position{row: currMovement.pos.row + 1, col: currMovement.pos.col}
		leftPos := position{row: currMovement.pos.row, col: currMovement.pos.col - 1}

		// Assume next step must be valid pipe
		switch currPipe {
		case '|':
			queue = append(queue, movement{pos: topPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: bottomPos, step: currMovement.step + 1})
		case '-':
			queue = append(queue, movement{pos: leftPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: rightPos, step: currMovement.step + 1})
		case 'L':
			queue = append(queue, movement{pos: rightPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: topPos, step: currMovement.step + 1})
		case 'J':
			queue = append(queue, movement{pos: leftPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: topPos, step: currMovement.step + 1})
		case '7':
			queue = append(queue, movement{pos: leftPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: bottomPos, step: currMovement.step + 1})
		case 'F':
			queue = append(queue, movement{pos: rightPos, step: currMovement.step + 1})
			queue = append(queue, movement{pos: bottomPos, step: currMovement.step + 1})
		case 'S':
		default:
			log.Fatalf("invalid tile %c", currPipe)
		}
	}
	return steps, maxStep
}

func SolvePart1() {
	tiles, start := parse()
	_, output := getMainLoopSteps(tiles, start)
	fmt.Printf("Part1: %d\n", output)
}

func SolvePart2() {
	tiles, start := parse()
	steps, _ := getMainLoopSteps(tiles, start)

	// mark those tiles which have been searched with -1
	steps[start.row][start.col] = -1

	output := 0

	for row, cols := range steps {
		for col, colValue := range cols {
			if colValue != 0 {
				// skip main loop or searched
				continue
			}

			// start bts
			enclosed := true
			queue := make([]position, 0)
			queue = append(queue, position{row: row, col: col})
			size := 0

			for len(queue) > 0 {
				// dequeue
				curr := queue[0]
				queue = queue[1:]

				if curr.row < 0 || curr.row >= len(steps) || curr.col < 0 || curr.col >= len(cols) {
					enclosed = false
					continue
				}

				if steps[curr.row][curr.col] != 0 {
					continue
				}

				steps[curr.row][curr.col] = -1
				size++

				// top
				queue = append(queue, position{row: curr.row - 1, col: curr.col})
				// top right
				queue = append(queue, position{row: curr.row - 1, col: curr.col + 1})
				// right
				queue = append(queue, position{row: curr.row, col: curr.col + 1})
				// bottom right
				queue = append(queue, position{row: curr.row + 1, col: curr.col + 1})
				// bottom
				queue = append(queue, position{row: curr.row + 1, col: curr.col})
				// bottom left
				queue = append(queue, position{row: curr.row + 1, col: curr.col - 1})
				// left
				queue = append(queue, position{row: curr.row, col: curr.col - 1})
				// top left
				queue = append(queue, position{row: curr.row - 1, col: curr.col - 1})
			}

			if enclosed {
				output += size
			}
		}
	}

	fmt.Println(steps)

	fmt.Printf("Part2: %d\n", output)
}
