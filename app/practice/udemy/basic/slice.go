package basic

import "fmt"

type User struct {
	ID   int
	Name string
}

func slice() {
	one()
}

func one() {
	var users []User
	fmt.Println(users)
	users = append(users, User{
		ID:   1,
		Name: "tanaka",
	})

	fmt.Println(users)
	// users.([]User)
	fmt.Println(users)
	fmt.Println(users[0].ID)
	fmt.Println(users[0].Name)

	makeda := make([]int, 5)
	fmt.Println(makeda)
	makeda[0] = 1000
	fmt.Println(makeda)

}
