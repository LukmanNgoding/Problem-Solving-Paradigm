package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) interface{} {
	// your code here

	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			for z := 0; z <= 100; z++ {
				z_res := int(math.Pow(float64(x), 2)) + int(math.Pow(float64(y), 2)) + int(math.Pow(float64(z), 2))
				if x+y+z == a && x*y*z == b && z_res == c {
					return []int{x, y, z}
				} else {
				}
			}
		}
	}
	return ("no solution")

}

func main() {
	fmt.Println(SimpleEquations(1, 2, 3))  //no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // [1,2,3]
}
