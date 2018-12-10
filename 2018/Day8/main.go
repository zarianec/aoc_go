package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

var cursor int

type Node struct {
	children []*Node
	meta     []int
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	tmp := strings.Split(string(input), " ")

	var license []int
	for _, val := range tmp {
		num, _ := strconv.Atoi(val)
		license = append(license, num)
	}

	cursor = 0
	tree := buildTree(license)

	fmt.Printf("Sum of meta: %d\n", sumMeta(tree))
	fmt.Printf("Root node score: %d\n", calcScore(tree))

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func buildTree(license []int) (node *Node) {
	cursor += 2

	nc := int(license[cursor-2])
	nm := int(license[cursor-1])

	node = &Node{}

	for i := 0; i < nc; i++ {
		node.children = append(node.children, buildTree(license))
	}

	node.meta = license[cursor : cursor+nm]
	cursor += nm

	return node
}

func sumMeta(node *Node) (sum int) {
	for i := 0; i < len(node.children); i++ {
		sum += sumMeta(node.children[i])
	}

	for _, v := range node.meta {
		sum += v
	}

	return sum
}

func calcScore(node *Node) (sum int) {

	if len(node.children) == 0 {
		return sumMeta(node)
	}

	for _, m := range node.meta {
		if m > 0 && m <= len(node.children) {
			sum += calcScore(node.children[m-1])
		}
	}

	return sum
}
