package users

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	Repo Repository
}

func (uc Usecase) Login(username, password string) (string, error) {
	user, err := uc.Repo.CheckUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: user.Username,
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	 return signedToken, nil
}

func (uc Usecase) AddUsers(user *User) error{
	hassPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hassPassword)
	
	result := uc.Repo.CreateUser(user)
	return result
}

func (uc Usecase) FindUsers(user *User) ([]User, error){
	hassPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hassPassword)

	result, err := uc.Repo.FindUsers()

	return result,  err
}

func (uc Usecase) DeleteUser(id string) error{
	err := uc.Repo.DeleteUser(id)

	return err
}
