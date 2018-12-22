package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

var minx, maxx, miny, maxy int

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")

	scan := parseInput(input)
	findPatch(scan)
	//visualise(scan)

	total, stale := countWater(scan)

	fmt.Println(total, stale)
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func findPatch(scan map[Point]string) {
	start := Point{500, 0}
	scan[start] = "+"

	queue := make([]Point, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		cur := queue[0]
		queue = unique(queue[1:])

		if cur.y >= maxy {
			continue
		}

		if isOpen(scan, cur.x, cur.y+1) {
			queue = append(queue, Point{cur.x, cur.y + 1})
			scan[Point{cur.x, cur.y + 1}] = "|"
			continue
		}

		if scan[cur] == "~" {
			//queue = append(queue, Point{cur.x, cur.y - 1})
			continue
		}

		if scan[cur] == "" {
			scan[cur] = "|"
		}

		leftX := cur.x - 1
		for isOpen(scan, leftX, cur.y) && !isOpen(scan, leftX+1, cur.y+1) {
			scan[Point{leftX, cur.y}] = "|"
			leftX--
		}

		rightX := cur.x + 1
		for isOpen(scan, rightX, cur.y) && !isOpen(scan, rightX-1, cur.y+1) {
			scan[Point{rightX, cur.y}] = "|"
			rightX++
		}

		if isOpen(scan, leftX+1, cur.y+1) {
			scan[Point{leftX + 1, cur.y + 1}] = "|"
			queue = append(queue, Point{leftX + 1, cur.y + 1})
		}

		if isOpen(scan, rightX-1, cur.y+1) {
			scan[Point{rightX - 1, cur.y + 1}] = "|"
			queue = append(queue, Point{rightX - 1, cur.y + 1})
		}

		if scan[Point{leftX, cur.y}] == "#" && scan[Point{rightX, cur.y}] == "#" {
			for x := leftX + 1; x <= rightX-1; x++ {
				scan[Point{x, cur.y}] = "~"
			}
			queue = append(queue, Point{cur.x, cur.y - 1})
		}
	}
}

func countWater(scan map[Point]string) (total int, stale int) {
	spring := 0

	for p, t := range scan {
		if p.y < miny {
			continue
		}

		if t == "|" {
			spring++
		}

		if t == "~" {
			stale++
		}
	}

	total = stale + spring

	return
}

func parseInput(input []byte) (scan map[Point]string) {
	scan = make(map[Point]string)
	minx = int(1e9)
	maxx = int(-1e9)
	miny = int(1e9)
	maxy = int(-1e9)

	lines := strings.Split(string(input), "\n")

	for _, l := range lines {
		axis := string(l[0])

		r := strings.NewReplacer("x=", "", "y=", "", ",", "", "..", " ")
		nums := strings.Split(r.Replace(l), " ")

		axisPoint, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		to, _ := strconv.Atoi(nums[2])

		for ; from <= to; from++ {
			p := Point{}
			if axis == "x" {
				p = Point{axisPoint, from}

				minx = int(math.Min(float64(axisPoint), float64(minx)))
				maxx = int(math.Max(float64(axisPoint), float64(maxx)))
				miny = int(math.Min(float64(from), float64(miny)))
				maxy = int(math.Max(float64(from), float64(maxy)))
			} else {
				p = Point{from, axisPoint}

				minx = int(math.Min(float64(from), float64(minx)))
				maxx = int(math.Max(float64(from), float64(maxx)))
				miny = int(math.Min(float64(axisPoint), float64(miny)))
				maxy = int(math.Max(float64(axisPoint), float64(maxy)))
			}

			scan[p] = "#"
		}
	}

	return
}

func visualise(scan map[Point]string) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < maxy+1; y++ {
		for x := minx - 1; x <= maxx+1; x++ {
			p := Point{x, y}

			if val, ok := scan[p]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println()

	//time.Sleep(500 * time.Millisecond)
}

func isOpen(scan map[Point]string, x int, y int) bool {
	p := Point{x, y}

	return scan[p] == "" || scan[p] == "|"
}

func unique(queue []Point) []Point {
	u := make([]Point, 0, len(queue))
	m := make(map[Point]bool)

	for _, val := range queue {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
