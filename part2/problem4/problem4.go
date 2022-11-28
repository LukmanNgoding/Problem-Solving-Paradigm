package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RomanNumerals(value int) string {
	//your code here
	numberN := 3999
	if value > numberN {
		return strconv.Itoa(value)
	}

	conversi := []struct {
		hasil int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var Roman strings.Builder
	for _, conversi := range conversi {
		for value >= conversi.hasil {
			Roman.WriteString(conversi.digit)
			value -= conversi.hasil
		}
	}
	return Roman.String()
}

func main() {
	fmt.Println(RomanNumerals(6))    //VI
	fmt.Println(RomanNumerals(9))    //IX
	fmt.Println(RomanNumerals(23))   //XXIII
	fmt.Println(RomanNumerals(2021)) //MMXXI
	fmt.Println(RomanNumerals(1646)) //MDCXLVI
}
