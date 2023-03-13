package main

import (
	"fmt"
)

func main() {
	var year int = 2023

	//if-else
	if age := year - 1999; age < 25 {
		fmt.Println(age)
	} else {
		fmt.Println("too old")
	}

	//switch-case
	score := 8

	switch score {
	case 8: //or condition like score==8
		fmt.Println("lolos")
	case 7:
		fmt.Println("lolos")
		fallthrough //meneruskan pengecekan ke fallthrough berikutnya
	case 6:
		fmt.Println("tidak lolos")
	default:
		fmt.Println("need to learn harder next time")
	}

	//for loop

	//3 type of loow
	// for i:= xx; i<x;i++{}
	// for i<x {}
	// for {i++ if i==3{break}}

	//for {continue}-->> used for skipping current loop to next loop (example loop when i=1, continue to i=2 without running rest of the code below continue)
	n := 0
marker: //this is label
	for n < 10 {
		for j := 0; j < 3; j++ {
			if n == 4 {
				break marker //break only exit the current loop, need label to exit the whole nested loop into the labeled line
			}
			fmt.Println(j)
		}
	}

	//array
	var studentscore [4]int //
	studentscore = [4]int{1, 2, 3, 4}
	var studname = [2][2]string{{"jonathan", "mahmud"}, {"bahrudin", "muinem"}}

	fmt.Printf("%#v\n", studentscore)
	fmt.Printf("%#v\n", studname)

	//looping through array can use normal loop or
	for index, value := range studentscore {
		fmt.Println(index, value)
	}

	//slice == array. but its a refference type. changing one element will change the other too
	var fruit = []string{"apple"}
	_ = fruit

	var price = make([]int, 3) //creating slice price with 3 element

	price[0] = 1
	price[1] = 2
	price[2] = 3

	//or

	var price2 = make([]int, 3)
	price2 = append(price2, 4, 5, 6)

	//adding other slice into another
	price = append(price, price2...) //"..." mean we take all the element

	//slicing slice. temp will share memory with price. changinge one element in one slice will also change the other element coresponding with the changed element
	var temp = price[2:3]
	fmt.Printf("%#v\n", temp)

	//cap(array.slice) will return the max size of the array/slice

	//map key:value. dictionary in python/javascript
	var person map[string]string //key: string, value: string
	person = map[string]string{} //initializatino for input
	person["name"] = "rifqi"

	//or

	var city = map[string]string{
		"City":   "ponorogo",
		"street": "sudirman",
	}
	fmt.Println(city["City"])
	fmt.Println(city["street"])

	for key, value := range city {
		fmt.Println(key, ":", value)
	}

	delete(person, "name")
	//can also combined with map

}
