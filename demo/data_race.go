package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0

	go func() {
		for j := 0; j < 5000; j++ {
			a++
		}
	}()
	go func() {
		for j := 0; j < 5000; j++ {
			a++
		}
	}

	time.Sleep(2*time.Second)

	fmt.Println(a)
}