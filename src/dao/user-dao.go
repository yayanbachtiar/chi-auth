package dao

import (
	"mama-chi/src/models"

	"gopkg.in/mgo.v2/bson"
)

// Login func
func Login(username string) (*models.Users, error) {
	var user models.Users
	if err := mongoDb.C("users").Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
