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
func UploadFile() (filename string, success bool) {

	filename = "hello.exe"
	success = true
	return
}

func main() {
	if filename, success := UploadFile(); success {
		fmt.Println("Upload success", filename)
	} else {
		fmt.Println("Upload failed", filename)
	}
}
