package main

import "fmt"

func swap(a, b *int) {
<<<<<<< HEAD
	result := 0
	result = *b
	b = a
	a = &result
=======
>>>>>>> 3c008fd81e72e6798fe33bb65e4c307d898b07c8

}
func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}
