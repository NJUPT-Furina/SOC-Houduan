package main

import (
	"fmt"
)

func main() {
	num := 121
	if num == reverseInt(num) {
		fmt.Println(true)
	}

}
func reverseInt(n int) int {
	reversed := 0
	for {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
		break
	}

	return reversed
}
