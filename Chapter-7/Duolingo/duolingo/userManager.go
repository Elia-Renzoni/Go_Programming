package duolingo

import (
	"math/rand"
	"time"
)

type UserManager struct {
	users map[int][]string
}

func (m *UserManager) addUser(values ...string) {
	m.users = make(map[int][]string)
	rand.Seed(time.Now().Unix())
	m.users[rand.Intn(100)] = values
}
