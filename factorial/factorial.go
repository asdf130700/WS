package main

import "fmt"

// func factorial(i int) int {
// 	if i == 0 {
// 		return 1
// 	}
// 	return i * factorial(i-1)
// }
func main() {
	// var f int
	// fmt.Println("계산하고자 하는 값을 입력하세요")
	// fmt.Scanf("%d", &f)
	// fmt.Println("%d 팩토리얼 갑은 ... %d", f, factorial(f))
	var a int = 95
	var b int = 88

	switch true {
	case b > 90, a > 90:
		fmt.Println("모두다 A 학점입니다")
	case a > 80, b > 80:
		fmt.Println("모두다 B 학점입니다")
	default:
		fmt.Println("모두다 F 학점입니다")
	}
}
