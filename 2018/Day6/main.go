package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	start := time.Now()

	input, err := ioutil.ReadAll(os.Stdin)
	check(err)

	coordinates := strings.Split(string(input), "\n")

	points := make([]Point, len(coordinates))
	minX := 1e9
	minY := 1e9
	maxX := 0e0
	maxY := 0e0

	for i, v := range coordinates {
		p := strings.Split(v, ", ")
		px, _ := strconv.ParseInt(p[0], 10, 64)
		py, _ := strconv.ParseInt(p[1], 10, 64)

		points[i] = Point{int(px), int(py)}

		minX = math.Min(minX, float64(px))
		minY = math.Min(minY, float64(py))

		maxX = math.Max(maxX, float64(px))
		maxY = math.Max(maxY, float64(py))
	}

	scores := make(map[Point]map[string][]int)
	region := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			tmpPoint := Point{int(x), int(y)}

			closest := findClosest(tmpPoint, points)

			if scores[closest] == nil {
				scores[closest] = map[string][]int{}
			}

			scores[closest]["x"] = append(scores[closest]["x"], int(x))
			scores[closest]["y"] = append(scores[closest]["y"], int(y))

			d := 0
			for _, p := range points {
				d += int(distance(p, tmpPoint))

				if d > 10000 {
					break
				}
			}

			if d < 10000 {
				region += 1
			}
		}
	}

	best := 0e0
	for _, p := range scores {
		if in_array(p["x"], int(minX)) || in_array(p["x"], int(maxX)) ||
			in_array(p["y"], int(minY)) || in_array(p["y"], int(maxY)) {
			continue
		}

		best = math.Max(best, float64(len(p["x"])))
	}

	fmt.Printf("Largest area size: %d\n", int(best))
	fmt.Printf("Best region size: %d\n", region)

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func findClosest(point Point, points []Point) (closest Point) {
	closest = points[0];
	shared := false

	for i := 1; i < len(points); i++ {
		p := points[i]

		dist := distance(point, p)
		bDist := distance(point, closest)

		if dist < bDist {
			closest = p
			shared = false
		} else if dist == bDist {
			shared = true
		}
	}

	if shared {
		return Point{-1, -1}
	}

	return closest
}

func distance(p1 Point, p2 Point) (d float64) {
	d = math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y))

	return d
}

func in_array(haystack []int, needle int) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
