package main

import (
	"fmt"
)

func main() {
	greet("rifqi", 12)
	var closure = func(number int) int { //closure function
		var result int
		result = number * 2
		return result
	}

	var timestwo = closure(2)
	fmt.Println(timestwo)
}

func greet(name string, age int) { //parameter can be writen name, address string if both has same type
	fmt.Printf("im %s and im %d years old\n", name, age)
}

func birth(name string, age int) (string, int) { //string on the outer bracket is return type
	return name, age
}

func calculate(top, side int) (area int, circumference int) { //predifined return
	area = top * side
	circumference = 2 * (top + side)
	return
}

//varadic function --> func name(name, ...string)-->
