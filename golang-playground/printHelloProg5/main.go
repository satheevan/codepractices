package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Print("Enter the value:")
	fmt.Scanf("%d", &a)

	for i := 1; i <= a; i++ {
		fmt.Println("Hello")
	}

}
