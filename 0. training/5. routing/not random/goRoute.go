package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mtx := sync.Mutex{}
	data1 := []interface{}{"bisa1", "bisa2", "bisa3", "bisa4"}
	data2 := []interface{}{"coba1", "coba2", "coba3", "coba4"}

	for i := 1; i <= 4; i++ {
		go goroute1(data1, i, &mtx)
		time.Sleep(1 * time.Millisecond)
		go goroute2(data2, i, &mtx)
		time.Sleep(1 * time.Millisecond)
	}
	//goroute1(data2)
}

func goroute1(input interface{}, i int, mtx *sync.Mutex) {

	mtx.Lock()
	fmt.Println(input, i)
	mtx.Unlock()
}
func goroute2(input interface{}, i int, mtx *sync.Mutex) {
	mtx.Lock()
	fmt.Println(input, i)
	mtx.Unlock()
}
