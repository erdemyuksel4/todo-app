package usecase

import (
	"todoapp/config"
	"todoapp/domain/model"
)

func GetUserById(id int) *model.User {
	for _, user := range config.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}
