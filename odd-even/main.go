package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	
	arr := initArr()

	for _, value := range arr {
		if value%2 == 0 {
			fmt.Println(value," is even")
		} else {
			fmt.Println(value," is odd")
		}
	}
}

func initArr() []int {
	var arr []int
	for i := 0; i <= 10; i++ {
		arr = append(arr, i)
	}
	return arr
}