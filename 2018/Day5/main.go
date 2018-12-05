package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

func main() {
	polymer, err := ioutil.ReadFile("input.txt")
	check(err)

	start := time.Now()

	//
	// Part 1
	//
	polymer = reactPolymer(polymer)

	fmt.Println("Part 1:")
	fmt.Printf("Polymer size: %d", len(polymer))
	fmt.Println()

	//
	// Part 2
	//
	elements := "abcdefghijklmnopqrstuvwxyz"
	stringPolymer := string(polymer)
	minSize := len(stringPolymer);
	for i := 0; i < len(elements); i++ {
		pos := string(elements[i])
		neg := string(elements[i] - 32);
		r := strings.NewReplacer(pos, "", neg, "")
		optimizedSize := len(reactPolymer([]byte(r.Replace(stringPolymer))))

		if minSize > optimizedSize {
			minSize = optimizedSize
		}
	}

	fmt.Println("Part2:")
	fmt.Printf("Optimized polymer size: %d", minSize)
	fmt.Println()

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func reactPolymer(polymer []byte) (reactedPolymer []byte) {
	for i := 0; i < len(polymer)-1; i++ {
		cur := float64(polymer[i]);
		next := float64(polymer[i+1]);
		if math.Abs(cur-next) == 32 {
			polymer = append(polymer[:i], polymer[i+2:]...)
			i = int(math.Max(float64(i-2), -1))
		}
	}

	return polymer
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
