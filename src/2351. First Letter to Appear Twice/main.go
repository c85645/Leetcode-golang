package main

import "fmt"

func repeatedCharacter(s string) byte {
	repeat := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if repeat[char] {
			return char
		}
		repeat[char] = true
	}
	return 0
}

func main() {
	case1 := "abccbaacz"
	case2 := "abcdd"
	result1 := repeatedCharacter(case1)
	result2 := repeatedCharacter(case2)
	fmt.Printf("第一個重複的字母是: %c (byte: %d)\n", result1, result1)
	fmt.Printf("第一個重複的字母是: %c (byte: %d)\n", result2, result2)
}
