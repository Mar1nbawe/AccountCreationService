package main

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id         uuid.UUID
	username   string
	password   string
	email      string
	dateJoined time.Time
}

func initUUID(u *User) {
	u.id = uuid.New()
}

func initTime(u *User) {
	u.dateJoined = time.Now()
}
