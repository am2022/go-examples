package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	var i int

	fmt.Print("enter number:")
	fmt.Scan(&num)

	for i = 2; i < num; i++ {
		if num%i == 0 {
			fmt.Print(num, " is not prime!")
			os.Exit(0)
		}
	}

	if num == i {
		fmt.Print(num, " is prime!")
	}
}
