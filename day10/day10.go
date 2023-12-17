package day10

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt
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

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

func SolvePart1() {
	tiles, start := parse()
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

	output := 0
	for len(queue) > 0 {
		currMovement := queue[0]
		queue = queue[1:]

		if steps[currMovement.pos.row][currMovement.pos.col] > 0 {
			continue
		}
		steps[currMovement.pos.row][currMovement.pos.col] = currMovement.step
		if currMovement.step > output {
			output = currMovement.step
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

	fmt.Println(steps)
	fmt.Printf("Part1: %d\n", output)
}

func SolvePart2() {
}
