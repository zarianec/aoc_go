package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Sample struct {
	before [4]int
	input  [4]int
	after  [4]int
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	samples, instructions := parseInput(input)

	// Part 1
	tc := 0
	m := make(map[int]map[int]int, 0)
	for _, s := range samples {
		oc := 0
		for i, f := range opsList {
			r := [4]int{s.before[0], s.before[1], s.before[2], s.before[3]}
			ra := [4]int{s.after[0], s.after[1], s.after[2], s.after[3]}

			r = f(r, s.input[1], s.input[2], s.input[3])
			if ra == r {
				oc++

				if m[i] == nil {
					m[i] = map[int]int{}
				}
				m[i][s.input[0]] += 1
			}
		}

		if oc >= 3 {
			tc++
		}
	}
	fmt.Printf("Amount of samples with 3 or more opcodes: %d\n", tc)

	// Part 2
	mapped := make(map[int]int)
	for len(m) > 0 {
		for opcid, _ := range opsList {
			for opc, _ := range m[opcid] {
				if _, ok := m[opcid][opc]; !ok {
					continue
				}

				if _, ok := mapped[opc]; ok {
					delete(m[opcid], opc)
					continue
				}

				if len(m[opcid]) == 1 {
					mapped[opc] = opcid
					delete(m[opcid], opc)
				}
			}

			if len(m[opcid]) == 0 {
				delete(m, opcid)
			}
		}
	}

	// Reset register
	r := [4]int{0, 0, 0, 0}
	for _, instr := range instructions {
		op := mapped[instr[0]]
		f := opsList[op]
		r = f(r, instr[1], instr[2], instr[3])
	}

	fmt.Printf("Register 0 value is: %d\n", r[0])

	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func parseInput(input []byte) (samples []Sample, instructions [][4]int) {
	parts := strings.Split(string(input), "\n\n\n\n")

	// Parse part 1
	samples = make([]Sample, 0)
	r := strings.NewReplacer("Before: ", "", "After:  ", "", "[", "", "]", "", "\n", " ", ",", "")

	sg := strings.Split(parts[0], "\n\n")
	for _, smpl := range sg {
		rs := r.Replace(smpl)
		rd := [12]int{}
		for i, s := range strings.Split(rs, " ") {
			rd[i], _ = strconv.Atoi(s)
		}

		s := Sample{
			[4]int{rd[0], rd[1], rd[2], rd[3]},
			[4]int{rd[4], rd[5], rd[6], rd[7]},
			[4]int{rd[8], rd[9], rd[10], rd[11]},
		}
		samples = append(samples, s)
	}

	// Parse part 2
	instructions = make([][4]int, 0)
	insList := strings.Split(parts[1], "\n")
	for _, ins := range insList {
		codes := strings.Split(ins, " ")
		i := [4]int{}
		i[0], _ = strconv.Atoi(codes[0])
		i[1], _ = strconv.Atoi(codes[1])
		i[2], _ = strconv.Atoi(codes[2])
		i[3], _ = strconv.Atoi(codes[3])

		instructions = append(instructions, i)
	}

	return
}
