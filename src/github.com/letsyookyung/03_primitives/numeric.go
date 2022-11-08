package main

import "fmt"

//func main() {
//	//n := 42
//	//var n unit16 = 42
//	var a float32 = 10.0
//	var b float32 = 3.0
//
//	fmt.Printf("%v, %T\n", a+b, a+b)
//	fmt.Printf("%v, %T\n", a-b, a-b)
//	fmt.Printf("%v, %T\n", a*b, a*b)
//	fmt.Printf("%v, %T\n", a/b, a/b)
//	//fmt.Printf("%v, %T\n", a%b, a%b)
//}

// type이 같아야 연산이 가능
//int8 : -128 ~ 127, int16 : -32 768 ~ 32 767
//unit8 : 0~255, unit16: 0~65 535 ~

//func main() {
//	a := 10 // 1010(binary 10)
//	b := 3  // 0011(3)
//
//	fmt.Println(a & b)  //0010 = 2
//	fmt.Println(a | b)  //1011 =11
//	fmt.Println(a ^ b)  //1001 =9
//	fmt.Println(a &^ b) //and not 0100 =8
//}

//func main() {
//	a := 8              //2^3
//	fmt.Println(a << 3) //=64 // 2^3 * 2^3 = 2^6
//	fmt.Println(a >> 3) //=1 // 2^3 / 2^3 = 2^0
//}

//func main() {
//	n := 3.14 // var n float64 = 3.14
//	n = 13.7e72
//	n = 2.1e14
//	fmt.Printf("%v, %T", n, n)
//}

//func main() {
//	var n complex64 = 1 + 2i
//	fmt.Printf("%v, %T", n, n)
//
//}

//func main() {
//	a := 1 + 2i
//	b := 2 + 5.2i
//	fmt.Println(a + b)
//	fmt.Println(a - b)
//	fmt.Println(a * b)
//	fmt.Println(a / b)
//
//	fmt.Printf("%v, %T", n, n)
//
//}

func main() {
	//var n complex128 = 1 + 2i
	//fmt.Printf("%v, %T\n", real(n), real(n))
	//fmt.Printf("%v, %T\n", imag(n), imag(n))
	var n complex128 = complex(5, 21)
	fmt.Printf("%v, %T\n", n, n)

}
