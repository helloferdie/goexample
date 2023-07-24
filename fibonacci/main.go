package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cache = map[int64]int64{}

// Fibonacci number sequence:
// in: 	0, 1, 2, 3, 4, 5, 6
// out:	0, 1, 1, 2, 3, 5, 8

func main() {
	fmt.Println("Sequence:")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	seq, _ := strconv.ParseInt(s, 10, 64)

	fmt.Printf("Fibonacci number is %v\n", getFibonacci(seq))
}

func getFibonacci(i int64) int64 {
	if i < 2 {
		return i
	}

	if r, ok := cache[i]; ok {
		return r
	}

	r := getFibonacci(i-1) + getFibonacci(i-2)
	cache[i] = r
	return r
}
