package main

type SquareGrid struct {
	width     int
	height    int
	obstacles []Point
}

func (g *SquareGrid) inBounds(p Point) bool {
	return p.x >= 0 && p.x < g.width && p.y >= 0 && p.y < g.height
}

func (g *SquareGrid) isPassable(p Point) bool {
	return !pointInRange(p, g.obstacles)
}

func (g *SquareGrid) neighbours(p Point) []Point {
	neighbours := []Point{
		{p.x, p.y - 1},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x + 1, p.y},
	}

	// Filter in bounds
	neighbours = filterPoints(neighbours, func(p Point) bool {
		return g.inBounds(p)
	})

	// Filter obstacles
	neighbours = filterPoints(neighbours, func(p Point) bool {
		return g.isPassable(p)
	})

	return neighbours
}

func findClosestTarget(start Unit, targets []Point, units []Unit, walls []Point, maxx int, maxy int) ([]Point, map[Point]Point) {
	queue := make([]Node, 0)
	queue = append(queue, Node{start.toPoint(), 0})

	cameFrom := make(map[Point]Point, 0)
	cameFrom[queue[0].p] = Point{}

	closest := make([]Point, 0)
	closestDist := -1

	grid := SquareGrid{maxx, maxy, buildObstaclesFromUnitsAndWalls(start, units, walls)}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if closestDist != -1 && cur.d > closestDist {
			return closest, cameFrom
		}

		if pointInRange(cur.p, getVisitedPoints(cameFrom)) {
			continue
		}

		if pointInRange(cur.p, targets) && !pointInRange(cur.p, closest) {
			closestDist = cur.d
			closest = append(closest, cur.p)
		}

		nbs := grid.neighbours(cur.p)
		nbs = sortInNegativeReadOrder(nbs)
		for _, next := range nbs {
			if !pointInRange(next, getVisitedPoints(cameFrom)) {
				queue = append(queue, Node{next, cur.d + 1})
				cameFrom[next] = cur.p
			}
		}
	}

	return closest, cameFrom
}

func buildObstaclesFromUnitsAndWalls(curUnit Unit, units []Unit, walls []Point) (obstacles []Point) {
	obstacles = walls // All the walls are obstacles

	for _, u := range units {
		if !u.dead {
			obstacles = append(obstacles, u.toPoint())
		}
	}

	return
}

func buildPath(cameFrom map[Point]Point, from Point, to Point) []Point {
	cur := to
	path := make([]Point, 0)

	for cur != from {
		path = append(path, cur)
		cur = cameFrom[cur]
	}
	path = reversePoints(path)

	return path
}

func reversePoints(points []Point) []Point {
	for left, right := 0, len(points)-1; left < right; left, right = left+1, right-1 {
		points[left], points[right] = points[right], points[left]
	}

	return points
}

func pointInRange(p Point, points []Point) bool {
	for _, p2 := range points {
		if p2 == p {
			return true
		}
	}
	return false
}

func getVisitedPoints(visited map[Point]Point) []Point {
	l := make([]Point, 0)

	for _, p := range visited {
		l = append(l, p)
	}

	return l
}

func filterPoints(points []Point, f func(p Point) bool) []Point {
	filtered := points[:0]
	for _, x := range points {
		if f(x) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}
