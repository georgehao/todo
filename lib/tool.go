package lib

import (
	"github.com/georgehao/todo/models"
	"time"
	"errors"
)

func Authenticate(email string, password string) (models.User, error) {
	msg := "invalid email or password"
	user := models.User{Email:email}

	if u, err := user.Where("email", user.Email); err != nil {
		return user, errors.New(msg)
	} else if u.ID < 1 {
		return user, errors.New(msg)
	} else if password != u.Password {
		return user, errors.New(msg)
	} else {
		u.LastLoginTime = time.Now()
		u.Save()
		return u, nil
	}
}
