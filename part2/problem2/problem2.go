package main

import "fmt"

func BottomUp(n int) int {
	// your code here
	a := 0
	b := 1
	if n < 2 {
		return n
	}
	for i := 2; i <= n; i++ {
		temp := a + b
		a = b
		b = temp
	}

	return b
}

func main() {
	fmt.Println(BottomUp(0))  //0
	fmt.Println(BottomUp(1))  //1
	fmt.Println(BottomUp(2))  //1
	fmt.Println(BottomUp(3))  //2
	fmt.Println(BottomUp(5))  //5
	fmt.Println(BottomUp(6))  //8
	fmt.Println(BottomUp(7))  //13
	fmt.Println(BottomUp(9))  //34
	fmt.Println(BottomUp(10)) //55
}
