package models

// Client Models
type Client struct {
	ID        int8   `json:"id" gorm:"AUTO_INCREMENT"`
	ClientKey string `json:"client_key" gorm:"client_key"`
	SecretKey string `json:"secret_key" gorm:"secret_key"`
	Name      string `json:"name" gorm:"name"`
}
