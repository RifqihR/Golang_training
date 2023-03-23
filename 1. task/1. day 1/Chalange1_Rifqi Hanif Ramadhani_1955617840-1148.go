package main

import (
	"encoding/hex"
	"fmt"
)

var i = 21
var j = true
var k float64 = 123.456

func main() {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j)

	fmt.Printf("%b\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)

	tmp := hex.EncodeToString([]byte("f"))
	out, err := hex.DecodeString(tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	tmp = hex.EncodeToString([]byte("F"))
	out, err = hex.DecodeString(tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	//unicode of character
	rsh := "Я"
	ck := []rune(rsh)
	fmt.Printf("%U\n", ck[0])
	/*
		for _, runeValue := range checking {
		fmt.Printf("%T\n", runeValue)
		}
	*/

	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)
}

//rifqi hanif ramadhani 1955617840-1148
