package MutexStore

import (
	users "AccountCreationService/Users"
	"AccountCreationService/encrypt"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func IsEmailinStore(email string) (bool, error) {
	UserStore.RLock()
	defer UserStore.RUnlock()

	for _, user := range UserStore.Users {
		decryptedEmail, err := encrypt.DecryptData(user.Email)
		if err != nil {
			log.Error("Error decrypting email: ", err)
			return false, err
		}
		log.Infof("Checking email: %s against stored email: %s", email, decryptedEmail)
		if decryptedEmail == email {
			return true, nil
		}
	}
	return false, nil
}

func AddUser(u users.User) {
	UserStore.Lock()
	defer UserStore.Unlock()
	UserStore.Users[u.Id] = u
}
