package main

type Unit struct {
	x      int
	y      int
	race   string
	hp     int
	damage int
	dead   bool
}

func (unit *Unit) hit(target *Unit) {
	target.hp -= unit.damage

	if 0 >= target.hp {
		target.dead = true
	}
}

func (unit *Unit) toPoint() Point {
	return Point{unit.x, unit.y}
}

// Unit factory
func newUnit(x int, y int, race string) Unit {
	return Unit{x, y, race, 200, 3, false}
}

func findByPoint(p Point, units []Unit) (int, Unit) {
	for idx, u := range units {
		if !u.dead && u.x == p.x && u.y == p.y {
			return idx, u
		}
	}

	return -1, Unit{dead: true}
}
