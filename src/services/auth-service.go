package services

import (
	"fmt"
	"mama-chi/src/dao"
	"mama-chi/src/helper"
	"mama-chi/src/models"
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt"
)

// GetAllClient fnc
func GetAllClient() map[string]string {
	data := dao.GetAllClient()
	m2 := make(map[string]string)

	for _, value := range data {
		m2[value.ClientKey] = value.SecretKey
	}
	fmt.Println(m2)
	return m2
}

// Login FUnc
func Login(users models.Users, key string) (models.OauthAccess, error) {
	user, err := dao.Login(users.UserName)

	if err != nil {
		return models.OauthAccess{}, err
	}

	erro := helper.CheckPasswordHash(users.Password, user.Password)
	if erro == false {
		return models.OauthAccess{}, fmt.Errorf("Wrong Password")
	}

	// check on db, user has already create token last hour or not
	now := time.Now()
	sec := now.Add(24 * 30 * 12 * time.Hour).Unix()
	hs256 := jwt.NewHS256("to_secret")
	jot := &jwt.JWT{
		Issuer:         "gbrlsnchs",
		Subject:        string(user.UserName),
		Audience:       user.Type,
		ExpirationTime: sec,
		NotBefore:      now.Add(30 * time.Minute).Unix(),
		IssuedAt:       now.Unix(),
		ID:             strconv.FormatInt(user.ID, 10),
	}
	jot.SetAlgorithm(hs256)
	jot.SetKeyID(strconv.FormatInt(user.ID, 10))
	payload, ferr := jwt.Marshal(jot)
	if ferr != nil {
		// handle error
	}
	token, _ := hs256.Sign(payload)
	oauthAccess := models.OauthAccess{
		AccessKey: string(token),
		AppID:     key,
		ExpiredAt: sec,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    strconv.FormatInt(user.ID, 10),
	}
	// save
	// dao.SaveClientAccess(oauthAccess)
	return oauthAccess, nil
}
