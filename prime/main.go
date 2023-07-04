package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

/*
Given this number sequence,
2, 3, 1, 5, 1, 7, 1, 1, 1, 11, 1, 13 ....
find the pattern and fulfill the function requirements.
The function takes one number input (i) and return array/vector as output. If there is no item to be returned, it will give empty array instead.
Contraints:
0 <= i <= 1000
Example:
Input:1; 	Output:[2]
Input:2; 	Output:[2,3]
Input:3; 	Output:[2,3,1]
Input:4; 	Output:[2,3,1,5]
Input:10;	Output:[2,3,1,5,1,7,1,1,1,11]
*/

func main() {
	// do not modify below here
	// this properly checks the function for you
	start := time.Now()

	fmt.Println(Solution(0))
	fmt.Println(Solution(1))
	fmt.Println(Solution(4))
	fmt.Println(Solution(15))
	fmt.Println(Solution(200))
	fmt.Println(Solution(1000))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func Solution(iValue int) (result []int) {
	//Write Your Code Here

	// Only accurate for input < 2⁶⁴
	for i := 1; i <= iValue; i++ {
		if big.NewInt(int64(i + 1)).ProbablyPrime(0) {
			result = append(result, i+1)
		} else {
			result = append(result, 1)
		}
	}
	return result
}
