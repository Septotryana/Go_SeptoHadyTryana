package main

import "fmt"

func main() {
	const pi float64 = 3.14
	var T = 20.0
	var r = 4.0
	var rumus float64 = 2 * pi * r * (T + r)
	fmt.Print("Menghitung Luas :")
	fmt.Print(rumus)
}
