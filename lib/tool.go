package lib

import (
	"github.com/georgehao/todo/models"
	"time"
	"errors"
	"fmt"
)

func Authenticate(email string, password string) (models.User, error) {
	msg := "invalid email or password"
	user := models.User{}

	if u, err := user.Where("email", email); err != nil {
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

func SignupUser(u *models.User) (uint, error) {
	isExist, err := u.UserExist()
	if err != nil {
		return 0, err
	}

	if isExist {
		msg := u.Email + "was already registered"
		return 0, errors.New(msg)
	}

	u.LastLoginTime = time.Now()
	u.CreatedAt = time.Now()

	fmt.Println(11111);
	err = u.Create()
	if err != nil {
		fmt.Println(1111111111111111)
		return 0, err
	}

	return u.ID, err
}
