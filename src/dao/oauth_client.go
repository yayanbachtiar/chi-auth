package dao

import (
	"log"
	"mama-chi/src/models"
)

// GetAllClient func
func GetAllClient() []models.Client {
	var oauthClient []models.Client
	if err := mongoDb.C("oauth_clients").Find(nil).All(&oauthClient); err != nil {
		log.Fatalf("errror %v", err)
	}
	log.Println(oauthClient)
	// mariaDb.Find(&oauthClient)
	return oauthClient
}
