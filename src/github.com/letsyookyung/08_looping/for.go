package main

import "fmt"

// initializer; for looping 조건;
//func main() {
//	for i, j := 0, 3; i < 5; i, j = i++, j+1 {
//		fmt.Println(i, j) //이렇게 하면 안됨
//	}

// 위에거와 아래거의 차이점은 i는 main scope, for문 scope에 속해서, for문 밖에서 안된다.
//func main() {
//	for i := 0; i < 5; i++ {
//		fmt.Println(i)
//		if i%2 == 0 {
//			i /= 2
//			fmt.Println("1)", i, "\n")
//		} else {
//			i = 2*i + 1
//			fmt.Println("2)", i, "\n")
//		}
//	}
//}

//func main() {
//	i := 0
//	for i = 3; i < 5; i++ {
//		fmt.Println(i)
//		if i%2 == 0 {
//			i /= 2
//			fmt.Println("1)", i, "\n")
//		} else {
//			i = 2*i + 1
//			fmt.Println("2)", i, "\n")
//		}
//	}
//	fmt.Println(i) //8
//}

//func main() {
//	for i < 5 {
//		i++
//		fmt.Println(i)
//	}
//}

// while 같이 infinite 돌수록
//func main() {
//	for i := 0; ; {
//		fmt.Println(i)
//		i++
//		if i == 100 {
//			break
//		}
//	}
//}

// continue
//func main() {
//	for i := 0; i < 10; i++ {
//		if i%2 == 0 {
//			continue
//		}
//		fmt.Println(i)
//	}
//}

//func main() {
//Loop: //이런 레이블을 쓰면 break가 어디에 먹혀야 하는지 표시해둘 수 있음
//	for i := 1; i <= 9; i++ {
//		fmt.Printf("%d단\n", i)
//		for j := 1; j <= 9; j++ {
//			fmt.Printf("%d x %d = %d\n", i, j, i*j)
//			if i*j >= 50 {
//				break Loop // 이렇게 하면 현재의 for문만 나감
//			}
//		}
//	}
//}

// looping of collections
func main() {
	s := []int{1, 2, 3}  // slices
	t := [3]int{1, 2, 3} // arrays
	maps := map[string]int{
		"cal": 3423,
		"tod": 1234,
	} // maps
	strings := "hello go!"
	fmt.Println(s[1])
	for k, v := range s {
		fmt.Println(k, v)
	}
	for k, v := range t {
		fmt.Println(k, v)
	}
	for _, v := range maps {
		fmt.Println(v)
	}
	for k, v := range strings {
		fmt.Println(k, string(v))
	}
}
