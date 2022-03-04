package main

import "fmt"

func playWithAsterisk(n int) {
	fmt.Println("/*")
	for i := 0; i != n; i++ {
		fmt.Println("\t*")
		for j := 0; j <= i; j++ {
			fmt.Print("\t*")
		}
	}
	fmt.Println("*/")
}
func main() {
	playWithAsterisk(5)
}