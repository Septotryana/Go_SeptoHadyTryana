package main

import "fmt"

func primeNumber(number int) bool{
	var result bool
	var i int
	var test int = 0
	if(number == 0 || number == 1){
	test = 1
}
	for i = 2; i <= number/2;i++{
		if number % i == 0{
			test = 1
			break;
		}
	}
	if test == 0 {
		result = true
	}else{
		result = false
	}
	return result
}
func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
