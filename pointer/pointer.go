package main

import "fmt"

// func main() {
// 	var numPtr *int
// 	fmt.Println(numPtr)//nil 값이 출력 stack 안에 생김 main 함수 안에 있기 때문

// }

type Data struct { //다른패키지에서 access 하려면 대문자로 써야함
	value int
	data  [100]int
}

func ChangeData(arg *Data) { //구조체의 인스턴스가 아닌 구조체의 포인터를 넣음
	// (*arg).data[99] = 5// 정확히 짜려면 하지만 go 언어는 생략이 가능하다
	// (*arg).value = 3
	arg.data[99] = 5 // 정확히 짜려면 하지만 go 언어는 생략이 가능하다
	arg.value = 3
}

// func changeInt(p *int) {
// 	*p = 11
// }
func main() {
	var ch rune = '한'
	var ech rune = 'A'

	str1 := "안녕하세요"
	str2 := "ABCDE"
	fmt.Printf("len(str1)= %d, len(str2) = %d\n", len(str1), len(str2))

	str3 := "hello 월드"
	for n, val := range str3 {
		fmt.Printf("index: %d, value :%c\n", n, val) //띄어쓰기도 1byte인식
	}
	str4 := "AAAAA"
	str5 := "a"
	str10 := str4 + str5
	// str6 := "BBB"
	// str7 := "ccccccccc" //ascii code 소문자가 더 큼

	fmt.Printf("%s > %s : %v\n", str4, str5, str4 > str5) //소문자 a
	fmt.Println(str10)
	fmt.Printf("%c, %c\n", ch, ech)
	fmt.Printf("%T, %T\n", ch, ech) //T는 변수의 타입을 찍어줌

	var data Data //100개의 변수가 자동으로 초기화된다.
	//ChangeData(data)//포인터로 선언해서 바로 데이터를 넘길수 없다.
	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[99]= %d\n", data.data[99])
}
