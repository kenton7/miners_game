package miner

import (
	coal_package "coal_mine/coal"
	factory_pack "coal_mine/factory"
	baseminer "coal_mine/miners/base_miner"
	"context"
	"errors"
	"fmt"
	"sync"
)

type Miner interface {
	Run(ctx context.Context, worker int) <-chan coal_package.Coal
	Info() baseminer.MinerClass
}

type MinerInfo interface {
	GetClass() baseminer.MinerClass
	GetSalary() int
	GetPower() int
	GetProfit() int
	GetBreakTime() int
}

func BuyAndStartMiner(class baseminer.MinerClass, workers int, ctx context.Context) error {
	var wg sync.WaitGroup
	coalChan := make(chan coal_package.Coal, workers)

	if workers <= 0 {
		return errors.New("❌ Введено неверное количество шахтёров")
	}

	for i := 1; i <= workers; i++ {
		miner := baseminer.New(class)

		if coal_package.GetCurrentBalance() < miner.GetSalary() {
			fmt.Printf("❌ Недостаточно средств. Стоимость шахтёра: %d\n", miner.GetSalary())
			return errors.New("❌ Недостаточно средств")
		}

		coal_package.PayForWork(miner.GetSalary())
		fmt.Printf("✅ Куплен %s #%d! Списано %d. Остаток: %d\n",
			miner.GetClass(), i, miner.GetSalary(), coal_package.GetCurrentBalance())

		factory_pack.AddMinerToFactory(miner)

		wg.Add(1)
		go func(id int, m *baseminer.BaseMiner) {
			defer wg.Done()
			minerChan := m.Run(ctx, id)
			for coal := range minerChan {
				coalChan <- coal
			}
		}(i, miner)
	}

	go func() {
		go func() {
			wg.Wait()
			close(coalChan)
		}()
		for coal := range coalChan {
			coal_package.AddToBalance(coal.CraftedCoals)
		}
	}()
	return nil
}
