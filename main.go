package main

import (
	"fmt"
	"math/rand"
)

func randomazzo(priorities_sum int) (priorities_result int) {

	random := rand.Intn(priorities_sum)

	priorities_result = random - priorities_sum

	return priorities_result
}

func main() {

	// input the strings separately.
	//	TODO import data from the file
	var jobs []string = []string{}
	for i := 0; ; i++ {
		fmt.Scan(&jobs[i])
		if jobs[i] == "end" {
			break
		}
	}
	fmt.Println("set priorities! ")
	var priorities_sum int
	var priority int
	for _, job := range jobs {

		fmt.Print(job, " : ")
		fmt.Scan(&priority)
		fmt.Println()
		priorities_sum += priority

	}

}
