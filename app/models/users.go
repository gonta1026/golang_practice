package models

import (
	"time"
	"todo_app/utils"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	utils.LogFatalln(err)

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?,email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	utils.LogFatalln(err)

	return err
}

func (u *User) DaleteUser() (err error) {
	cmd := `DELETE FROM USERS WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	utils.LogFatalln(err)
	return err
}
