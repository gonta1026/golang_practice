package array

import (
	"fmt"
	"strings"
)

func Comma() {
	// stringならJoin一発
	fruits := []string{"apple", "orange", "lemon"}
	fmt.Println(strings.Join(fruits, ",")) // => "apple,orange,lemon"
	fmt.Println(strings.Join(fruits, "#")) // => "apple#orange#lemon"
	fmt.Println(strings.Join(fruits, "-")) // => "apple-orange-lemon"
	//Joinはstringしかできない。。map,reduceみたいなのもないし。やるとしたら一旦中身をstringにしてそこからjoinする
	numbers := []int{55, 49, 100, 100, 0}
	strNumbers := []string{}
	for _, num := range numbers {
		newNum := fmt.Sprintf("%d", num)
		strNumbers = append(strNumbers, newNum)
	}
	fmt.Println(strings.Join(strNumbers, "-")) //
	numbers02 := []int{55, 49, 100, 100, 0}
	str := ""
	for _, v := range numbers02 {
		str += fmt.Sprintf("%d,", v)
	}
	str = strings.TrimRight(str, ",") //右端の","を取り除く
	fmt.Println(str)                  // => "55,49,100,100,0"

	//チートなやり方としては、Sprintfの整形を利用する。。
	str = fmt.Sprintf("%v", numbers02)
	str = strings.Trim(str, "[]")
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println(str) // => "55,49,100,100,0"

	//うーんこれもかなり辛い。。
	fruits2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
	}
	str = ""
	for _, v := range fruits2 {
		w, ok := v.([]interface{})
		if !ok {
			str += fmt.Sprintf("%v,", v)
		} else {
			for _, x := range w {
				str += fmt.Sprintf("%v,", x)
			}
		}
	}
	str = strings.TrimRight(str, ",") //右端の","を取り除く
	fmt.Println(str)                  // => "3,apple,250,orange,400"
}
