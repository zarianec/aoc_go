package main

var opsList [16]func([4]int, int, int, int) [4]int

func init() {
	// Map operations
	opsList = [16]func([4]int, int, int, int) [4]int{
		addr, addi,
		mulr, muli,
		banr, bani,
		borr, bori,
		setr, seti,
		gtir, gtri, gtrr,
		eqir, eqri, eqrr,
	}
}

// Addition
func addr(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] + r[b]

	return r
}

func addi(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] + b

	return r
}

// Multiplication
func mulr(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] * r[b]

	return r
}

func muli(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] * b

	return r
}

// Bitwise AND
func banr(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] & r[b]

	return r
}

func bani(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] & b

	return r
}

// Bitwise OR
func borr(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] | r[b]

	return r
}

func bori(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a] | b

	return r
}

// Assignment
func setr(r [4]int, a int, b int, c int) [4]int {
	r[c] = r[a]

	return r
}

func seti(r [4]int, a int, b int, c int) [4]int {
	r[c] = a

	return r
}

// Greater-than testing
func gtir(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if a > r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func gtri(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if r[a] > b {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func gtrr(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if r[a] > r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

// Equality testing
func eqir(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if a == r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func eqri(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if r[a] == b {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func eqrr(r [4]int, a int, b int, c int) [4]int {
	bitSet := 0
	if r[a] == r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}
