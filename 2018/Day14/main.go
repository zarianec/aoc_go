package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	r1 := findBest(635041)
	r2 := findPosition(635041)

	fmt.Printf("Best receipt: %s\n", r1)
	fmt.Printf("Input appear first at: %d\n", r2)

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

// Part1
func findBest(n int) string {
	recipes := make([]int, 2)
	recipes[0] = 3
	recipes[1] = 7

	e1 := 0
	e2 := 1
	for len(recipes) < n+10 {
		r1 := recipes[e1]
		r2 := recipes[e2]
		nr := r1 + r2
		parts := splitNum(nr)
		recipes = addRecipe(recipes, parts)

		e1 = nextPos(e1, r1, len(recipes))
		e2 = nextPos(e2, r2, len(recipes))
	}

	return intsToString(recipes[n : n+10])
}

// Part 2
func findPosition(n int) int {
	nstr := strconv.Itoa(n)

	recipes := make([]int, 2)
	recipes[0] = 3
	recipes[1] = 7

	e1 := 0
	e2 := 1
	count := 0
	for count == 0 {
		r1 := recipes[e1]
		r2 := recipes[e2]
		nr := r1 + r2
		parts := splitNum(nr)
		recipes = addRecipe(recipes, parts)

		if len(recipes) > len(nstr) {
			ok := false
			if count, ok = compareLast(recipes, nstr); ok {
				break
			}
		}

		e1 = nextPos(e1, r1, len(recipes))
		e2 = nextPos(e2, r2, len(recipes))
	}

	return count
}

func addRecipe(list []int, new []int) []int {
	if len(new) != 0 {
		list = append(list, new...)
	} else {
		// Argh! Stupid Go don't want to add 0 to slice and think that it's empty
		list = append(list, 0)
	}

	return list
}

func compareLast(l []int, s string) (int, bool) {
	last := l[len(l)-len(s):]
	if intsToString(last) == s {
		return len(l) - len(s), true
	}

	last = l[len(l)-len(s)-1 : len(l)-1]
	if intsToString(last) == s {
		return len(l) - len(s) - 1, true
	}

	return 0, false
}

func nextPos(curIdx int, curVal int, len int) int {
	newP := (len + ((curVal + 1) + curIdx)) % len

	return newP
}

func splitNum(num int) []int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}

	sort.Slice(digits, func(i, j int) bool { return i > j })
	return digits
}

func intsToString(ints []int) string {
	str := ""
	for _, d := range ints {
		str += strconv.Itoa(d)
	}

	return str
}
