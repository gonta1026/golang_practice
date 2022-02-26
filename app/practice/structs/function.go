package structs

import "fmt"

// 型名(構造体の名前)を type の後に書く

type User struct {
	// フィールド名と型を記述する
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name + "です")
}
func (u *User) SetName(name string) {
	u.Name = name
}

func StructFunctionMethods() {
	i := 5
	p := &i
	fmt.Println(p)
	user1 := &User{
		Name: "User",
		Age:  33,
	}

	fmt.Println(user1.Age)
	user1.SayName()
	user1.SetName("akira")
	user1.SayName()
}
