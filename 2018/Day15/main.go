package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

// Simple point type
type Point struct {
	x int
	y int
}

type Node struct {
	p Point
	d int
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")

	units, walls, maxx, maxy := parseInput(input)

	units1 := make([]Unit, len(units))
	copy(units1, units)
	runPart1(units1, walls, maxx, maxy)

	units2 := make([]Unit, len(units))
	copy(units2, units)
	runPart2(units2, walls, maxx, maxy)

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func runPart1(units []Unit, walls []Point, maxx int, maxy int) {
	battleFinished := false
	roundsCount := -1
	for !battleFinished {
		units, battleFinished, roundsCount = tick(units, walls, roundsCount, maxx, maxy)
	}

	fmt.Println(calculateOutcome(roundsCount, units))
}

func runPart2(units []Unit, walls []Point, maxx int, maxy int) {
	battleFinished := false
	deadElves := false
	roundsCount := -1

	damage := 3
	for true {
		curUnits := make([]Unit, len(units))
		copy(curUnits, units)

		deadElves = false
		roundsCount = -1
		damage++

		for !deadElves && !battleFinished {
			curUnits = giveNanoWeaponToElves(curUnits, damage)
			curUnits, battleFinished, roundsCount = tick(curUnits, walls, roundsCount, maxx, maxy)

			for _, u := range curUnits {
				if u.race == "E" && u.dead == true {
					battleFinished = false
					deadElves = true
					break
				}
			}
		}

		if battleFinished {
			fmt.Println(calculateOutcome(roundsCount, curUnits))
			break
		}
	}
}

func tick(units []Unit, walls []Point, roundsCount int, maxx int, maxy int) ([]Unit, bool, int) {
	//visualize(walls, units, maxx, maxy)

	battleFinished := false
	moveLog := make(map[int]bool)
	for y := 0; y < maxy; y++ {
		if battleFinished {
			break
		}

		for x := 0; x < maxx; x++ {
			battleFinished = checkIfBattleIsFinished(units)
			if battleFinished {
				break
			}

			idx, u := findByPoint(Point{x, y}, units)
			if u.dead || moveLog[idx] == true {
				continue
			}
			moveLog[idx] = true

			targets := buildTargets(u, units)
			closest, moves := findClosestTarget(u, targets, units, walls, maxx, maxy)
			closest = sortInReadOrder(closest)

			if len(closest) == 0 {
				tryToHitSomeone(u, units)
				continue
			}

			path := buildPath(moves, u.toPoint(), closest[0])
			// Most probably path contains the current unit itself
			if len(path) == 0 {
				tryToHitSomeone(u, units)
				continue
			}

			np := path[0]
			u.x, u.y = np.x, np.y
			units[idx] = u

			tryToHitSomeone(u, units)
		}
	}

	roundsCount++

	return units, battleFinished, roundsCount
}

func parseInput(input []byte) (units []Unit, walls []Point, maxx int, maxy int) {
	walls = make([]Point, 0)
	maxx = 0
	maxy = 0

	lines := strings.Split(string(input), "\n")
	for y, l := range lines {
		tiles := strings.Split(l, "")

		for x, t := range tiles {
			if t == "#" {
				walls = append(walls, Point{x, y})
			}
			if t == "E" || t == "G" {
				units = append(units, newUnit(x, y, t))
			}

			maxx = int(math.Max(float64(maxx), float64(x)))
		}

		maxy = int(math.Max(float64(maxy), float64(y)))
	}

	return
}

func giveNanoWeaponToElves(units []Unit, power int) []Unit {
	for i, u := range units {
		if u.race == "E" {
			u.damage = power
			units[i] = u
		}
	}

	return units
}

func checkIfBattleIsFinished(units []Unit) bool {
	gCount := 0
	eCount := 0

	for _, u := range units {
		if !u.dead && u.race == "E" {
			eCount++
		}

		if !u.dead && u.race == "G" {
			gCount++
		}
	}

	return gCount == 0 || eCount == 0
}

func tryToHitSomeone(u Unit, units []Unit) {
	eidx, enemy := nearTarget(u, units)
	if !enemy.dead {
		u.hit(&enemy)
		units[eidx] = enemy
	}
}

func calculateOutcome(rounds int, units []Unit) int {
	hp := 0
	for _, u := range units {
		if !u.dead {
			hp += u.hp
		}
	}

	return hp * rounds
}

func sortInReadOrder(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		}
		return false
	})

	return points
}

func sortInNegativeReadOrder(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		return points[i].y > points[j].y
	})
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x > points[j].y
		}
		return false
	})

	return points
}

func buildTargets(cur Unit, units []Unit) (targets []Point) {
	targets = make([]Point, 0)
	for _, t := range units {
		if t.dead || t.race == cur.race {
			continue
		}

		zone := []Point{
			{t.x, t.y + 1},
			{t.x + 1, t.y},
			{t.x - 1, t.y},
			{t.x, t.y - 1},
		}

		targets = append(targets, zone...)
	}

	return
}

func nearTarget(u Unit, units []Unit) (idx int, enemy Unit) {
	selectedUnit := Unit{hp: 999, dead: true}
	selectedIdx := -1

	points := []Point{
		{u.x, u.y - 1},
		{u.x - 1, u.y},
		{u.x + 1, u.y},
		{u.x, u.y + 1},
	}

	for _, p := range points {
		idx, enemy = findByPoint(p, units)
		if !enemy.dead && enemy.race != u.race && selectedUnit.hp > enemy.hp {
			selectedUnit = enemy
			selectedIdx = idx
		}
	}

	return selectedIdx, selectedUnit
}

func visualize(walls []Point, units []Unit, maxx int, maxy int) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < maxy+1; y++ {
		for x := 0; x < maxx+1; x++ {
			p := Point{x, y}
			if pointInRange(p, walls) {
				fmt.Print("#")
				continue
			}

			_, u := findByPoint(p, units)
			if !u.dead {
				fmt.Print(u.race)
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()

	time.Sleep(200 * time.Millisecond)
}
