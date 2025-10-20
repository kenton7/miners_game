package baseminer

import (
	coal_package "coal_mine/coal"
	"context"
	"fmt"
	"sync"
	"time"
)

var mtx sync.Mutex

type BaseMiner struct {
	Class     MinerClass
	salary    int
	Power     int
	profit    int
	breakTime int
	startAt   string
	IsWorking bool
}

type MinerClass string

const (
	LittleMinerClass MinerClass = "маленький шахтёр"
	NormalMinerClass MinerClass = "обычный шахтёр"
	StrongMinerClass MinerClass = "сильный шахтёр"
)

type MinerSalary int

const (
	LittleMinerSalary MinerSalary = 5
	NormalMinerSalary MinerSalary = 50
	StrongMinerSalary MinerSalary = 450
)

func New(class MinerClass) *BaseMiner {

	var (
		salary    int
		power     int
		profit    int
		breakTime int
	)

	switch class {
	case LittleMinerClass:
		salary = int(LittleMinerSalary)
		power = 30
		profit = 1
		breakTime = 3
	case NormalMinerClass:
		salary = int(NormalMinerSalary)
		power = 45
		profit = 3
		breakTime = 2
	case StrongMinerClass:
		salary = int(StrongMinerSalary)
		power = 60
		profit = 10
		breakTime = 1
	}

	return &BaseMiner{
		Class:     class,
		salary:    salary,
		Power:     power,
		profit:    profit,
		breakTime: breakTime,
		startAt:   time.Now().Format("15:04:05"),
		IsWorking: false,
	}
}

// func NewMiner(class MinerClass, salary, power, profit, breakTime int) BaseMiner {
// 	return BaseMiner{
// 		class:     class,
// 		salary:    salary,
// 		power:     power,
// 		profit:    profit,
// 		breakTime: breakTime,
// 		startAt:   time.Now().Format("15:04:05"),
// 	}
// }

func (b *BaseMiner) Run(ctx context.Context, worker int) <-chan coal_package.Coal {
	coalChan := make(chan coal_package.Coal, b.Power)
	ticker := time.NewTicker(time.Duration(b.breakTime) * time.Second)

	power := b.GetPower()

	fmt.Printf("[%s] %s начал добывать уголь...\n", b.startAt, b.Class)
	go func() {
		defer ticker.Stop()
		defer close(coalChan)

		for i := 1; i <= power; i++ {
			b.IsWorking = true
			select {
			case <-ctx.Done():
				b.IsWorking = false
				continue
			case <-ticker.C:
				coalChan <- coal_package.Coal{CraftedCoals: b.profit}
				coal_package.AddToBalance(b.profit)
				b.LessPower()
				if b.Class == StrongMinerClass {
					b.Boost(3)
				}
			}
		}
		b.IsWorking = false
	}()
	return coalChan
}

func (b *BaseMiner) Info() *BaseMiner {
	return b
}

func (b BaseMiner) GetClass() MinerClass {
	return b.Class
}

func (b BaseMiner) GetSalary() int {
	return b.salary
}

func (b BaseMiner) GetPower() int {
	return b.Power
}

func (b BaseMiner) GetProfit() int {
	return b.profit
}

func (b BaseMiner) GetBreakTime() int {
	return b.breakTime
}

func (b BaseMiner) GetStartTime() string {
	return b.startAt
}

func (b *BaseMiner) Boost(amount int) {
	mtx.Lock()
	b.profit += amount
	mtx.Unlock()
}

func (b *BaseMiner) LessPower() {
	mtx.Lock()
	b.Power -= 1
	mtx.Unlock()
}
