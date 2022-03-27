package basic

import "fmt"

func array() {
	var arr [3]string = [3]string{"a", "b", "c"}
	arr4 := [...]string{"c", "aaa"}
	fmt.Println(arr4)
	fmt.Println(arr)
}
