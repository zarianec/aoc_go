package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type Operation struct {
	opcode string
	a      int
	b      int
	c      int
}

func main() {
	start := time.Now()
	input, _ := ioutil.ReadFile("input.txt")

	initPointer, ops := parseInput(input)

	runProgram([6]int{0, 0, 0, 0, 0, 0}, initPointer, ops)
	runProgram([6]int{1, 0, 0, 0, 0, 0}, initPointer, ops)

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func runProgram(reg [6]int, ip int, ops map[int]Operation) {
	// registry 4 is hardcoded here and it's specific for my input
	var targetRegistry = 4

	curPointer := reg[ip]
	for curPointer < len(ops) {
		if reg[ip] == 1 {
			sum := 0;
			for i := 1; i <= reg[targetRegistry]; i++ {
				if reg[targetRegistry]%i == 0 {
					sum += i
				}
			}
			fmt.Printf("Registry 0 value: %d\n", sum)
			break
		}

		curPointer = reg[ip]
		op := ops[curPointer]
		f := opsList[op.opcode]
		reg = f(reg, op.a, op.b, op.c)
		reg[ip]++
	}
}

func parseInput(input []byte) (pointer int, ops map[int]Operation) {
	lines := bytes.Split(input, []byte("\n"))

	pointer, _ = strconv.Atoi(string(lines[0][4]))
	ops = make(map[int]Operation)

	for i, op := range lines[1:] {
		parts := bytes.Split(op, []byte(" "))
		a, _ := strconv.Atoi(string(parts[1]))
		b, _ := strconv.Atoi(string(parts[2]))
		c, _ := strconv.Atoi(string(parts[3]))

		ops[i] = Operation{string(parts[0]), a, b, c}
	}

	return
}
