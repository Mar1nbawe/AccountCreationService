package MutexStore

import (
	users "AccountCreationService/Users"
	"github.com/google/uuid"
	"sync"
)

var UserStore = struct {
	sync.RWMutex
	Users map[uuid.UUID]users.User
}{Users: make(map[uuid.UUID]users.User)}
