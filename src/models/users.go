package models

import (
	"strconv"
	"time"
)

// Users Models
type Users struct {
	ID           int64      `bson:"id" json:"column:id;AUTO_INCREAMENT"`
	Name         string     `bson:"name" json:"name"`
	UserName     string     `bson:"username" json:"username"`
	Password     string     `bson:"password" json:"password"`
	Email        string     `bson:"email" json:"email"`
	Type         string     `bson:"type" json:"type"`
	RegisteredAt time.Time  `bson:"registered_at" json:"registered_at"`
	CreatedAt    time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `bson:"deleted_at" json:"deleted_at" sql:"index"`
}

// GetID func
func (u Users) GetID() string {
	return strconv.FormatInt(u.ID, 10)
}
