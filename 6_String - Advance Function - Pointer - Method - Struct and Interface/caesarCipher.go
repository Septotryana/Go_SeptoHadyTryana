package main

import (
	"fmt"
	"strings"
)
func caesar(offset int, input string) string {
	var result string
	
	for i := 0; i < len(input); i++ {
		if isUpper(input[i]){
			result = string(char(int(input[i]+ offset-65)%26 + 65))
			
		}else{
			result = string(char(int(input[i]+ offset-97)%26 + 97))
		}
	}
	return result
}
func main{
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
}