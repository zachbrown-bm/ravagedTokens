package usermanager

import "ravagedTokens/models"

func GetCurrent() models.User {
	return models.User{
		Id:    100,
		Name:  "Zach Brown",
		Email: "myname@gmail.com",
		Type:  models.Admin,
	}
}
