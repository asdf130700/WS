package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	i := 0

// 	for {
// 	L1:
// 		for {
// 			for {
// 				if i == 0 {
// 					break L1
// 				}
// 			}
// 			fmt.Println("inner OK")
// 		}
// 		fmt.Println("outer Lable OK") //L1
// 	}
// 	fmt.Println("OK") //무한루프
// }

func main() {

	// s := []int{0, 1, 2, 3, 4, 5, 6}
	// t := s[2:5] //새로 슬라이스를 만듬 2-5번까지? 2에서 5-1번까지
	// fmt.Printf("%v", t)

	// s := []int{0, 1, 2, 3, 4, 5}
	// t := s
	// fmt.Printf("%v, %v", s, t)
	// fmt.Printf("%v %v", &s[0], &t[0])

	// s := []int{10, 11, 21, 31, 41, 51}
	// // t := make([]int, len(s)) //메모리 공간 할당
	// t := make([]int, 7, 7) //메모리 공간 할당
	// fmt.Printf("%d, %d\n", s, t)

	// for n, val := range s { //n에 0가 return
	// 	t[n] = val
	// }
	// fmt.Printf("%d, %d", s, t)

	s := []int{10, 3, 21, 7, 41, 59}
	fmt.Printf("%d\n", s) //오름차순 정렬
	sort.Ints(s)
	fmt.Printf("%d\n", s)
	sort.Sort(sort.IntSlice(s)) //오름차순 정렬
	fmt.Printf("%d\n", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s))) //내림차순 정렬
	fmt.Printf("%d\n", s)
}

// 	var b [3]int
// 	fmt.Printf("%v\n", &b[0])
// 	fmt.Printf("%x\n", &b[0])
// 	var a []int        //슬라이스
// 	a = []int{1, 2, 3} //슬라이스 초기화
// 	a[1] = 10
// 	fmt.Println(a)

// 	s := make([]int, 5, 10)        // 5개의 원소를 가지고 10개를 저장할수 있음
// 	fmt.Println(s, len(s), cap(s)) //cap 최대 용량량을 리턴
// }
