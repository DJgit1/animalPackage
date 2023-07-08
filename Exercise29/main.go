package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Print(i)
	}

	fmt.Println()

	for j := 0; j < 100; j++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("Iteration %v x = %d\t", j, x)
		fmt.Printf(" y = %d\t", y)

		if x < 4 && y < 4 {
			fmt.Println("x and y both are less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y both are less than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		} else if y != 5 {
			fmt.Println(" y is not 5")
		} else {
			fmt.Println("none of the previous cases were met")
		}
	}
}
