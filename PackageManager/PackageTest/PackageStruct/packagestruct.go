package structPackage //세부항목에 있는 패키지를 의미하는것
import "fmt"

type StructPackage struct { //패키지명과 버전 가정
	Name    string
	Version string
}

//func ShowPackageInfo(sp *StructPackage) {//포인터로 넘겨주면 new로 인자를 받는게 가능?
//sp := StructPackage{}//값을 default로 초기화
func ShowPackageInfo() { //내부에서 변수를 생성
	sp := StructPackage{Name: "goblock", Version: "v1.0.0"} //값을 default로 초기화
	fmt.Printf("Package Name : %s, Version : %s\n", sp.Name, sp.Version)
}
