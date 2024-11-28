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
	var zadania []string = []string{}
	for n := 0; ; n++ {
		fmt.Scan(&zadania[n])
		if zadania[n] == "end" {
			break
		}
	}
}
