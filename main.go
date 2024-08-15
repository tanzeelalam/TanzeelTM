package main

import (
	"fmt"

	"github.com/tanzeelalam/puppy"
)

func main() {
	fmt.Println("Hello 😁")
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
}
