package main

// rune 타입은 유니코드(UTF-8/Unicode)를 쉽게 제어하기 위한 타입
// go 언어의 rune타입은 int32을 재정의해서 사용
// rune : 유니코드(UTF-8)을 표현하는 타입
// 유니코드 인코딩에서 한글은 3byte, 영어는 1byte를 사용함
// 유니코드는 한글을 지원하는 인코딩 방식이기 때문에, 한글을 사용하는 프로그램 개발에 유용하게 쓰일 수 있음
// go site가서 rune api 문서 읽어보기

import (
	"fmt"
)

//func main() {
//	r := 'a'
//	fmt.Printf("%v, %T\n", r, r) // 97, int32 (int32로)
//}

func main() {
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32 (go에서 rune은 int32이라서)
}
