package strc

import "fmt"

func Map() {
	fmt.Println("hogehogehoge")
	m := map[int]User{
		1: {Name: "keisei1111", Age: 111},
		2: {Name: "keisei333", Age: 333},
		4: {Name: "keisei444", Age: 444},
	}
	m[1] = User{Name: "aefefekeisei1111", Age: 111}
	fmt.Println(m[1])

	// m2 := map[User]string{
	// 	{Name: "keke"}: "value1",
	// 	{Name: "aaa"}:  "value2",
	// }
	// m3 := map[User]User{
	// 	{Name: "bbbbbbbb"}: {Name: "bbbbbbbbb"},
	// 	{Name: "ccccc"}:    {Name: "ccccccc"},
	// }
}
