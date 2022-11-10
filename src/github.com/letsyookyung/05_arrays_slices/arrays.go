package main

import "fmt"

// array는 하나의 한 타입만 가능
// array로 묶어서 돌리면 조금 더 빠름 (이미 메모리에 있음)
//func main() {
//	//grades := [3]int{97, 85, 93}
//	grades := [...]int{97, 85, 93, 99, 100, 101} // [...] = 정해놓지 않을 수도 있음
//	fmt.Printf("Grades: %v\n", grades)
//
//	var students [10]string
//	fmt.Printf("students: %v\n", students)
//	students[0] = "lisa"
//	students[2] = "ivy"
//	students[1] = "ron"
//	fmt.Printf("students: %v\n", students)
//	fmt.Printf("the number of students: %v\n", len(students))
//
//}
//
//func main() {
//	//var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} //[[1 0 0] [0 1 0] [0 0 1]]
//	//fmt.Println(identityMatrix)
//
//	var identityMatrix [3][3]int
//	identityMatrix[0] = [3]int{1, 0, 0}
//	identityMatrix[1] = [3]int{0, 1, 0}
//	identityMatrix[2] = [3]int{0, 0, 1}
//	fmt.Println(identityMatrix)
//
//}

//func main() {
//	a := [...]int{1, 2, 3}
//	b := a // call by value 형식
//	b[1] = 5
//	fmt.Println(a)
//	fmt.Println(b)
//}

func main() {
	a := [...]int{1, 2, 3}
	b := &a // call by reference 형식
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
