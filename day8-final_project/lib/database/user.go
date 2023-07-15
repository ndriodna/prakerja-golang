package database

import (
	"day_eight/config"
	"day_eight/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func FindUser(id int) (interface{}, error) {
	var user models.Users
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user models.Users) (interface{}, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(id int, val models.Users) (interface{}, error) {
	var user models.Users
	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}

	user.Name = val.Name
	user.Email = val.Email
	if e := config.DB.Updates(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.Users
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}
