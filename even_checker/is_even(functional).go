package main

import "fmt"

func is_even(num int) int {
	if num%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	var num int

	fmt.Print("enter number:")
	fmt.Scan(&num)

	i := is_even(num)

	if i == 1 {
		fmt.Println("number:", num, "is even")
	} else {
		fmt.Println("number:", num, "is odd")
	}
}
