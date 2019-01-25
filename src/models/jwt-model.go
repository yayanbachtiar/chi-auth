package models

import "github.com/gbrlsnchs/jwt"

// Token Model
type Token struct {
	*jwt.JWT
	IsLoggedIn  bool   `json:"isLoggedIn"`
	CustomField string `json:"customField,omitempty"`
}
