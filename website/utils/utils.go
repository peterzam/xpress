package utils

import (
	"Xpress/models"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.Get(key).(string)
}

func CheckUserInputs(u models.User) string {
	// Check Empty
	if u.Name == "" || u.Phone == "" || u.Address == "" || u.Password == "" {
		return "Inputs cannot be blanks"
	}

	// Check passowrd 8 characters
	if len(u.Password) < 8 {
		return "Password must be at least 8 characters"
	}

	// Check if already in DB
	if DB.Where("phone = ?", u.Phone).First(&models.User{}).Error == nil {
		return "Phone number is already registered"
	}

	return ""
}
