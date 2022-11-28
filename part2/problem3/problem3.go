package main

import (
	"fmt"
)

func Frog(jumps []int) int {
	// your code here
	var hasil int
	var hasil2 int
	var j int
	var panjangR int = len(jumps) // len panjangan array
	i := 0
	jumps1 := 0
	for i < panjangR {
		if jumps1 == 0 {
			hasil2 += Plus(jumps[j] - jumps[j+1])
			j++
		} else if j+2 <= panjangR-1 {
			hasil2 += Plus(jumps[j] - jumps[j+2])
			j += 2
		} else if j+1 <= panjangR-1 {
			hasil2 += Plus(jumps[j] - jumps[j+1])
			j++
		} else {
			j += 2
		}
		jumps1 = 2

		if i+2 <= panjangR-1 {
			hasil += Plus(jumps[i] - jumps[i+2])
		} else if i+1 <= panjangR-1 {
			hasil += Plus(jumps[i] - jumps[i+1])
		}
		i += 2

	}
	if hasil < hasil2 {
		return hasil
	}

	return hasil2
}
func Plus(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
