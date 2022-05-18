package main

import (
	"fmt"
)

func sliceFunc(i []int, num int) []int {
	for j := 0; j < num; j++ {
		i = append(i, 3)
	}
	i[0] = num
	return i
}
func main() {
	s := make([]int, 0, 3)
	s = append(s, 1)
	fmt.Printf("cap=%d\ts=%v\n", cap(s), s)
	fmt.Println("=========")
	s1 := sliceFunc(s, 2)
	fmt.Printf("cap=%d\ts=%v\n", cap(s), s)
	fmt.Printf("cap=%d\ts1=%v\n", cap(s1), s1)
	fmt.Println("=========")
	s2 := sliceFunc(s, 3)
	fmt.Printf("cap=%d\ts=%v\n", cap(s), s)
	fmt.Printf("cap=%d\ts1=%v\n", cap(s1), s1)
	fmt.Printf("cap=%d\ts2=%v", cap(s2), s2)
}
