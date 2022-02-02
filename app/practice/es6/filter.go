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

type Person struct {
	Name string
	Age  int
}

func FilterMethod() {
	person := []Person{
		{Name: "John", Age: 33},
		{Name: "aa", Age: 22},
		{Name: "Saaamith", Age: 24},
		{Name: "Saaaaamith", Age: 23},
		{Name: "Smidddth", Age: 21},
	}
	var filteringPerson []Person
	for _, v := range person {
		if v.Age <= 25 {
			filteringPerson = append(filteringPerson, v)
		}
	}

	fmt.Println(filteringPerson)
}
