package main

import "fmt"

// case should be unique
//func main() {
//	switch 2 {
//	case 1, 5, 10:
//		fmt.Println("one, five, ten")
//	case 2, 4:
//		fmt.Println("two, four")
//	default:
//		fmt.Println("not one or two")
//	}
//}

//func main() {
//	switch i := 2 + 3; i {
//	case 1, 5, 10:
//		fmt.Println("one, five, ten")
//	case 2, 4:
//		fmt.Println("two, four")
//	default:
//		fmt.Println("not one or two")
//	}
//}

// tagless syntax, allow to overlap
//func main() {
//	i := 3
//	switch {
//	case i <= 10 && i == 3:
//		i += 100
//		fmt.Println("less than or equal to ten", i)
//		//fallthrough // 다음으로 무조건, 조건에 안맞아도 넘어감
//	case i <= 20:
//		fmt.Println("less than or equal to twenty", i)
//	default:
//		fmt.Println("greater than twenty", i)
//	}
//}

// type switch
func main() {
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break // 이거 하면 다음꺼 안나옴
		fmt.Println("it is too late to print out")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is an array")
	default:
		fmt.Println("i is another type")
	}

}
