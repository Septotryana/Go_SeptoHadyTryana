package main

import "fmt"

func pangkat(base , pangkat int) int {
	result := 1
	if pangkat != 0{
		for i:= 1; i <= pangkat; i++{
			result *= base
		}
	}
	return result
}
func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}