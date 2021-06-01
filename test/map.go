package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)

	m[12] = "aaaa"
	m[23] = "bbbb"
	m[3939] = "cccc" //key와 value 쌍이 만들어짐

	str := m[12]
	fmt.Println(str)

	noData := m[33]
	fmt.Println(noData)

	var n map[string]int
	n = make(map[string]int)

	n["오후"] = 10

	value, exists := n["오후"]
	fmt.Println(value, exists)

	value, exists = n["이순신"]
	fmt.Println(value, exists)

	//fmt.Println(n["이순신"]) // 정수 타입이므로 default 값인 0이 나옴

	delete(n, "오후")
	value, exists = n["오후"]
	fmt.Println(value, exists)
}
