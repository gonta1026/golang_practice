package basic

import "fmt"

func normal() {
	// fmt.Println("basic!")
	var number, number2 int = 100, 200

	fmt.Println(number)
	fmt.Println(number2)

	var (
		hoge int    = 200
		str  string = "sss"
	)
	var fuga int16 = 2000

	fmt.Println(hoge, str, fuga)

	var str1 string
	var num2 int

	str1 = "sss"
	num2 = 1222
	fmt.Println(str1, num2)

}
