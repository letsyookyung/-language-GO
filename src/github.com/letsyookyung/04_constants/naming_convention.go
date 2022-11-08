package main

import "fmt"

//func main() {
//	const myConst int = 42
//	//myConst := 22 //  cannot assign to myConst
//	fmt.Printf("%v, %T\n", myConst, myConst) // 42, int
//}

//func main() {
//	//const myConst float64 = math.Sin(1.57)
//	//fmt.Printf("%v, %T\n", myConst, myConst) // compile time는 뭔가를 실행하는 함수같은 것을(runtime에 돌아야하는) 동시에 쓰면 안됨
//
//}

//	func main() {
//		const a int = 14
//		const b string = "foo"
//		const c float32 = 3.14
//		fmt.Print(a, b, c) // 14foo3.14
//	}

//const a int16 = 10
//
//// 안에 있는 const가 win, package level에서 선언했어도 다시 선언될 수 있음
//func main() {
//	const a int = 14
//	//const b string = "foo"
//	//const c float32 = 3.14
//	fmt.Print(a) // 14
//}

func main() {
	const a int = 42
	//var b int = 27
	var b int16 = 27                 // (mismatched types int and int16)
	fmt.Printf("%v, %T\n", a+b, a+b) // 69, int
}

func main() {
	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b) // 69, int16
}