package main

import (
	"os"

	"github.com/cppcho/adventofcode_2023_go/day9"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		panic("invalid argument")
	}

	if args[0] == "day9" {
		day9.SolvePart1()
		day9.SolvePart2()
	} else {
		panic("invalid argument")
	}
}
