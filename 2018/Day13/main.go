package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Point struct {
	x         int
	y         int
	pointType int
	sign      string
}

type Cart struct {
	x    int
	y    int
	dir  int
	turn int
}

func (c *Cart) makeTurn(turnType int) {
	if turnType == 0 {
		return
	}

	if turnType == 3 {
		c.dir += c.turn
		c.dir = (4 + c.dir) % 4
		c.turn++

		if c.turn > 1 {
			c.turn = -1
		}

		return
	}

	hash := c.dir + turnType
	if hash%2 == 0 {
		c.dir = (c.dir + 1) % 4
	} else {
		c.dir = (4 + (c.dir - 1)) % 4
	}
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	track, carts, maxX, maxY := buildTrack(string(input))

	// I'm to lazy to implement deep copy of maps in Go
	// so if you decided to run this (who am i kidding - no one is interested in this ;))
	// you need to uncomment needed part of code manually

	// Part 1 start
	//crashed := false
	//crashPoint := [2]int{}
	//for crashed != true {
	//	// visualise(track, carts, maxX, maxY)
	//	crashed, crashPoint, carts = tickTillCrash(track, carts, maxX, maxY)
	//}
	//
	//fmt.Printf("First crash position: %d,%d\n", crashPoint[0], crashPoint[1])
	// Part 1 end

	// Part 2 start
	totalCarts := len(carts)
	lastCart := Cart{}
	for totalCarts > 1 {
		//visualise(track, carts, maxX, maxY)
		carts, totalCarts, lastCart = tickTillLast(track, carts, maxX, maxY)
	}
	// Part 2 end

	fmt.Printf("Last cart position: %d,%d\n", lastCart.x, lastCart.y)

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

func tickTillCrash(track map[int]map[int]Point, carts map[int]map[int]Cart, maxX int, maxY int) (crashed bool, crashPoint [2]int, newCarts map[int]map[int]Cart) {
	newCarts = make(map[int]map[int]Cart)

	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if !cartIsset(x, y, carts) {
				continue
			}

			c := carts[x][y]
			np := getNextPoint(track, c)
			c.x, c.y = np.x, np.y
			c.makeTurn(np.pointType)

			if cartIsset(np.x, np.y, carts) {
				crashPoint = [2]int{np.x, np.y}
				crashed = true
			}

			if cartIsset(np.x, np.y, newCarts) {
				crashPoint = [2]int{np.x, np.y}
				crashed = true
			}

			if crashed {
				return
			}

			if newCarts[c.x] == nil {
				newCarts[c.x] = make(map[int]Cart)
			}
			newCarts[c.x][c.y] = c
			delete(carts[x], y)
		}
	}

	return
}

func tickTillLast(track map[int]map[int]Point, carts map[int]map[int]Cart, maxX int, maxY int) (newCarts map[int]map[int]Cart, totalCarts int, lastCart Cart) {
	newCarts = make(map[int]map[int]Cart)
	totalCarts = 0
	lastCart = Cart{}

	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if !cartIsset(x, y, carts) {
				continue
			}
			crashed := false

			c := carts[x][y]
			np := getNextPoint(track, c)
			c.x, c.y = np.x, np.y
			c.makeTurn(np.pointType)

			if cartIsset(np.x, np.y, carts) {
				delete(carts[np.x], np.y)
				crashed = true
			}

			if cartIsset(np.x, np.y, newCarts) {
				delete(newCarts[np.x], np.y)
				crashed = true
			}

			if crashed {
				totalCarts--
				continue
			}

			if newCarts[c.x] == nil {
				newCarts[c.x] = make(map[int]Cart)
			}
			newCarts[c.x][c.y] = c
			lastCart = c

			delete(carts[x], y)
			totalCarts++
		}
	}

	return
}

func visualise(track map[int]map[int]Point, carts map[int]map[int]Cart, maxX int, maxY int) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			p := track[x][y]

			if p.sign == "" {
				fmt.Print(" ")
				continue
			}

			if _, ok := (carts)[x][y]; ok {
				c := carts[x][y]

				fmt.Print("\033[1;31m" + getCartSign(c.dir) + "\033[0m")
				continue
			}

			if p.sign == "<" || p.sign == ">" || p.sign == "v" || p.sign == "^" {
				fmt.Print(".")
			} else {
				fmt.Print(p.sign)
			}
		}

		fmt.Println()
	}

	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}

func buildTrack(input string) (track map[int]map[int]Point, carts map[int]map[int]Cart, maxX int, maxY int) {
	lines := strings.Split(input, "\n")
	track = make(map[int]map[int]Point)
	carts = make(map[int]map[int]Cart)

	for y, l := range lines {
		c := strings.Split(l, "")

		for x, p := range c {
			if track[x] == nil {
				track[x] = make(map[int]Point)
			}

			track[x][y] = Point{x, y, getPointType(p), p}

			if isCart(p) {
				if carts[x] == nil {
					carts[x] = make(map[int]Cart)
				}
				carts[x][y] = buildCart(x, y, p)
			}

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	return
}

func getNextPoint(track map[int]map[int]Point, cart Cart) Point {
	np := Point{}

	if cart.dir == 0 {
		np = track[cart.x+1][cart.y]
	} else if cart.dir == 1 {
		np = track[cart.x][cart.y+1]
	} else if cart.dir == 2 {
		np = track[cart.x-1][cart.y]
	} else if cart.dir == 3 {
		np = track[cart.x][cart.y-1]
	}

	return np
}

func isCart(p string) bool {
	if p == "<" || p == ">" || p == "v" || p == "^" {
		return true
	}

	return false
}

func buildCart(x int, y int, c string) Cart {
	dir := 0

	switch c {
	case ">":
		dir = 0
	case "v":
		dir = 1
	case "<":
		dir = 2
	case "^":
		dir = 3
	}

	return Cart{x, y, dir, -1}
}

func getPointType(point string) int {
	t := 0

	switch point {
	case "/":
		t = 1
	case "\\":
		t = 2
	case "+":
		t = 3
	}

	return t
}

func getCartSign(dir int) string {
	typeToCartMap := map[int]string{
		0: ">",
		1: "v",
		2: "<",
		3: "^",
	}

	return typeToCartMap[dir]
}

func cartIsset(x int, y int, carts map[int]map[int]Cart) bool {
	if _, ok := (carts)[x][y]; ok {
		return true
	}

	return false
}
