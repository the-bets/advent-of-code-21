package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func findDepths(fn string) (nums []int, err error) {
	a, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(a), "\r\n")
	// avoids resize on append
	nums = make([]int, 0, len(lns))

	for _, y := range lns {
		if len(y) == 0 {
			continue
		}
		n, err := strconv.Atoi(y)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func findIncreasing(lines []int) int {
	// count how many numbers in the ordered list are smaller than the previous
	current := 333333
	ans := 0

	for _, y := range lines {
		if y >= current {
			ans++
		}
		current = y
	}
	return ans
}

func findWindows(numbers []int) int {
	// implementation of sliding window algorithm
	current := 333333
	ans := 0

	for i := 0; i < len(numbers)-2; i++ {
		sum := numbers[i] + numbers[i+1] + numbers[i+2]
		if sum > current {
			ans++
		}
		current = sum
	}
	return ans
}

func main() {
	numbers, err := findDepths("day1.txt")
	if err != nil {
		panic(err)
	}

	answer1 := findIncreasing(numbers)
	fmt.Printf("The answer to challenge one is %d\n", answer1)

	answer2 := findWindows(numbers)
	fmt.Printf("The answer to challenge two is %d\n", answer2)
}
