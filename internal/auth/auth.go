package auth

import (
	"golang.org/x/crypto/bcrypt"
	models_ "webBooks/internal/models "
	repository_ "webBooks/internal/repository "
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(repo *repository_.Repository, username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	user := &models_.Users{Username: username, Password: hashedPassword}
	return repo.CreateUser(user)
}

func AuthenticateUser(repo *repository_.Repository, username, password string) (*models_.Users, bool) {
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		return nil, false
	}
	return user, CheckPasswordHash(password, user.Password)
}
