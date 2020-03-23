package dao

const (
	addUser = `insert into s_user()`
)

type User struct {
}

func (u User) AddUser() {
	db.Query(addUser)
}
