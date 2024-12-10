package Users

import (
	"AccountCreationService/Models"
	encrypt "AccountCreationService/encrypt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         uuid.UUID
	Username   string `username:"username" validate:"required"`
	Password   string `password:"password" validate:"required,min=8"`
	Email      string `email:"email" validate:"required"`
	DateJoined time.Time
}

func initUUID(u *User) {
	u.Id = uuid.New()
}

func initTime(u *User) {
	u.DateJoined = time.Now()
}

func initUsername(u *User, inputUsername string) {
	u.Username = inputUsername
}

func initEmail(u *User, inputEmail string, errChan chan error) {
	encryptedEmail, err := encrypt.EncryptData(inputEmail)

	if err != nil {
		errChan <- err
		return
	}
	u.Email = encryptedEmail
	errChan <- nil
}

func initPassword(u *User, inputPassword string, errChan chan error) {
	encryptedPassword, err := encrypt.HashPassword(inputPassword)
	if err != nil {
		errChan <- err
		return
	}
	u.Password = encryptedPassword
	errChan <- nil
}

func InitUser(u *User, input Models.UserInput) error {
	errChan := make(chan error, 2)
	go initEmail(u, input.Email, errChan)
	go initPassword(u, input.Password, errChan)
	initUUID(u)
	initUsername(u, input.Username)
	initTime(u)

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}
