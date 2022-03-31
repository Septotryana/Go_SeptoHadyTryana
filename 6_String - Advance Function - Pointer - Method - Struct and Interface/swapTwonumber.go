package main

import "fmt"

func swap(a, b *int) {
	result := 0
	result = *b
	b = a
	a = &result

}
func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}
