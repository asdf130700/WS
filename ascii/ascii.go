package main

import (
	// "os"
	// "bufio"
	"fmt"
)

func main() {
	// var ch1 byte = 65
	// var ch2 byte = 0101 //01 8진수
	// var ch3 byte = 0x41 //0x 16진수
	// var ch4 byte = 'A'

	// fmt.Printf("%d %c %c %c", ch1, ch2, ch3, ch4)

	// var uni1 rune = 44032
	// var uni2 rune = 0126000
	// var uni3 rune = 0xAC00
	// var uni4 rune = '가'

	// fmt.Printf("\n%c %c %c %c", uni1, uni2, uni3, uni4)

	// rawLiteral := `아리랑\n
	// 아리랑\n
	// 	아라리오`
	// interLiteral := "\n아리랑아리랑\n아라리오"

	// fmt.Print(rawLiteral)
	// fmt.Print(interLiteral)

	// fmt.Println()
	// fmt.Print(unsafe.Sizeof(rawLiteral))
	// fmt.Println()
	// fmt.Print(unsafe.Sizeof(interLiteral))

	// var a int
	// var b int

	// n, err := fmt.Scan(&a, &b)
	// if err != nil {
	// 	fmt.Println(n, err)
	// } else {
	// 	fmt.Println(n, a, b)
	// }

	// stdio := bufio.NewReader(os.Stdin) //새로운 버퍼객체를 리턴해준다.
	// //os가 핸들링 하고 있는 것을 어플리케이션에서 처리하게 받아옴

	// var a int
	// var b int

	// n, err := fmt.Scanln(&a, &b)

	// if err != nil {
	// 	fmt.Println(err) //오류 메시지를 그대로 출력
	// 	fmt.Println("다시 입력해주세요")
	// 	stdio.ReadString('\n') //넘겨진 인자 값을 만날때 까지 버퍼를 지움
	// } else {
	// 	fmt.Println(n, a, b)
	// }

	// n, err = fmt.Scanln(&a, &b)

	// if err != nil {
	// 	fmt.Println(err) //오류 메시지를 그대로 출력
	// 	fmt.Println("다시 입력해주세요")
	// 	stdio.ReadString('\n') //넘겨진 인자 값을 만날때 까지?
	// } else {
	// 	fmt.Println(n, a, b)
	// }

	// 	var i int = 10
	// 	var j int = 20
	// 	fmt.Println(i, j)
	// 	swap(&i, &j)
	// 	fmt.Println(i, j)
	// }
	// func swap(a *int, b *int) {
	// 	//*b, *a = *a, *b
	// 	temp := *a
	// 	*a = *b
	// 	*b = temp

	// a := 7      // 00000111
	// b := a << 2 // 00011100: a의 비트를 오른쪽으로 2번 이동
	// // fmt.Printf("%08b", b)//2진수 출력
	// fmt.Printf("%d", b)

	var c float64 = 0.1
	var d float64 = 0.2
	fmt.Println(1 == 1)             // true: 두 정수가 같으므로 true
	fmt.Println(3.5 == 3.5)         // true: 두 실수가 같으므로 true
	fmt.Println("Hello" == "Hello") // true: 두 문자열이 같으므로 true
	fmt.Println(c+d == 0.3)         //c + d가 의도와 다르게 저장됨
	fmt.Println(c + d)

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b) // true: 두 배열이 같으므로 true
}
