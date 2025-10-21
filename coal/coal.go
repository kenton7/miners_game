package coal_package

import (
	"sync"
	"time"
)

type Coal struct {
	CraftedCoals int
}

var balance = 1
var mtx sync.RWMutex

func GetCurrentBalance() int {
	mtx.RLock()
	defer mtx.RUnlock()
	return balance
}

func IncreaseBalancePerSecond() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			mtx.Lock()
			balance += 1
			mtx.Unlock()
		}
	}()
}

func AddToBalance(amount int) {
	mtx.Lock()
	defer mtx.Unlock()
	balance += amount
}

func PayForWork(cost int) bool {
	mtx.Lock()
	defer mtx.Unlock()
	if balance >= cost {
		balance -= cost
		return true
	}
	return false
}
