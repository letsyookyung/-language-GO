package main

import (
	"fmt"
)

// byte : ASCII코드로 표현하는 타입
// 8bit

//
//func main() {
//	s := "this is a string"
//	fmt.Printf("%v, %T\n", s, s)               // this is a string, stirng
//	fmt.Printf("%v, %T\n", s[2], s[2])         // 105, unit8 => byte is an alias for a string
//	fmt.Printf("%v, %T\n", string(s[2]), s[2]) // i, unit8
//
//	s[2] = "u" // 이런거 안됨
//
//}

//func main() {
//	s := "this"
//	s2 := "thathat"
//	fmt.Printf("%v, %T\n", s+s2, s+s2)
//}

// string -> byte converting -> ASCII value로
//func main() {
//	s := "this is a string"
//	b := []byte(s)               //collection of bytes
//	fmt.Printf("%v, %T\n", b, b) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
//}

func main() {
	b := [16]byte{116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 115, 116, 114, 105, 110, 103} //collection of bytes
	s := string(b[:])
	fmt.Printf("%v, %T\n", s, s) // this is a string, string
}
