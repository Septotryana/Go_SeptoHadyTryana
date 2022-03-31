package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var result string
	res := strings.Contains(a, b)
	if res == true {
		if len(a) < len(b) {
			result = a
		} else {
			result = b
		}
	}
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
}
