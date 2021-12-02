package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func grabFile(fn string) []string {
	f, err := os.Open(fn) // can't use ReadFile + strings.Split here like day1
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()

	var lns []string
	sn := bufio.NewScanner(f)
	for sn.Scan() {
		lns = append(lns, sn.Text())
	}
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return lns
}

func covertDirs(s string) (string, int) {
	dirs := strings.Fields(s)
	return dirs[0], convertDirVals(dirs[1])
}

func convertDirVals(s string) int {
	dirVal, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return dirVal
}

func main() {
	// Part One
	data := grabFile("day2.txt")
	horizon, depth, result := 0, 0, 0

	for _, dir := range data {
		dir, val := covertDirs(dir)
		if dir == "forward" {
			horizon += val
		} else if dir == "down" {
			depth += val
		} else {
			depth -= val
		}
	}
	result = depth * horizon

	// Part Two
	horizon, depth, aim, result2 := 0, 0, 0, 0
	for _, dir := range data {
		dir, val := covertDirs(dir)
		if dir == "forward" {
			horizon += val
			depth += aim * val
		} else if dir == "down" {
			aim += val
		} else {
			aim -= val
		}
	}
	result2 = depth * horizon

	fmt.Printf("Part One result is: %d\n", result)
	fmt.Printf("Part Two result is %d\n", result2)
}

//TODO: reformat if:else to be cleaner, maybe with switch cases -- an abstract out some functions
//	for modularity
