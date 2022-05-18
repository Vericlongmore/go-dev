package main

import "fmt"

func main() {
	var a []int
	a = append(a, 0, 1, 2)
	a = append(a, 0, 1, 2)
	a = append(a, 0, 1, 2)

	aa := len(a)
	fmt.Println(cap(a))
	fmt.Println(aa)

}
