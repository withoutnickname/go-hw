package main

import (
	"errors"
	"fmt"
)

func convertString(str string) (string, error) {
	var result string
	count := 1

	for i := 0; i < len(str); i++ {
		char := str[i]

		if !isAllowedSymbol(char) {
			return "", errors.New("non-Latin and non-camel symbols are not allowed")
		}
		if i+1 < len(str) && char == str[i+1] {
			count++
		} else {
			result += string(char) + fmt.Sprint(count)
			count = 1
		}
	}
	return result, nil
}

func isAllowedSymbol(char byte) bool {
	return (char >= 'a' && char <= 'z')
}

func main() {
	str := "aaaaabbcccdddddddddddddddeeeeffffffg"
	converted, err := convertString(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Converted:", converted)
}
