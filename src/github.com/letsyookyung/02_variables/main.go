package main

import (
	"fmt"
)

//var actorName string = "elisabeth sladen"
//var companion string = "sarah jane smith"
//var doctorNumber int = 3
//var season int = 11

var (
	actorName    string = "elisabeth sladen"
	companion    string = "sarah jane smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	I int = 1
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T", j, j)
}
