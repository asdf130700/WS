package main

import "fmt"

func main() {
	var a int
	var b int = 10
	var p *int

	a = b
	fmt.Println(a)
	fmt.Println(b)

	p = &b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

}
