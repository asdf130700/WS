package main

import "kerygma.com/greeting"

//kerygma.com/greeting => ../greeting//이 문장을 만나면 대체해라 greeting폴더와 packageimporttest가 같은 경로에 있음

//소스에서만 컴파일 모르는 함수 호출 자동완성이 되지 않음 sayhello 모듈에서 하는 기능
func main() {
	greeting.SayHello()
}
