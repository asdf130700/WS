package main

import "fmt"

func main() {
	// var a int
	// var b int = 10
	// var p *int

	// a = b
	// fmt.Println(a)
	// fmt.Println(b)

	// p = &b
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(p)
	// fmt.Println(*p)
	// fmt.Println(&p)

	// var v int = 10
	// fmt.Scanf("%d", v)
	// fmt.Println("%v", &v)
	// fmt.Println(v)

	var names []string = []string{"이순신", "홍길동", "강감찬"}

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(names[i])
	// }

	// for n, val := range names {
	// 	fmt.Println(n, val)
	// }

	// for _, val := range names { //_받는 인덱스를 무시하고 싶을때
	// 	fmt.Println(val)
	// }

	for n, _ := range names { //_명시적으로 두번째를 처리하지 않게 함
		fmt.Println(n)
	}

}
