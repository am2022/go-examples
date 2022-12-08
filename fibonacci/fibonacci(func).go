package main

import "fmt"

func fib(num int) {
	var f1 int = 1
	var f2 int = 1
	var sum int

	for i := 1; i <= num; i++ {
		if i == 1 {
			fmt.Println(f1)
			fmt.Println(f2)
		} else {
			sum = f1 + f2
			fmt.Println(sum)
			f1 = f2
			f2 = sum
		}
	}
}

func main() {
	var num int

	fmt.Print("enter number:")
	fmt.Scan(&num)

	fib(num)
}
