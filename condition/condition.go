package main

import "fmt"

//전역변수 --> data segment
var cnt int = 0

//
// void IncreaseAndReturn(){

// }
func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	//지역변수 --> stach segemnt

	// tmp := 30

	// fmt.Scanf("%d", &tmp)

	// if tmp >= 28 {
	// 	fmt.Println("에어컨을 켜세요")
	// } else if tmp < 5 {
	// 	fmt.Println("에어컨을 끄세요")
	// } else {
	// 	fmt.Println("에어컨이 동작중입니다")
	// }
	// fmt.Println("종료!!")

	// var i int = 10

	// i = IncreaseAndReturn()//한번은 함수가 호출 판단문에서는 호출이 되지 않음

	if true || IncreaseAndReturn() < 5 { //왼쪽이 호출되어 거짓이므로 실행되지 않음 오른쪽은 cnt 값으로 대체

		fmt.Println("1증가")
	}

}
