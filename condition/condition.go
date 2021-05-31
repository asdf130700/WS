package main

import "fmt"

func main() {
	tmp := 30

	fmt.Scanf("%d", &tmp)

	if tmp >= 28 {
		fmt.Println("에어컨을 켜세요")
	} else if tmp < 5 {
		fmt.Println("에어컨을 끄세요")
	} else {
		fmt.Println("에어컨이 동작중입니다")
	}
	fmt.Println("종료!!")
}
