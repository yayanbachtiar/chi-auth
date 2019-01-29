package models

import (
	"time"
)

// OauthAccess models
type OauthAccess struct {
	ID         int8      `bson:"AUTO_INCREMENT" json:"-"`
	AppID      string    `bson:"app_id" json:"-"`
	UserID     string    `bson:"user_id" json:"user_id"`
	AccessKey  string    `bson:"access_key" json:"access_key" `
	UpdatedAt  time.Time `bson:"updated_at" json:"-" `
	CreatedAt  time.Time `bson:"created_at" json:"-" `
	ExpiredAt  int64     `bson:"expired_at" json:"expired_at" `
	RefreshKey string    `bson:"refresh_key" json:"-" `
}
