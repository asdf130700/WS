package main

import (
	"fmt"
)

func main() {
	s := []int{10, 3, 21, 7, 41, 1}
	t := make([]int, 4, 20)

	copy(t, s) //range 와의 차이점 range는 length의 길이에 따라 오류 copy시 모자란 만큼 복사
	fmt.Println(t)
}
