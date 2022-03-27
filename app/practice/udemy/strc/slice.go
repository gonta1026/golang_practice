package strc

import "fmt"

type Users []User

func Slice() {

	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 30}
	user5 := User{Name: "user4", Age: 30}

	users := Users{}
	newUsers := append(users, user1, user2, user3, user4, user5)
	makeUsers := make([]User, 0)
	newMakeUsers := append(makeUsers, makeUsers...)
	fmt.Println(newMakeUsers)
	fmt.Println(users)
	fmt.Println(newUsers)
	for _, user := range newUsers {
		fmt.Println(user.Name)
	}
}
