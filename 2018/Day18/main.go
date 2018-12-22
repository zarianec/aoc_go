package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	tiles, maxx, maxy := parseInput(string(input))
	score := 0
	scores := make(map[int]int)
	hashes := make(map[string]int)
	scoreAfter10 := 0
	scoreAfter1e9 := 0

	for i := 0; i < 10000; i++ {
		tiles, score = tick(tiles, maxx, maxy)
		scores[i] = score
		//visualise(tiles, maxx, maxy)

		if i == 9 {
			scoreAfter10 = score
		}

		tilesString := tilesToString(tiles, maxx, maxy)
		hasher := md5.New()
		hasher.Write([]byte(tilesString))
		hash := hex.EncodeToString(hasher.Sum(nil))

		if idx, ok := hashes[hash]; ok {
			patternLength := i - idx
			targetStep := (1e9 - idx - 1) % patternLength
			scoreAfter1e9 = scores[idx+targetStep]
			break
		}
		hashes[hash] = i
	}

	fmt.Printf("Score after 10 seconds: %d\n", scoreAfter10)
	fmt.Printf("Score after 1e9 seconds: %d\n", scoreAfter1e9)

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func parseInput(input string) (tiles map[int]map[int]string, maxx int, maxy int) {
	tiles = make(map[int]map[int]string)
	lines := strings.Split(input, "\n")
	maxx = 0
	maxy = 0

	for y, l := range lines {
		for x, t := range strings.Split(l, "") {
			if tiles[x] == nil {
				tiles[x] = make(map[int]string)
			}
			tiles[x][y] = t

			maxx = int(math.Max(float64(maxx), float64(x)))
		}
		maxy = int(math.Max(float64(maxy), float64(y)))
	}

	return
}

func tick(tiles map[int]map[int]string, maxx int, maxy int) (newTiles map[int]map[int]string, score int) {
	totalTrees := 0
	totalLumbers := 0

	newTiles = make(map[int]map[int]string)
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			newType := determineNewType(x, y, tiles)

			if newTiles[x] == nil {
				newTiles[x] = make(map[int]string)
			}
			newTiles[x][y] = newType

			if newType == "|" {
				totalTrees++
			}
			if newType == "#" {
				totalLumbers++
			}
		}
	}

	score = totalTrees * totalLumbers

	return
}

func determineNewType(curx int, cury int, tiles map[int]map[int]string) string {
	curType := tiles[curx][cury]
	adjacents := make([]string, 0)

	adjacents = append(adjacents, tiles[curx-1][cury-1])
	adjacents = append(adjacents, tiles[curx][cury-1])
	adjacents = append(adjacents, tiles[curx+1][cury-1])
	adjacents = append(adjacents, tiles[curx-1][cury])
	adjacents = append(adjacents, tiles[curx+1][cury])
	adjacents = append(adjacents, tiles[curx-1][cury+1])
	adjacents = append(adjacents, tiles[curx][cury+1])
	adjacents = append(adjacents, tiles[curx+1][cury+1])

	newType := curType
	switch curType {
	case ".":
		newType = checkOpen(curType, adjacents)
		break
	case "|":
		newType = checkTree(curType, adjacents)
		break
	case "#":
		newType = checkLumber(curType, adjacents)
		break
	}

	return newType
}

func checkOpen(cur string, adjacents []string) string {
	treesCount := 0

	for _, t := range adjacents {
		if t == "|" {
			treesCount++
		}
	}

	if treesCount >= 3 {
		return "|"
	}

	return cur
}

func checkTree(cur string, adjacents []string) string {
	lumbersCount := 0

	for _, t := range adjacents {
		if t == "#" {
			lumbersCount++
		}
	}

	if lumbersCount >= 3 {
		return "#"
	}

	return cur
}

func checkLumber(cur string, adjacents []string) string {
	lumbersCount := 0
	treesCount := 0

	for _, t := range adjacents {
		if t == "#" {
			lumbersCount++
		}
		if t == "|" {
			treesCount++
		}
	}

	if lumbersCount > 0 && treesCount > 0 {
		return "#"
	}

	return "."
}

func tilesToString(tiles map[int]map[int]string, maxx int, maxy int) string {
	str := ""

	for y := 0; y < maxy+1; y++ {
		for x := 0; x < maxy+1; x++ {
			str += tiles[x][y]
		}
	}

	return str
}

func visualise(tiles map[int]map[int]string, maxx int, maxy int) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			fmt.Print(tiles[x][y])
		}
		fmt.Println()
	}
	fmt.Println()

	time.Sleep(100 * time.Millisecond)
}
