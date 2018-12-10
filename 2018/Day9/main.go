package main

import (
	"container/ring"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Highest score is: %d\n", playTheGame(430, 71588))
	fmt.Printf("Highest score is: %d\n", playTheGame(430, 71588*100))

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func playTheGame(players int, maxMarble int) int {
	scores := make([]int, players)
	circle := ring.New(1)
	circle.Value = 0

	curPlayer := 0
	for i := 1; i <= maxMarble; i++ {
		score := 0

		if i%23 == 0 {
			circle = circle.Move(-8)
			removed := circle.Unlink(1)
			circle = circle.Move(1)

			score += i + removed.Value.(int)
		} else {
			circle = circle.Move(1)
			s := ring.New(1)
			s.Value = i

			circle.Link(s)
			circle = circle.Move(1)
		}

		scores[curPlayer] += score

		curPlayer = (curPlayer + 1) % players
	}

	return maxScore(scores)
}

func maxScore(scores []int) int {
	max := 0
	for _, s := range scores {
		if max < s {
			max = s
		}
	}

	return int(max)
}
