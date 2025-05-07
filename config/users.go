package config

import (
	"todoapp/domain/model"
)

var Users = []*model.User{
	{
		ID:       1,
		Type:     model.TodoUser,
		UserName: "Erdem",
		Password: "123",
	},
	{
		ID:       2,
		Type:     model.TodoUser,
		UserName: "user2",
		Password: "321",
	},
	{
		ID:       3,
		Type:     model.Admin,
		UserName: "admin",
		Password: "123",
	},
}
