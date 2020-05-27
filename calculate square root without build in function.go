package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var diff float64 = 0
	var z float64 = 1
	for i := 0; i < 1000; i++ {
		diff = z*z - x
		if diff > 0 {
			z = z - (diff * 0.1)
			if z < 0 {
				z = z * (-1)
			}
		} else {
			z = z + (diff * (-0.1))
			if z < 0 {
				z = z * (-1)
			}
		}
	}
	return z
}

func main() {
	fmt.Printf(" ___ ____ ____\n")
	fmt.Printf("| i | sqrt| diff|\n")
	fmt.Printf("|___|_____|_____|\n")
	for i := 0.0; i < 101; i++ {
		sqrt := Sqrt(i)
		diff := math.Sqrt(i) - sqrt
		abs_diff := diff
		if diff < 0 {
			abs_diff = diff * (-1)
		}
		fmt.Printf("|%3d|%5.3f|%5.3f|\n", int(i), sqrt, abs_diff)
		fmt.Printf("|___|_____|_____|\n")
	}
}
