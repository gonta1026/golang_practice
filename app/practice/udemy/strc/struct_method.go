package strc

import "fmt"

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) setName(name string) {
	u.Name = name
}
func (u *User) setName2(name string) {
	u.Name = name
}

func StructMethod() {
	fmt.Println("aaaa")
	user1 := User{Name: "user1"}
	user1.SayName()
	user1.setName2("変更後のuser1")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.setName("変更後のuser2")
	user2.SayName()
}
