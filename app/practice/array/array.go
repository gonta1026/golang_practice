package array

import "fmt"

// sliceにmapをappendするパターン
func arrayTemp() {

	tmp := []interface{}{"apple", 10, 2.5}
	fmt.Println(tmp) // => "[apple 10 2.5]"

	fruites2 := []interface{}{
		3,
		[]interface{}{"apple", 250},
		[]interface{}{"orange", 400},
		[]interface{}{"lemon", 300},
	}
	fmt.Println(fruites2)
	hoge := fruites2[0].(int) + 1
	fmt.Println(hoge)
	fmt.Println(fruites2[0])
	hogehoge := fruites2[1].([]interface{})
	fmt.Println("hogehogeは", hogehoge)
	f, _ := fruites2[1].([]interface{})
	fmt.Println(f[1]) // => "250"

	f, _ = fruites2[3].([]interface{})
	fmt.Println(f[0]) // => "lemon"
}

func sliceOfSlice() {
	hoge := make([]int, 4)
	hoge[3] = 3
	fuga := make([]int, 5) // 5この
	fmt.Println("fugaは", fuga)
	fmt.Println(hoge) // => "[0 0 0 0]"
	dary := make([][]int, 4)
	fmt.Println(dary) // => "[[0 7 0] [0 7 0] [0 7 0] [0 7 0]]"
	for i := range dary {
		dary[i] = make([]int, 3)
	}
	dary[0][1] = 7
	dary[1][1] = 7
	dary[2][1] = 7
	dary[3][1] = 7
	fmt.Println(dary) // => "[[0 7 0] [0 0 0] [0 0 0] [0 0 0]]"
}

type MapType = map[string]string
type MapType02 = []map[string]string
type MapIntType = map[int]string

func sliceInMap() {
	/* MapType */
	var m1 = MapType{ // 値を指定するとき
		"dora": "doraのvalue",
		"nobi": "nobiのvalue",
	}
	value, ok := m1["dora"]
	value02, ok02 := m1["doraa"]
	fmt.Println(value, ok)
	fmt.Println(value02, ok02)
	var slice []MapType
	fmt.Println(slice)
	slice = append(slice, m1)
	fmt.Println(slice[0]["dora"] + slice[0]["dora"])
	/* MapIntType */
	m2 := MapIntType{
		0: "aaaa",
		1: "bbb",
	}
	fmt.Println(m2)
	var slice02 []MapIntType
	slice02 = append(slice02, m2)
	fmt.Println(slice02)
}

// sliceにmapを最初に入れているパターン
func sliceInMap02() {
	slice1 := []MapType{
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
	slice2 := MapType02{
		{
			"dora": "doraのvalue",
			"nobi": "nobiのvalue111",
		},
		{
			"dora": "doraのvalue",
			"nobi": "nobiのvalue222",
		}}
	fmt.Println(slice2)

	for _, value := range slice2 {
		fmt.Println(value["nobi"])
	}
}

func ArrayPractice() {
	// arrayTemp()
	// sliceOfSlice()
	// sliceInMap()
	sliceInMap02()
}
