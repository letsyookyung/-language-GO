package main

//func main() {
//	if false {
//		fmt.Println("the test is true.")
//	}
//}

//func main() {
//	statePopulations := map[string]int{ //map[KEY_type]VALUE_type{KEY:VALUE}
//		"California": 3925,
//		"texas":      2786,
//		"new york":   5555,
//		"ohio":       1234,
//		"seoul":      8888,
//	}
//	//if pop, ok := statePopulations["texas"]; ok {
//	//	fmt.Println(pop)
//	//}
//}

// comparison / operator (&&, <=, != ~)
//func main() {
//	number := 140
//	guess := 100
//
//	if guess < 1 {
//		fmt.Println("the guess must be greater than 1")
//	} else if guess > 100 {
//		fmt.Println("the guess must be less than 100")
//	} else {
//		if guess < number {
//			fmt.Println("TOO LOW")
//		}
//		if guess > number {
//			fmt.Println("TOO high")
//		}
//		if guess == number {
//			fmt.Println("you got it")
//		}
//	}
//}

import (
	"fmt"
	"math"
)

func main() {
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("these are the same")
	} else {
		fmt.Println(("these are different"))
	}

}
