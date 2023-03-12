package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat malam"

	charCount := make(map[string]int)

	for _, char := range kalimat {
		charStr := string(char)
		fmt.Println(charStr)
		charCount[charStr]++
	}

	fmt.Println(charCount)
}
