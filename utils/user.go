package utils

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(u1 *UserBalance, u2 *UserBalance, amount int) {
	u1.Lock()
	fmt.Println("lock", u1.Name)
	u1.Change(amount)

	time.Sleep(1 * time.Second)

	u2.Lock()
	fmt.Println("lock", u2.Name)
	u2.Change(amount)

	u1.Unlock()
	u2.Unlock()
}
