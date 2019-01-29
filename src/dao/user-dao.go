package dao

import (
	"log"

	"mama-chi/src/models"

	"gopkg.in/mgo.v2/bson"
)

// Login func
func Login(username string) (*models.Users, error) {
	var user models.Users
	if err := mongoDb.C("users").Find(bson.M{"username": "ngadimin"}).One(&user); err != nil {
		log.Println(username)
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
