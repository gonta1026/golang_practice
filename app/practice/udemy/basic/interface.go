package basic

import "fmt"

func interfaceFunc() {
	var x interface{}
	fmt.Println(x)
	x = "sss"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
