package main

import "fmt"

//const a = iota // iota : counter, when having enumerate =0

//const (
//	a = iota
//	//b = iota
//	b
//	c
//	//c = iota
//)
//
//const (
//	a2 = iota
//)
//
//// compiler가 알아서 새롭게 enumerate 하는지 확인?
//func main() {
//	fmt.Printf("%v\n", a)
//	fmt.Printf("%v\n", b)
//	fmt.Printf("%v\n", c)
//	fmt.Printf("%v\n", a2) // 0 1 2 0
//}

// ---
//
//const (
//	catSpecialist = iota
//	dogSpecialist
//	snakeSpecialist
//)
//
//func main() {
//	var specialistType int = catSpecialist
//	fmt.Printf("%v\n", specialistType == catSpecialist) // 0으로 true
//	fmt.Printf("%v\n", dogSpecialist)                   // 1
//}

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v %v \n", dogSpecialist, specialistType) // 2 0

}
