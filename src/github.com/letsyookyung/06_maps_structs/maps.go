package main

import "fmt"

// 1.what is map?
// 2.creating
// 3.manipulation

func main() {
	statePopulations := make(map[string]int) // 정의하고 나중에 값을 넣어두 됨
	statePopulations = map[string]int{       //map[KEY_type]VALUE_type{KEY:VALUE}
		"California": 3925,
		"texas":      2786,
		"new york":   5555,
		"ohio":       1234,
		"seoul":      8888,
	}
	//m := map[[]int]string{}         // array is valid key type, but slices is not
	m2 := map[[3]int]string{
		{1}: "hello",
		{2}: "bye",
	} // working

	fmt.Println("1) ", statePopulations, "\n") // map[California:3925 new york:5555 ohio:1234 texas:2786]
	statePopulations["georgia"] = 1666
	fmt.Println("2) ", statePopulations, "\n") // map[California:3925 georgia:1666 new york:5555 ohio:1234 texas:2786]
	// 뭔가를 추가했을 때 map의 순서가 달라짐
	delete(statePopulations, "seoul")
	fmt.Println("3) ", statePopulations, "\n")
	fmt.Println(statePopulations["seoul"]) // delete했는데 0나옴, 이유
	fmt.Println(statePopulations["oho"])   // 오타난 key를 넣어도? 0나옴, 이유

	pop, ok := statePopulations["oho"] // 오타 //ok 넣어서 진짜로 map에 있는지 없는지 확인 가능
	fmt.Println(pop, ok)               // 0, false
	fmt.Println(m2)

	fmt.Println(len(statePopulations))

	// map을 복사?해올 때 call by reference 방식으로 되서, 원래의 map에도 변형이 가해짐
	sp := statePopulations
	delete(sp, "ohio")
	fmt.Println("\nsp: ", sp)                             // map[California:3925 georgia:1666 new york:5555 texas:2786]
	fmt.Println("\nstatePopulations: ", statePopulations) //map[California:3925 georgia:1666 new york:5555 texas:2786]
}
