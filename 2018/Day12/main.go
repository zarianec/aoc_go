package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	init := strings.Replace(lines[0], "initial state: ", "", -1)
	rules := readRules(lines[2:])

	// Part 1
	fmt.Printf("Sum of pots after 20 gens is: %d\n", evaluatePlants(init, rules, 20))

	// Part 2 - not sure that it's a most correct way and if my constant value will work for other inputs
	// but after some large amount of generations plants become patternized
	// So we don't need to calculate all 50e9 generations and need to know only the difference between 2 gens
	// after pattern is stabilized
	const gensBeforePattern = 125 // It's a minimal amount of steps for my input

	sum1 := evaluatePlants(init, rules, gensBeforePattern)
	sum2 := evaluatePlants(init, rules, gensBeforePattern+1)
	genDiff := sum2 - sum1
	sumFor50e9 := ((50e9 - gensBeforePattern) * genDiff) + sum1

	fmt.Printf("Sum of pots after 50 billions gens is: %d\n", sumFor50e9)

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func evaluatePlants(init string, rules map[string]string, times int) int {
	// Get pots with plants
	plants := make([]int, 0)
	for i := 0; i < len(init); i++ {
		if init[i] == '#' {
			plants = append(plants, i)
		}
	}

	for ; times > 0; times-- {
		curr := make([]int, 0)

		minPot := min(plants)
		maxPot := max(plants)

		for i := minPot-2; i < maxPot+2; i++ {
			pat := ""
			for k := -2; k <= 2; k++ {
				if in(plants, i+k) {
					pat = pat + "#"
				} else {
					pat = pat + "."
				}
			}

			if rules[pat] == "#" {
				curr = append(curr, i)
			}
		}

		// Print the pots to detect a pattern
		//for i := min(curr); i <= max(curr); i++ {
		//	if in(curr, i) {
		//		fmt.Print("#")
		//	} else {
		//		fmt.Print(".")
		//	}
		//}
		//fmt.Println()

		plants = curr
	}

	return sum(plants)
}

func readRules(input []string) map[string]string {
	rules := make(map[string]string)
	for _, r := range input {
		rules[r[:5]] = r[9:]
	}

	return rules
}

// Helpers
// It's a Go - so you need to do all the simple things manually ;)

func min(arr []int) int {
	min := int(1e9)
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}

func max(arr []int) int {
	max := int(-1e9)
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}

func in(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true;
		}
	}

	return false
}

func sum(arr []int) int {
	sum := 0;
	for _, v := range arr {
		sum += v
	}

	return sum
}
