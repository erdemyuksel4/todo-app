package model

type UserType int

const (
	TodoUser UserType = 1
	Admin    UserType = 2
)

type User struct {
	ID        int         `json:"userId"`
	Type      UserType    `json:"user_type"`
	UserName  string      `json:"username"`
	Password  string      `json:"password"`
	TodoLists []*TodoList `json:"todolists"`
}
