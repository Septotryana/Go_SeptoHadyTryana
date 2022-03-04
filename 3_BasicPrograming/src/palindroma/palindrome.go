package	main

import "fmt"

func palindrome(input string) bool {
	var result bool
	var e byte
	pembanding := ""
	for i:= len(input)-1; i >= 0;i--{
		e = input[i]
		pembanding = string(e)
	}
	if (input == pembanding){
		result = true
		//saya gatau kenapa masuknya ke else semua kak
	}else{
		result = false
	}
	return result
}
func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}