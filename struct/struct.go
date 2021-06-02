package main

import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var p1 person
// 	p2 := person{age: 30, name: "lee"} //필드 순서는 상관 없음
// 	//var p1 person
// 	//var p2 person
// 	p3 := new(person)
// 	//var p3 *person
// 	fmt.Println(p1, p2)
// 	fmt.Println(p1.age, p2.age, p3.age)
// 	fmt.Printf("%p, %p, %p\n", &p1, &p2, p3)
// }
type User struct {
	Name string
	age  int
	Id   string
}
type VIPuser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	normalUser := User{Name: "이순신", age: 55, Id: "Lee"}
	vipUser := VIPuser{User{Name: "이순신", age: 55, Id: "Lee"}, 5, 100000}
	fmt.Println(normalUser, vipUser) //구조체는 출력시 중괄호로 나옴 배열은 []대괄호
	fmt.Printf("유저이름 : %s, 나이 : %d, 아이디 : %s\n", normalUser.Name, normalUser.age, normalUser.Id)
	fmt.Printf("VIP유저 이름 : %s, ", vipUser.UserInfo.Id)

	vipUser2 := vipUser
	fmt.Println(vipUser2)
}
