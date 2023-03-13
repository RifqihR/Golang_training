package main

import (
	"fmt"
)

func main() {
	result := make(map[string]int)

	var data = "selamat malam"
	for _, char := range data {
		fmt.Printf("%c\n", char)
		result[string(char)]++
	}
	fmt.Println(result)
}
