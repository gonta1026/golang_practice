package structs

import "fmt"

// 型名(構造体の名前)を type の後に書く
type Profile struct {
	// フィールド名と型を記述する
	Name string
	Age  int
}

type Uber struct {
	Name string
	Age  int
}

type MapUber map[string]string

func base() {
	// mapは参照渡しになるのだー
	mapUser := MapUber{"otsuka": "keisei", "yamada": "akira"}
	newMapUser := mapUser
	mapUser["otsuka"] = "akirasa-----"
	fmt.Println(newMapUser)

	/*************************:***/
	// structのuserなのだー
	user := &Uber{
		Name: "kesiei",
		Age:  33,
	}
	newUser := user
	user.Name = "新しいkesiei------------"
	fmt.Println(newUser)
}

func StructMethods() {
	base()
}
