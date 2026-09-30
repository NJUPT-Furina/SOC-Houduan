package main

import (
	"fmt"
)

func main() {
	num := 121
	reversenum, err := reverseInt(num)
	if err != nil {
		fmt.Println("输入不正确！")
		return
	}
	if num == reversenum {
		fmt.Println("True")
	}
	fmt.Println("False")

}
func reverseInt(n int) (int, error) {
	reversed := 0
	for {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
		break
	}
	return reversed, nil
}
