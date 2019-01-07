package main

import (
	"container/heap"
	"fmt"
	"log"
	"time"
)

var depth = 6969
var target = Region{point: Point{9, 796}, regionType: 0}

var allowed map[terrain]equipment

func main() {
	start := time.Now()

	fmt.Printf("Risk in region: %d\n", calcRisk(target, drawMap(target, depth, 0)))
	fmt.Printf("Best time: %d\n", bestTime(target, drawMap(target, depth, 100)))

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func drawMap(target Region, depth int, buffer int) map[int]map[int]Region {
	regionMap := make(map[int]map[int]Region)

	if regionMap[target.point.x] == nil {
		regionMap[target.point.x] = make(map[int]Region)
	}
	regionMap[target.point.x][target.point.y] = target

	risk := 0
	for y := 0; y <= target.point.y+buffer; y++ {
		for x := 0; x <= target.point.x+buffer; x++ {
			r := Region{point: Point{x, y}}

			r.geoIndex = calculateIndex(r, regionMap)
			if target.point.x == x && target.point.y == y {
				r.geoIndex = 0
			}

			r.erosionLevel = (r.geoIndex + depth) % 20183
			r.regionType = terrain(r.erosionLevel % 3)

			if regionMap[x] == nil {
				regionMap[x] = make(map[int]Region)
			}
			regionMap[x][y] = r

			risk += int(r.regionType)
		}
	}

	return regionMap
}

func calcRisk(target Region, regionMap map[int]map[int]Region) (risk int) {
	risk = 0
	for y := 0; y <= target.point.y; y++ {
		for x := 0; x <= target.point.x; x++ {
			r := regionMap[x][y]
			risk += int(r.regionType)
		}
	}

	return
}

func bestTime(target Region, regionMap map[int]map[int]Region) int {
	costSoFar := make(map[Move]int)

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)

	initMove := Move{Point{0, 0}, torch}
	item := &Item{initMove, 0, 0}
	heap.Push(&queue, item)

	costSoFar[initMove] = 0

	for queue.Len() > 0 {
		cur := heap.Pop(&queue).(*Item)

		if cur.move.point == target.point && cur.move.equipment == torch {
			return costSoFar[cur.move]
		}

		adj := getAdjacent(cur.move.point, cur.move.equipment, regionMap)
		for _, next := range adj {
			changeGearCost := 1
			if next.equipment != cur.move.equipment {
				changeGearCost += 7
			}

			nm := Move{next.point, next.equipment}
			newCost := costSoFar[cur.move] + changeGearCost
			if cost, ok := costSoFar[nm]; !ok || newCost < cost {
				costSoFar[nm] = newCost
				heap.Push(&queue, &Item{nm, newCost, 0})
			}
		}
	}

	return -1
}

func getAdjacent(cur Point, equip equipment, regionMap map[int]map[int]Region) (next []Move) {
	next = make([]Move, 0)

	dir := []Direction{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for _, d := range dir {
		if n, ok := regionMap[cur.x+d.x][cur.y+d.y]; ok {
			if equip&allowed[n.regionType] != 0 {
				next = append(next, Move{n.point, equip})
				next = append(next, Move{n.point, equip ^ allowed[n.regionType]})
			}
		}
	}

	return
}

func calculateIndex(r Region, regionMap map[int]map[int]Region) int {
	if r.point.x == 0 && r.point.y == 0 {
		return 0
	}

	if r.point.x == 0 {
		return r.point.y * 48271
	}

	if r.point.y == 0 {
		return r.point.x * 16807
	}

	r1 := regionMap[r.point.x-1][r.point.y]
	r2 := regionMap[r.point.x][r.point.y-1]

	return r1.erosionLevel * r2.erosionLevel
}
