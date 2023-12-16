package day9

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func SolvePart1() {
	lines := strings.Split(input, "\n")
	output := 0
	for _, line := range lines {
		nums := make([]int, 0)
		for _, str := range strings.Split(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic("invalid num")
			}
			nums = append(nums, num)
		}
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

}
