package main

import "fmt"

func reverseString(str string) string {
	byte_str := []rune(str) // converts string into array of characters
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i] // replacing 1st with last character & so on
	}
	return string(byte_str)
}

func main() {
	fmt.Println(reverseString("himalaya"))
	fmt.Println(reverseString("taj"))
	fmt.Println(reverseString("tropical"))
}
