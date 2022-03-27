package strc

import "fmt"

type T struct {
	User User
}

func Embedded() {
	fmt.Println("Embedded!!")
	t := T{User: User{
		Name: "aaa",
		Age:  10,
	}}
	t.User.setName2("keisei")
	fmt.Println(t)
}
