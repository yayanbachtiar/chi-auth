package services

import (
	"fmt"
	"mama-chi/src/dao"
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
