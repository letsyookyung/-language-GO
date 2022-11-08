package main

import (
	"fmt"
)

//const (
//	_  = iota             // ignore the first value by assigning to blank identifier
//	KB = 1 << (10 * iota) // left shift 10 times iota
//	MB
//	GB
//	TB
//	PB
//	EB
//	ZB
//	YB
//)

//func main() {
//	fileSize := 40000000000000000000.
//	fmt.Printf("%.2fGB\n", fileSize/GB)
//	fmt.Printf("%.2fZB\n", fileSize/ZB)
//}

//----

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthKorea
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is hq? %v\n", isHeadquarters&roles == isHeadquarters)

}
