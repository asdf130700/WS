package main

import "kerygma.com/structPackage"

func main() {
	original_sp := structPackage.StructPackage{}                       //객체를 선언 new가 아니기 때문에 포인터가 아님
	structPackage.ShowPackageInfo(&original_sp, "stopBlock", "v1.0.1") //객체의 주소가 넘어가서 받아옴
}
