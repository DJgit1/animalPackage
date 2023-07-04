package main

import "fmt"

var x = 42

const y = 43

func main() {
	fmt.Printf("the value of x is %v and type of x is %T\n", x, x)
	fmt.Printf("the value of y is %v and type of y is %T\n", y, y)
	z := 44
	fmt.Printf("the value of z is %v and tzpe of z is %T\n", z, z)

}
