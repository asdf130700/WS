package main

import "fmt"

func main() {
	// var p = 0x0001
	// var a = 10

	// fmt.Printf("%p\n", &p)
	// fmt.Printf("%p\n", &a)

	i := [2][2][3]int{
		{
			{0, 1, 2},
			{3, 4, 5},
		},
		{
			{6, 7, 8},
			{9, 10, 11},
		},
	}
	fmt.Println(i)
}
