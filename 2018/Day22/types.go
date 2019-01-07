package main

const (
	rocky  terrain = 0
	wet    terrain = 1
	narrow terrain = 2
)

const (
	torch   equipment = 1
	gear    equipment = 2
	neither equipment = 4
)

type terrain int
type equipment int

type Point struct {
	x, y int
}

type Region struct {
	point        Point
	regionType   terrain
	erosionLevel int
	geoIndex     int
}

type Move struct {
	point     Point
	equipment equipment
}

type Direction struct {
	x, y int
}

func init() {
	allowed = make(map[terrain]equipment)
	allowed[rocky] = gear | torch
	allowed[wet] = gear | neither
	allowed[narrow] = torch | neither
}
