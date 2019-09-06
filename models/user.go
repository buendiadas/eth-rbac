package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Account string `json:"account"`
	Role    string `json:"role"`
}

// Given an address, returns a user.
func GetUserByAddress(address string) (user User, err error) {
	err = DB.First(&user, "account = ?", strings.ToLower(address)).Error
	return
}
