package main

import "fmt"

//1. what is struct?
//2. creating
//3. naming convention
//4. embedding
//5. tags

// 비슷한 내용을 갖고 있는 것끼리 묶음, 하나의 struct 안에 하나하나씩의 타입에 대한 제한이 없음
//type Doctor struct {
//	number     int
//	actorName  string
//	companions []string //slice
//	episodes   []string
//}
//
//func main() {
//	aDoctor := Doctor{
//		number:    3,
//		actorName: "ivy",
//		companions: []string{
//			"fork",
//			"pycharm",
//		},
//	}
//
//	//bDoctor := Doctor{
//	//	4,
//	//	"ron",
//	//	[]string{
//	//		"fork",
//	//		"pycharm",
//	//	},
//	//} // 이렇게 되긴 하는데, 헷갈려지니
//
//	fmt.Println(aDoctor) // {3 ivy [fork pycharm]}
//	fmt.Println(aDoctor.actorName)
//	fmt.Println(aDoctor.companions[1])

//fmt.Println(bDoctor)
//}

// naming convention
// package level 에서는 앞글자 대문자 = is going to exported from the package
// starts with 소문자 = is going to be internal to the package

// map 과 다르게 copy 해도 call by value 형식
//func main() {
//	aDoctor := struct{ name string }{name: "ivy lee"}
//	anotherDoctor := aDoctor
//	anotherDoctor.name = "yookyung"
//	anotherNurse := &aDoctor
//	anotherNurse.name = "nexters"
//	fmt.Println(aDoctor)
//	fmt.Println(anotherDoctor)
//	fmt.Println(anotherNurse) // &참조하면 aDoctor까지 바뀜
//
//}

// embedded, composition through embedding

import (
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"` //tag
	Origin string
}

type Bird struct {
	Animal   // embedding
	SpeedKPH float32
	CanFly   bool
}

func main() {
	//b := Bird{}
	//b.Name = "Emu"
	//b.Origin = "Australia"
	//b.SpeedKPH = 48
	//b.CanFly = false
	//fmt.Println(b) // {Emu Australia} 48 false}
	//fmt.Println(b.Name) // Emu

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   true,
	}
	fmt.Println(c) // {{Emu Australia} 48 false}

	// tag 활용
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
