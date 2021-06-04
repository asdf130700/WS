package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d * %d = %d", i, j, i*j)
		}
	}

}
