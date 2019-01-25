package dao

import "mama-chi/src/models"

// GetAllClient func
func GetAllClient() []models.Client {
	var oauthClient []models.Client

	mariaDb.Find(&oauthClient)
	return oauthClient
}
