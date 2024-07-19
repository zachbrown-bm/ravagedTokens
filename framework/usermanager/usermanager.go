package usermanager

import "ravagedTokens/models"

func GetCurrent() models.User {
	return models.User{
		Name:  "Zach Brown",
		Email: "myname@gmail.com",
		Type:  models.Admin,
	}
}
