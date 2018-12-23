package main

var opsList map[string]func([6]int, int, int, int) [6]int

func init() {
	opsList = make(map[string]func([6]int, int, int, int) [6]int, 16)
	opsList["addr"] = addr
	opsList["addi"] = addi

	opsList["mulr"] = mulr
	opsList["muli"] = muli

	opsList["banr"] = banr
	opsList["bani"] = bani

	opsList["borr"] = borr
	opsList["bori"] = bori

	opsList["setr"] = setr
	opsList["seti"] = seti

	opsList["gtir"] = gtir
	opsList["gtri"] = gtri
	opsList["gtrr"] = gtrr

	opsList["eqir"] = eqir
	opsList["eqri"] = eqri
	opsList["eqrr"] = eqrr
}

// Addition
func addr(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] + r[b]

	return r
}

func addi(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] + b

	return r
}

// Multiplication
func mulr(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] * r[b]

	return r
}

func muli(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] * b

	return r
}

// Bitwise AND
func banr(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] & r[b]

	return r
}

func bani(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] & b

	return r
}

// Bitwise OR
func borr(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] | r[b]

	return r
}

func bori(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a] | b

	return r
}

// Assignment
func setr(r [6]int, a int, b int, c int) [6]int {
	r[c] = r[a]

	return r
}

func seti(r [6]int, a int, b int, c int) [6]int {
	r[c] = a

	return r
}

// Greater-than testing
func gtir(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if a > r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func gtri(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if r[a] > b {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func gtrr(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if r[a] > r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

// Equality testing
func eqir(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if a == r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func eqri(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if r[a] == b {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}

func eqrr(r [6]int, a int, b int, c int) [6]int {
	bitSet := 0
	if r[a] == r[b] {
		bitSet = 1
	}
	r[c] = bitSet

	return r
}
