package main

import "fmt"

func cetakTabelPerkalian(number int){
	for i := 1; i <= number; i++ {
		for j:= 1; j <= number; j++{
			fmt.Print(i * j)
		}
		fmt.Print("\n\n")
	}
}
func main () {
	cetakTabelPerkalian(9)
}