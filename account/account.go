package main

import "fmt"

type account struct {
	balance int // balance : 잔고
}

func withdrawFunc(a *account, amount int) { //withdraw 인출
	a.balance -= amount
}
func (a *account) withdrawMethod(amount int) { //a *account 인자로 계좌번호를 받음 포인터로 access 가 가능하다 a account 함수 내부에서만 바뀌고 수정이 불가해서 포인터를 써야함
	a.balance -= amount
}
func main() {
	a := &account{100} //a가 포인터형 주소를 담고 있음
	withdrawFunc(a, 30)
	a.withdrawMethod(30)           //메소드를 써서 함수를 호출 멤버함수를 호출하는 코드
	fmt.Printf("%d \n", a.balance) //(*a).balnce ??
	//두개의 함수를 호출하여 포인터를 넘겻기 때문에 40이 나옴

	var b myInt = 10
	fmt.Println(b.add(30))

}

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}
