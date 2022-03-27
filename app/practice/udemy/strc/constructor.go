package strc

import "fmt"

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

type Comment struct {
	Content string
}

type Entity struct {
	User    User
	Comment Comment
}

func NewComment(content string) *Comment {
	return &Comment{Content: content}
}
func NewEntity(user User, comment Comment) *Entity {
	return &Entity{User: user, Comment: comment}
}

func Constructor() {
	user := NewUser("keisei_desu", 1000)
	user.setName("hogehoge")
	user.Age = 20001000
	fmt.Println(user)

	comment := NewComment("commentだー")
	fmt.Println(comment.Content)

	user1 := User{"keisei", 100}
	var comment1 Comment

	entity := NewEntity(user1, comment1)
	entity.User.setName("kkeekekek")
	fmt.Println(entity.User)
}
