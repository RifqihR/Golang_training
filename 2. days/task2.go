package main

import (
	"fmt"
)

func main() {
	for k := 1; k <= 2; k++ {
		if k == 1 {
			for i := 1; i <= 4; i++ {
				fmt.Println("Niali i = ", i)
			}
		} else {
			for j := 1; j <= 10; j++ {
				if j == 5 {

					var ch = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
					for idx, val := range ch {
						fmt.Printf("character %U '%c' start at byte position %d \n", val, val, idx*2)
					}

				} else {
					fmt.Println("Niali j = ", j)
				}
			}
		}
	}
}
