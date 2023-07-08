package main

import "fmt"

func main() {

	// x := 0

	// for x < 5 {
	// 	fmt.Printf("Iteration no. of outer loop %v\n", x)
	// 	x++
	// 	y := 0
	// 	for y < 5 {
	// 		fmt.Printf("Iteration no. of inner loop %v\t", y)
	// 		y++
	// 	}
	// 	fmt.Println()
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
		fmt.Println()
	}
}
