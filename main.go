package main

import (
	"GoLearning/week1"
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ret, err := week1.DeleteAt(s, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
	fmt.Println()

	s2 := []string{"4", "5", "6"}
	ret2, err := week1.DeleteAt(s2, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret2)
	}
}
