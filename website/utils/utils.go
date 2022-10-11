package utils

import (
	"Xpress/models"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.Get(key).(string)
}

func CheckUserInputs(u models.User) string {

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

func GeneratePackageCode(l int) string {

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	s := make([]byte, l)
	for i := range s {
		s[i] = charset[seededRand.Intn(len(charset))]
	}
	return (string(s))
}
