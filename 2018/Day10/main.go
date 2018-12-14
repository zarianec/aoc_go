package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X  int
	Y  int
	vX int
	vY int
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	points := parseInput(input)

	// 15000 is an assumption of how much seconds it may take
	// it's based on input data review and requires Human Brain™ to be used:)
	second := getSmallestBoundarySecond(points, 15000)

	minX := 1e9
	maxX := -1e9
	minY := 1e9
	maxY := -1e9
	pixels := make(map[int]map[int]bool)

	for _, p := range points {
		curX := p.X + (p.vX * second)
		curY := p.Y + (p.vY * second)

		maxX = math.Max(maxX, float64(curX))
		maxY = math.Max(maxY, float64(curY))

		minX = math.Min(minX, float64(curX))
		minY = math.Min(minY, float64(curY))

		if pixels[curX] == nil {
			pixels[curX] = make(map[int]bool)
		}

		pixels[curX][curY] = true
	}

	for i := minY - 1; i < maxY+2; i++ {
		for k := minX - 1; k < maxX+2; k++ {
			if _, ok := pixels[int(k)][int(i)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Seconds needed: %d\n", second)

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func getSmallestBoundarySecond(points []Point, assumedLimit int) int {
	boundaries := make(map[int]int)

	for i := 0; i < assumedLimit; i++ {
		minX := 1e9
		maxX := -1e9
		minY := 1e9
		maxY := -1e9

		for _, p := range points {
			curX := p.X + (p.vX * i)
			curY := p.Y + (p.vY * i)

			maxX = math.Max(maxX, float64(curX))
			maxY = math.Max(maxY, float64(curY))

			minX = math.Min(minX, float64(curX))
			minY = math.Min(minY, float64(curY))
		}

		boundaries[i] = int(maxX - minX + maxY - minY)
	}

	minb := 1e9
	mins := 0
	for s, b := range boundaries {
		if minb > float64(b) {
			minb = float64(b)
			mins = s
		}
	}

	return mins;
}

func parseInput(input []byte) []Point {
	lines := strings.Split(string(input), "\n")
	points := make([]Point, 0)
	r, _ := regexp.Compile("-?[0-9]*\\d")

	for _, l := range lines {
		c := r.FindAllString(l, -1)

		X, _ := strconv.Atoi(c[0])
		Y, _ := strconv.Atoi(c[1])
		vX, _ := strconv.Atoi(c[2])
		vY, _ := strconv.Atoi(c[3])

		point := Point{X, Y, vX, vY}
		points = append(points, point)
	}

	return points
}
