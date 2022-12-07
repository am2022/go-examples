package main

import (
	"fmt"
	"os"
)

func main() {
	var num int = 0
	var max int = 0

	for {
		fmt.Print("enter number(enter 0 for exit):")
		fmt.Scan(&num)

		if num == 0 {
			fmt.Print("max number:", max)
			os.Exit(0)
		} else if max < num {
			max = num
		}
	}
}
