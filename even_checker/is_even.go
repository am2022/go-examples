package main

import "fmt"

func main(){
	var num int;

	fmt.Print("enter number:");
	fmt.Scan(&num);

	if num % 2 == 0{
		fmt.Println(num, "is even!");
	}else{
		fmt.Println(num, "is odd!");
	}
}