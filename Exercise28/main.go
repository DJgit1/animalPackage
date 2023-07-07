package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)

	/*
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
	*/

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y both are less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y both are less than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println(" y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}
