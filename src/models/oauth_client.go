package models

// Client Models
type Client struct {
	ID        int8   `json:"id" gorm:"AUTO_INCREMENT" bson:"_id"`
	ClientKey string `json:"client_key" gorm:"client_key" bson:"key"`
	SecretKey string `json:"secret_key" gorm:"secret_key" bson:"secret"`
	Name      string `json:"name" gorm:"name" bson:"name"`
}
