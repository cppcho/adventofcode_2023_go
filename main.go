package main

import (
	"log"
	"os"

	"github.com/cppcho/adventofcode_2023_go/day10"
	"github.com/cppcho/adventofcode_2023_go/day9"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		panic("invalid argument")
	}

	switch args[0] {
	case "day9":
		day9.SolvePart1()
		day9.SolvePart2()
	case "day10":
		day10.SolvePart1()
		day10.SolvePart2()
	default:
		log.Fatalf("Invalid argument: %s", args[0])
	}
}
