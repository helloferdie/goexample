package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
The function below wants to print all the names on the nameList
Modify the solution() function so that this code gets all the name within 1 second
Input: length of []nameList
Rule:
-You must get the names from getName() function. Any other way of inputting names to result is considered invalid.
-No modification to any of the existing functions are allowed, with the exception of solution().
*/

func solution(length int) []string {
	//YOU CAN ONLY MODIY OR WRITE CODE HERE
	//var result []string

	var wg sync.WaitGroup
	result := make([]string, length)
	for i := 0; i < length; i++ {
		go func(result []string, i int) {
			result[i] = getName(i)
			defer wg.Done()
		}(result, i)

	}
	wg.Wait()
	fmt.Println(result)
	return result
}

func main() {
	// do not modify below here
	// this properly checks the function for you
	start := time.Now()

	fmt.Println(solution(5))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func getName(n int) string {
	// do not modify below here
	// this properly checks the function for you
	var nameList = []string{"Budi", "Putra", "Putri", "Dwi", "Sarah"}
	time.Sleep(750 * time.Millisecond)
	return nameList[n]
}
