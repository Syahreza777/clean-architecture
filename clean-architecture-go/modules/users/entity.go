package users

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "yaza barudak"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("opewfjdi3f84f339fu3")

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
}

type MyClaims struct {
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	jwt.StandardClaims
}

