package es6

import "fmt"

func square(x int) int {
	return x * x
}

func NumMapping(f func(int) int, array []int) []int {
	buff := make([]int, len(array))
	for i, v := range array {
		buff[i] = f(v)
	}
	return buff
}
func Mapping(f func(int) int, array []int) []int {
	buff := make([]int, len(array))
	for i, v := range array {
		buff[i] = f(v)
	}
	return buff
}

func filterd() {

	type Person struct {
		Name string
		Age  int
	}
	type NewPerson struct {
		Name string
		Age  int
		Bool bool
		Hoge string
	}

	// newPerson := NewPerson{Name: "ssss", Age: 10, Bool: false}
	persons := []NewPerson{
		{Name: "John", Age: 33},
		{Name: "aa", Age: 22},
		{Name: "Saaamith", Age: 24},
	}
	// newPerson.
	var filteredPerson []NewPerson
	for _, v := range persons {
		if v.Age <= 25 {
			v.Bool = false
			v.Hoge = ""
			v.Hoge = ""
			filteredPerson = append(filteredPerson, v)
		}
	}

	fmt.Println(filteredPerson)
}
func groupItems() {
	// newPerson := NewPerson{Name: "ssss", Age: 10, Bool: false}
	// type Person struct {
	// 	Name string
	// 	Age  int
	// 	Bool bool
	// 	Hoge string
	// }

	type GroupItem struct {
		ItemName string
		ItemId   int
	}
	type Group struct {
		GroupName  string
		GroupId    int
		GroupItems []GroupItem
	}

	firstGroupItem01 := GroupItem{
		ItemName: "aa", ItemId: 1,
	}
	firstGroupItem02 := GroupItem{
		ItemName: "aa", ItemId: 2,
	}
	firstGroupItems := []GroupItem{
		firstGroupItem01,
		firstGroupItem02,
	}

	// firstGroupItems :=
	lastGroupItem := GroupItem{
		ItemName: "aa", ItemId: 1,
	}
	lastGroupItem02 := GroupItem{
		ItemName: "aa", ItemId: 1,
	}
	lastGroupItems := []GroupItem{
		lastGroupItem, lastGroupItem, lastGroupItem02,
	}

	groups := []Group{
		{GroupName: "GroupName01", GroupId: 33, GroupItems: firstGroupItems},
		{GroupName: "GroupName02", GroupId: 34, GroupItems: lastGroupItems},
		{GroupName: "GroupName02", GroupId: 35, GroupItems: lastGroupItems},
	}
	type Item struct {
		GroupId int
		Name    string
	}
	type NewItem struct {
		GroupId int
		Name    string
		Hoge    bool
	}
	items := []Item{
		{
			GroupId: 34,
			Name:    "肉まん",
		},
		{
			GroupId: 35,
			Name:    "肉まん",
		},
		{
			GroupId: 36,
			Name:    "肉まん",
		},
		{
			GroupId: 37,
			Name:    "肉まん",
		},
	}
	var newItems []NewItem
	// var filteredPerson []NewPerson

	for _, item := range items {
		for _, group := range groups {
			if item.GroupId == group.GroupId {
				newItem := NewItem{
					GroupId: item.GroupId,
					Name:    item.Name,
					Hoge:    false,
				}
				newItems = append(newItems, newItem)
			}
		}

	}
	fmt.Println(newItems)
}

func Method() {
	groupItems()
}
