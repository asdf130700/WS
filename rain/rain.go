package main

import (
	"fmt"
	// "bufio"
	// "os"
)

func rainover80(per int) bool {
	if per >= 80 {
		return true
	} else {
		return false
	}
}
func rainover20(per int) bool {
	if per >= 20 {
		return true
	} else {
		return false
	}
}

func rainunder20(per int) bool {
	if per < 20 {
		return true
	} else {
		return false
	}
}
func main() {
	var per int
	var degree int

	fmt.Println("온도, 습도를 입력하세요.")
	fmt.Scanf("%d %d", &degree, &per)
	fmt.Printf("온도 : %d 습도 : %d\n", degree, per)

	// if rainover80(per) == true && degree >= 25 {
	// 	fmt.Println("덥고 비가옵니다.")
	// } else if rainover20(per) == true && degree >= 25 {
	// 	fmt.Println("덥고 습합니다.")
	// } else if rainunder20(per) == true && degree >= 25 {
	// 	fmt.Println("야외 활동하기 좋습니다.")
	// } else if !(degree >= 25) && rainover80(per) == true || degree < 10 {
	// 	fmt.Println("야외 할동하기 좋지 않습니다")
	// } else {
	// 	fmt.Println("좋은 날씨 입니다.")
	// }

	if rain := rainover80(per); (rain == true) && degree >= 25 {
		fmt.Println("덥고 비가옵니다.")
	} else if rain := rainover20(per); (rain == true) && degree >= 25 {
		fmt.Println("덥고 습합니다.")
	} else if rain := rainunder20(per); (rain == true) && degree >= 25 {
		fmt.Println("야외 활동하기 좋습니다.")
	} else if rain := rainover80(per); !(degree >= 25) && (rain == true) || degree < 10 {
		fmt.Println("야외 할동하기 좋지 않습니다")
	} else {
		fmt.Println("좋은 날씨 입니다.")
	}
}

// func main() {
// 	var per int
// 	var degree int
// 	fmt.Println("기온과 강수 확률을 입력하세요")
// 	fmt.Scanf("%d, %d", &degree, &per)
// 	if per = rainOver80(rain); per && degree >= 25 {
// 		fmt.Println("덥고 비가옵니다.")
// 	} else if rainOver80 && degree >= 25 {
// 		fmt.Println("덥고 비가옵니다.")
// 	}
// }

// fmt.Print("입력한 온도 : %d\n 입력한 습도 : %d\n", degree, per)

// if rain := rainover80(per); (rain == true) && (degree >= 25) {
// 	printf("dhseh dd")
// }
