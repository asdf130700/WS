package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Korean  int
	English int
	Math    int
}

func main() {
	// for i := 2; i < 9; i++ {
	// 	for j := 1; j < 9; j++ {
	// 		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	// 	}
	// }
	//3.

	// a := 5
	// var b = 3.141592
	// c := 'A'
	// d := int64(128)
	// var e float32 = 0.125

	// fmt.Printf("%T,%T,%T,%T,%T", a, b, c, d, e)
	//1.

	//5.

	// for j := 0; j < len(i); j++ {
	// 	i := [2][2][3]int{
	// 	fmt.Printf(i[j])
	// 	}

	// }

	// a := [5]byte{'A', 'B', 'C', 'D', 'E'}
	// for n, val := range a {
	// 	fmt.Println(n, val)
	// }

	User := Score{Korean: 90, English: 92, Math: 82}
	// fmt.Println(User.Korean, User.English, User.Math)

	s := []int{User.Korean, User.English, User.Math}
	fmt.Printf("%d\n", s) //오름차순 정렬
	// sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sort.Sort(sort.IntSlice(s))

	// fmt.Printf("%d\n", s)
	// sort.Sort(sort.IntSlice(s)) //오름차순 정렬
	// fmt.Printf("%d\n", s)
	// sort.Sort(sort.Reverse(sort.IntSlice(s))) //내림차순 정렬
	// fmt.Printf("%d\n", s)

	// i := [2][2][3]int{
	// 	{
	// 		{
	// 			{1, 2, 3},
	// 			{4, 5, 6},
	// 		},

	// 		{
	// 			{7, 8, 9},
	// 			{10, 11, 12},
	// 		},
	// 	},
	// 	{
	// 		{
	// 			{13, 14, 15},
	// 			{16, 17, 18},
	// 		},

	// 		{
	// 			{19, 20, 21},
	// 			{22, 23, 24},
	// 		},
	// 	},
	// }
	// fmt.Println(i)
}
