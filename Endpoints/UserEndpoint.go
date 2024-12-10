package Endpoints

import (
	"AccountCreationService/Models"
	"AccountCreationService/MutexStore"
	users "AccountCreationService/Users"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()
var validate = validator.New()

func Userhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Warn("Invalid request method")
		return
	}

	var input Models.UserInput
	var newUser users.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		log.Error("Error decoding JSON: ", err)
		return
	}

	if err := validate.Struct(input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Error("Invalid input: ", err)
		return
	}

	emailExists, err := MutexStore.IsEmailinStore(input.Email)
	if err != nil {
		http.Error(w, "Error checking if user exists", http.StatusInternalServerError)
	}
	if emailExists {
		http.Error(w, "User already exists", http.StatusConflict)
		log.Warn("User already exists")
		return
	}

	if err := users.InitUser(&newUser, input); err != nil {
		http.Error(w, "Error initializing user", http.StatusInternalServerError)
		log.Error("Error initializing user: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created: %s", newUser.Username)
	MutexStore.AddUser(newUser)
	log.Info("User created: ", MutexStore.UserStore.Users[newUser.Id].Username)

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Warn("Invalid request method")
		return
	}

}
