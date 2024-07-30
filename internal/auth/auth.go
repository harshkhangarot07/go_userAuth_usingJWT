package auth

import (
	"errors"
	"go_userAuth/internal/models"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]models.User{}
var mu sync.Mutex

func RegisterUser(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users[username] = models.User{Username: username, Password: string(hashedPassword)}
	return nil
}

func AuthenticateUser(username, password string) bool {
	mu.Lock()
	defer mu.Unlock()

	user, exists := users[username]
	if !exists {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
