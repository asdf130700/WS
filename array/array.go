package main

import "fmt"

func main() {
	var a = [2][3]int{ //2행 3열
		{1, 2, 3},
		{4, 5, 6},
	}

	var b = [3]int{ //3열 배열 [원소 하나처럼 보기]
		1, 2, 3,
	}

	var c = [2][2][3]int{ //3차원배열
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	var d = [4][2][2][3]int{ //4차원 배열 전체를 싸는 중괄호가 필요
		{
			{
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},

		{
			{
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},

		{
			{
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},

		{
			{
				{1, 2, 3},
				{4, 5, 6},
			},

			{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
	}

	var e = [5]float64{
		1.0, 2.0, 3.0, 4.0, 5.0,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("%.2f, %.2f, %.2f, %.2f, %.2f\n", e[0], e[1], e[2], e[3], e[4])

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%.2f, ", e[i])
	// }

	for i := 0; i < 4; i++ { //i는 for문 안에서만 유효함
		fmt.Printf("%.2f, ", e[i])

	}
	fmt.Printf("%.2f\n", e[4])

	var day = [7]string{
		"월요알", "화요일", "수요일", "목요일", "금요일", "토요일", "일요일",
	}

	for f := 0; f < 7; f++ {
		fmt.Printf("%s ", day[f])
	}

	var test [7]string

	for q := 0; q < 7; q++ {
		fmt.Scanf("%s", &test[q])
	}

	fmt.Printf("%s", test)

}
