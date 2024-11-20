/*
testing random numbers for learning purpouses.
added zeroes before single digit number just for aesthetics
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for j := 0; j < 30; j++ {
		for i := 0; i < 30; i++ {
			liczba := rand.Intn(99)
			if liczba < 10 {
				fmt.Print("0")
			}
			fmt.Print(liczba, " ")
		}
		fmt.Println()
	}
}
