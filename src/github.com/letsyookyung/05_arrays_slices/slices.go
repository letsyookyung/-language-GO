package main

import "fmt"

//func main() {
//	a := []int{1, 2, 3}
//	//b := &a // slices로 하면, &안써도 call bu reference가 됨
//	b := a
//	b[1] = 5
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Printf("len: %v\n", len(a))
//	fmt.Printf("capacity:%v\n", cap(a))
//}

//func main() {
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	b := a[:]
//	c := a[3:]
//	d := a[:6]
//	e := a[3:6]
//	a[5] = 42
//	fmt.Println(a) // [1 2 3 4 5 42 7 8 9 10]
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//	fmt.Println(e) // [4 5 42]
//}

//func main() {
//	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 배열: 입력하는 데이터 수에 맞게 자동으로 크기 조정
//	b := a[:]
//	c := a[3:]
//	d := a[:6]
//	e := a[3:6]
//	a[5] = 42
//	fmt.Println(a) // [1 2 3 4 5 42 7 8 9 10]
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//	fmt.Println(e) // [4 5 42]
//}

//func main() {
//	a := make([]int, 3)
//	fmt.Println(a) //[0 0 0]
//	fmt.Printf("len: %v\n", len(a)) // len:3
//	fmt.Printf("capacity:%v\n", cap(a)) // cap:3
//
//}

//func main() {
//	a := make([]int, 3, 100 )
//	fmt.Println(a) //[0 0 0]
//	fmt.Printf("len: %v\n", len(a)) //len:3
//	fmt.Printf("capacity:%v\n", cap(a)) // cap:100
//
//}

//func main() {
//	a := []int{}
//	fmt.Println(a)                      // [
//	fmt.Printf("len: %v\n", len(a))     //len:0
//	fmt.Printf("capacity:%v\n", cap(a)) // cap:0
//	a = append(a, 1)
//	fmt.Println(a)                      // [1]
//	fmt.Printf("len: %v\n", len(a))     //len:1
//	fmt.Printf("capacity:%v\n", cap(a)) // cap: 2
//	a = append(a, 1, 2, 3, 4, 5)
//	fmt.Println(a)                      // [1 1 2 3 4 5]
//	fmt.Printf("len: %v\n", len(a))     //len: 6
//	fmt.Printf("capacity:%v\n", cap(a)) // cap: 6
//}

// concatenate
//func main() {
//	a := []int{}
//	fmt.Println(a)                      // []
//	fmt.Printf("len: %v\n", len(a))     // len : 0
//	fmt.Printf("capacity:%v\n", cap(a)) // cap :0
//	a = append(a, []int{2, 3, 4, 5}...)
//	fmt.Println(a)                      // [2 3 4 5]
//	fmt.Printf("len: %v\n", len(a))     // len:4
//	fmt.Printf("capacity:%v\n", cap(a)) //cap:4
//	a = append(a, []int{2, 3, 4, 5}...)
//	fmt.Println(a)                      //[2 3 4 5 2 3 4 5]
//	fmt.Printf("len: %v\n", len(a))   //len:8
//	fmt.Printf("capacity:%v\n", cap(a)) //cap:8
//}

// stack, push, pop 등
//func main() {
//	a := []int{1, 2, 3, 4, 5}
//	b := a[1:]
//	fmt.Printf("-a슬라이스 : %v, 메모리 주소:%p, 길이: %d, 용량: %d\n", a, &a, len(a), cap(a))
//	//-슬라이스 : [1 2 3 4 5], 메모리 주소:0xc000010018, 길이: 5, 용량: 5
//	fmt.Printf("-b슬라이스 : %v, 메모리 주소:%p, 길이: %d, 용량: %d\n", b, &b, len(b), cap(b))
//	//-슬라이스 : [2 3 4 5], 메모리 주소:0xc000010030, 길이: 4, 용량: 4
//	a = append(a, 6)
//	var c []int
//	c = a
//	b = append(a, 10)
//	fmt.Printf("-a슬라이스 : %v, 메모리 주소:%p, 길이: %d, 용량: %d\n", a, &a, len(a), cap(a))
//	//-슬라이스 : [1 2 3 4 5 6], 메모리 주소:0xc000010018, 길이: 6, 용량: 10
//	fmt.Printf("-b슬라이스 : %v, 메모리 주소:%p, 길이: %d, 용량: %d\n", b, &b, len(b), cap(b))
//	//-슬라이스 : [1 2 3 4 5 6], 메모리 주소:0xc000010030, 길이: 6, 용량: 10
//	fmt.Printf("-c슬라이스 : %v, 메모리 주소:%p, 길이: %d, 용량: %d\n", c, &c, len(c), cap(c))
//
//}

// 추가되는 값이 용량을 넘어설 경우 용량이 2배 증가(주의해야됨!)
// 1. 복사된 값은 변형하여도 원본에 영향을 주지 않는 것을 확인 가능 = 값 참조
// 2. 잘라낸 값은 변형하자 원본에 영향을 주는 것을 확인 == 메모리(주소) 참조

func main() {
	e := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", e)
	f := e[1:]
	fmt.Printf("%v\n", f)
	g := e[:len(e)-1]
	fmt.Printf("%v\n", g)
	h := append(e[:2], e[4:]...)
	fmt.Printf("%v\n\n", h)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
}
