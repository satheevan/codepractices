package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Print("Enter the number :")
	var str string = ""
	fmt.Scanf("%s", &str)
	fmt.Println(reflect.TypeOf(str))
	intVar, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("String to Integer :", intVar, reflect.TypeOf(intVar))
}
