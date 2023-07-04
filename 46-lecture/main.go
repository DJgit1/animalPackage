package main

import (
	"fmt"

	"github.com/DJgit1/puppy"
)

func main() {
	puppy.From11()
	puppy.From12()
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

}
