package main

import (
	"fmt"
	"time"
)

func main() {
	//wg := sync.WaitGroup{}
	data1 := []interface{}{"bisa1", "bisa2", "bisa3", "bisa4"}
	data2 := []interface{}{"coba1", "coba2", "coba3", "coba4"}

	for i := 1; i <= 4; i++ {
		go goroute1(data1, i)
		go goroute2(data2, i)
	}
	//goroute1(data2)
	time.Sleep(2 * time.Second)
}

func goroute1(input interface{}, i int) {
	fmt.Println(input, i)
}
func goroute2(input interface{}, i int) {
	fmt.Println(input, i)

}
