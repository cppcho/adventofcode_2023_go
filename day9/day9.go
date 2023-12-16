package day9

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parse() [][]int {
	lines := strings.Split(input, "\n")
	parsed := make([][]int, len(lines))
	for i, line := range lines {
		nums := make([]int, 0)
		for _, str := range strings.Split(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic("invalid num")
			}
			nums = append(nums, num)
		}
		parsed[i] = nums
	}
	return parsed
}

func SolvePart1() {
	parsed := parse()
	output := 0
	for _, nums := range parsed {
		for i := 0; i < len(nums); i++ {
			allZeroes := true
			for j := 1; j < len(nums)-i; j++ {
				nums[j-1] = nums[j] - nums[j-1]
				if nums[j-1] != 0 {
					allZeroes = false
				}
			}
			if allZeroes {
				break
			}
		}
		sum := 0
		for _, v := range nums {
			sum += v
		}
		output += sum
	}
	fmt.Printf("Part1: %d\n", output)
}

func SolvePart2() {
	parsed := parse()
	output := 0
	for _, nums := range parsed {
		for i := 0; i < len(nums); i++ {
			allZeroes := true
			for j := len(nums) - 1; j > i; j-- {
				nums[j] = nums[j] - nums[j-1]
				if nums[j] != 0 {
					allZeroes = false
				}
			}
			if allZeroes {
				break
			}
		}
		curr := 0
		for i := len(nums) - 1; i >= 0; i-- {
			curr = nums[i] - curr
		}
		output += curr
	}
	fmt.Printf("Part2: %d\n", output)
}
