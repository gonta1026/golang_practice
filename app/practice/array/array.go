package array

import "fmt"

// sliceにmapをappendするパターン
func ArrayTemp() {

	tmp := []interface{}{"apple", 10, 2.5}
	fmt.Println(tmp) // => "[apple 10 2.5]"

	fruites2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
		[]interface{}{"lemon", 300},
	}
	fmt.Println((fruites2[0].(int)) + 11111)
	// fmt.Println(fruites2[0].(int))
	f, _ := fruites2[1].([]interface{})
	fmt.Println(f[1]) // => "250"

	f, _ = fruites2[3].([]interface{})
	fmt.Println(f[0]) // => "lemon"
}

func SliceOfSlice() {
	dary := make([][]int, 4)
	fmt.Println(dary) // => "[[] [] [] []]"
	for i := range dary {
		dary[i] = make([]int, 3)
	}
	dary[0][1] = 7
	fmt.Println(dary) // => "[[0 7 0] [0 0 0] [0 0 0] [0 0 0]]"
}

func sliceInMap() {
	var m1 = map[string]string{ // 値を指定するとき
		"dora": "doraのvalue",
		"nobi": "nobiのvalue",
	}
	var slice []map[string]string
	slice = append(slice, m1)
	fmt.Println(slice[0]["dora"])
}

// sliceにmapを最初に入れているパターン
func sliceInMap02() {
	type Map map[string]string
	slice1 := []Map{
		{
			"dora": "doraのvalue",
			"nobi": "nobiのvalue111",
		},
		{
			"dora": "doraのvalue",
			"nobi": "nobiのvalue222",
		}}
	fmt.Println(slice1)

	for _, value := range slice1 {
		fmt.Println(value["nobi"])
	}
}

func ArrayPractice() {
	// ArrayTemp()
	SliceOfSlice()
	// sliceInMap()
	// sliceInMap02()
}
