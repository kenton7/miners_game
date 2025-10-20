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

// func BuyMiner(class baseminer.MinerClass, amount int) {
// 	//reader := bufio.NewReader(os.Stdin)
// 	//var workers int
// 	//var err error
// 	//fmt.Print("Введите количество шахтёров для покупки: ")

// 	for {
// 		//workersStr, _ := reader.ReadString('\n')
// 		//workersStr = strings.TrimSpace(workersStr)
// 		//workers, err = strconv.Atoi(workersStr)
// 		if amount <= 0 {
// 			fmt.Println("❌ Введено неверное количество, попробуйте снова.")
// 			continue
// 		}
// 		break
// 	}

// 	switch class {
// 	case baseminer.LittleMinerClass:
// 		startMiner(func() baseminer.BaseMiner {
// 			return baseminer.NewBaseMinerInfo(baseminer.LittleMinerClass, 5, 30, 1, 3)
// 		}, amount)

// 	case baseminer.NormalMinerClass:
// 		startMiner(func() baseminer.BaseMiner {
// 			return baseminer.NewBaseMinerInfo(baseminer.NormalMinerClass, 50, 100, 5, 2)
// 		}, amount)

// 	case baseminer.StrongMinerClass:
// 		startMiner(func() baseminer.BaseMiner {
// 			return baseminer.NewBaseMinerInfo(baseminer.StrongMinerClass, 450, 300, 20, 1)
// 		}, amount)
// 	default:
// 		fmt.Println("❌ Неверный класс шахтёра.")
// 	}
// }

// func BuyMiner(m Miner, ctx context.Context) {
// 	reader := bufio.NewReader(os.Stdin)
// 	var workersNum int
// 	var err error
// 	fmt.Print("Введите количество шахтёров для покупки: ")

// 	for {
// 		workersStr, _ := reader.ReadString('\n')
// 		workersStr = strings.TrimSpace(workersStr)
// 		workersNum, err = strconv.Atoi(workersStr)
// 		if err != nil || workersNum <= 0 {
// 			fmt.Println("❌ Введено неверное количество, попробуйте снова.")
// 			continue
// 		}
// 		break
// 	}

// 	switch m.(type) {
// 	case *littleminer.LittleMiner:
// 		startMiner(func() Miner {
// 			return littleminer.New()
// 		}, workersNum, ctx)
// 	case *normalminer.NormalMiner:
// 		startMiner(func() Miner {
// 			return normalminer.New()
// 		}, workersNum, ctx)
// 	case *strongminer.StrongMiner:
// 		startMiner(func() Miner {
// 			return strongminer.New()
// 		}, workersNum, ctx)
// 	}
// }

func BuyAndStartMiner(class baseminer.MinerClass, workers int, ctx context.Context) error {
	var wg sync.WaitGroup
	coalChan := make(chan coal_package.Coal, workers)

	if workers <= 0 {
		return errors.New("❌ Введено неверное количество шахтёров")
	}

	// for {
	// 	if workers <= 0 {
	// 		fmt.Println("❌ Введено неверное количество, попробуйте снова.")
	// 		continue
	// 	}
	// 	break
	// }

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

// func startMiner(factory func() baseminer.BaseMiner, workers int) {
// 	var wg sync.WaitGroup
// 	coalChan := make(chan coal_package.Coal, workers)

// 	for i := 1; i <= workers; i++ {
// 		miner := factory() // создаём отдельного шахтёра для каждой горутины
// 		factory_pack.AddMinerToFactory(miner)

// 		if coal_package.GetCurrentBalance() < miner.Info().GetSalary() {
// 			fmt.Printf("❌ Недостаточно средств. Стоимость шахтёра: %d\n", miner.Info().GetSalary())
// 			continue
// 		}

// 		coal_package.PayForWork(miner.Info().GetSalary())
// 		fmt.Printf("✅ Куплен %s #%d! Списано %d. Остаток: %d\n",
// 			miner.Info().GetClass(), i, miner.Info().GetSalary(), coal_package.GetCurrentBalance())

// 		wg.Add(1)
// 		go func(id int, m baseminer.BaseMiner) {
// 			defer wg.Done()
// 			minerChan := m.Run(ctx, id)
// 			for coal := range minerChan {
// 				coalChan <- coal
// 			}
// 		}(i, miner)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(coalChan)
// 	}()

// 	// 🟢 Читаем из канала, иначе всё повиснет
// 	go func() {
// 		for coal := range coalChan {
// 			coal_package.AddToBalance(coal.CraftedCoals)
// 		}
// 	}()
// }
