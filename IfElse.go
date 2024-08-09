package main

import (
	"fmt"
	"strconv"
)

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 25; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 12 {
		fmt.Println(num, "has 1 digit")
	} else if len(strconv.Itoa(num)) > 3 {
		fmt.Println(num, "has multiple digits with greater than three")
	} else {
		fmt.Println("don't know what to do with", num)
	}
}
