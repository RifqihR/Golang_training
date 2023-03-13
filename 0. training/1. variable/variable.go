package main

import (
	"fmt"
)

func main() {
	var name string = "ariel"
	var names string
	names = "string"
	var age int = 20
	//uint 8-64 and int 8-64 --->uint is integer without the negative part

	//no type declaration
	var notype = "input can be anything"

	//short declaration
	othernames := "string"

	//multi declaration
	var student1, student2, student3 string = "satu", "dua", "tiga"
	var str, integer = "str", 20

	fmt.Println(name, names, othernames, age, notype)
	fmt.Println(student1, student2, student3)
	fmt.Println(str, integer)

	/*
		fmt.PrintF()--->digunakan tergantung flag yang diberikan
			--https://pkg.go.dev/fmt-->read this documentation

	*/
	fmt.Printf("%T", name) //---> %T calling datatype of the variable

	//operator still the same with other language
}
