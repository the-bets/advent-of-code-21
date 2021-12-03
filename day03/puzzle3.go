package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := grabBinaries("day3.txt")
	ans1 := powerLevel(data)

	ansOx, _ := strconv.ParseUint(lifeSupport(data, true, 0), 2, 64)
	ansCO, _ := strconv.ParseUint(lifeSupport(data, false, 0), 2, 64)
	ans2 := ansOx * ansCO

	fmt.Printf("The answer to part one is %d\n", ans1)
	fmt.Printf("The answer to part two is %d\n", ans2)
}

func grabBinaries(fn string) []string {
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

func powerLevel(slc []string) int {
	binaries := len(slc[0])
	var g strings.Builder
	var e strings.Builder

	for x := 0; x < binaries; x++ {
		var ones, zeroes = 0, 0
		for _, y := range slc {
			if y[x] == '0' {
				zeroes++
			} else if y[x] == '1' {
				ones++
			}
		}
		if ones > zeroes {
			g.WriteString("1")
			e.WriteString("0")
		} else {
			g.WriteString("0")
			e.WriteString("1")
		}
	}
	gamma, _ := strconv.ParseInt(g.String(), 2, 64)
	epsilon, _ := strconv.ParseInt(e.String(), 2, 64)
	solution := int(gamma * epsilon)
	return solution
}

// I'm going to use recursion because I started iteratively and it got messy...
func lifeSupport(slc []string, common bool, index int) string {
	if len(slc) == 1 {
		return slc[0]
	}

	oSlc, cSlc := make([]string, 0), make([]string, 0)

	for _, num := range slc {
		if num[index] == '1' {
			oSlc = append(oSlc, num)
		} else {
			cSlc = append(cSlc, num)
		}
	}
	if len(oSlc) >= len(cSlc) == common {
		return lifeSupport(oSlc, common, index+1) // a la instructions, can also be read as 'return value(s) with One'
	}
	return lifeSupport(cSlc, common, index+1) // a la instructions, can also be read as 'return value(s) with Zero'
}
