package main

import (
	"fmt"
	"sort"
	"strconv"
)

func DragonOfLoowater(dragonHead, knightHeight []int) string {
	// your code here
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	dragon := len(dragonHead)
	night := len(knightHeight)
	var i, j, hasil int
	var res string

	for j = 0; j < night; j++ {
		if knightHeight[j] >= dragonHead[i] {
			hasil += knightHeight[j]
			i++
		}
		if i == dragon {
			break
		}
	}
	if i < dragon {
		res += "Knight Fall"
	} else {
		res += strconv.Itoa(hasil)
	}
	return res
}

func main() {
	fmt.Println(DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}))    //11
	fmt.Println(DragonOfLoowater([]int{5, 10}, []int{5}))         // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})) // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})) // 10
}
