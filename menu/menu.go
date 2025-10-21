package menu

import (
	"bufio"
	coal_package "coal_mine/coal"
	factory_pack "coal_mine/factory"
	miner "coal_mine/miners"
	baseminer "coal_mine/miners/base_miner"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println()
	fmt.Println("--- МЕНЮ --- ")
	fmt.Println("Введите соответствующую цифру для выполнения комады:")
	fmt.Println("1 - Проверить баланс")
	fmt.Println("2 - Нанять маленького шахтёра за 5 угля")
	fmt.Println("3 - Нанять среднего шахтёра за 50 угля")
	fmt.Println("4 - Нанять сильного шахтёра за 450 угля")
	fmt.Println("5 - Получить информацию о шахтёрах")
	fmt.Println("6 - Купить кирку на предприятие за 3000 угля")
	fmt.Println("7 - Купить вентиляцию в шахту за 15000 угля")
	fmt.Println("8 - Купить вагонетки за 50000 угля")
	fmt.Println("9 - Завершить работу всех шахтёров")
	fmt.Println("10 - Показать меню")
	fmt.Println("0 - Выйти из игры")
}

func ControlGame(command int, ctx context.Context, stopMiners func()) {

	//factoryCtx, factoryStop := context.WithCancel(context.Background())

	switch command {
	case 1:
		//TODO: - Сделать запрос баланса по API
		fmt.Printf("💰 Ваш текущий баланс: %d\n", coal_package.GetCurrentBalance())

	case 2:
		workers := checkWorkers()
		miner.BuyAndStartMiner(baseminer.LittleMinerClass, workers, ctx)
	case 3:
		workers := checkWorkers()
		miner.BuyAndStartMiner(baseminer.NormalMinerClass, workers, ctx)
	case 4:
		workers := checkWorkers()
		miner.BuyAndStartMiner(baseminer.StrongMinerClass, workers, ctx)
	case 5:
		factory_pack.GetInfoAboutMiners()
	case 6:
		if err := factory_pack.BuyItem(factory_pack.Pickaxe); err != nil {
			fmt.Println(err.Error())
		}
		// if err := factory_pack.BuyItem(factory_pack.Pickaxe, factory_pack.PickaxeCost); err != nil {
		// 	fmt.Println(err.Error())
		// }
		// if err := factory_pack.BuyItem(factory_pack.NewItem("Кирка", 3000)); err != nil {
		// 	fmt.Println(err.Error())
		// }
	case 7:
		if err := factory_pack.BuyItem(factory_pack.Ventilation); err != nil {
			fmt.Println(err.Error())
		}
		// if err := factory_pack.BuyItem(factory_pack.Ventilation, factory_pack.VentilationCost); err != nil {
		// 	fmt.Println(err.Error())
		// }
	case 8:
		if err := factory_pack.BuyItem(factory_pack.Wagon); err != nil {
			fmt.Println(err.Error())
		}
		// if err := factory_pack.BuyItem(factory_pack.Wagon, factory_pack.WagonCost); err != nil {
		// 	fmt.Println(err.Error())
		// }
	case 9:
		fmt.Println("Предприятие решило завершить работу всех шахтёров...")
		stopMiners()
		return
	case 10:
		ShowMenu()
	case 0:
		fmt.Println("Выход из игры...")
		return
	default:
		fmt.Println("Неизвестная команда, попробуйте снова.")
	}
}

// func ControlGame(ctx context.Context, stopMiners func()) {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("\nВведите команду: ")
// 		commandStr, _ := reader.ReadString('\n')
// 		commandStr = strings.TrimSpace(commandStr)
// 		command, err := strconv.Atoi(commandStr)
// 		if err != nil {
// 			fmt.Println("❌ Введена неверная команда.")
// 			continue
// 		}

// 		switch command {
// 		case 1:
// 			//TODO: - Сделать запрос баланса по API
// 			fmt.Printf("💰 Ваш текущий баланс: %d\n", coal_package.GetCurrentBalance())
// 		case 2:
// 			workers := checkWorkers()
// 			miner.BuyAndStartMiner(baseminer.LittleMinerClass, workers, ctx)
// 		case 3:
// 			workers := checkWorkers()
// 			miner.BuyAndStartMiner(baseminer.NormalMinerClass, workers, ctx)
// 		case 4:
// 			workers := checkWorkers()
// 			miner.BuyAndStartMiner(baseminer.StrongMinerClass, workers, ctx)
// 		case 5:
// 			factory_pack.GetInfoAboutMiners()
// 		case 6:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Кирка", 3000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 7:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Вентиляция", 15000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 8:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Вагонетка", 50000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 9:
// 			fmt.Println("Предприятие решило завершить работу всех шахтёров...")
// 			stopMiners()
// 			return
// 		case 10:
// 			showMenu()
// 		case 0:
// 			fmt.Println("Выход из игры...")
// 			return
// 		default:
// 			fmt.Println("Неизвестная команда, попробуйте снова.")
// 		}
// 	}
// }

// func StartGame(ctx context.Context, stopMiners func()) {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("\nВведите команду: ")
// 		commandStr, _ := reader.ReadString('\n')
// 		commandStr = strings.TrimSpace(commandStr)
// 		command, err := strconv.Atoi(commandStr)
// 		if err != nil {
// 			fmt.Println("❌ Введена неверная команда.")
// 			continue
// 		}

// 		switch command {
// 		case 1:
// 			fmt.Printf("💰 Ваш текущий баланс: %d\n", coal_package.GetCurrentBalance())
// 		case 2:
// 			//miner.BuyMiner(baseminer.LittleMinerClass, ctx)
// 			//miner.BuyAndStartMiner(baseminer.LittleMinerClass, 2, ctx)
// 			//miner.BuyMiner(littleminer.New(), ctx)
// 		case 3:
// 			//miner.BuyMiner(normalminer.New(), ctx)
// 			miner.BuyMiner(baseminer.NormalMinerClass, ctx)
// 		case 4:
// 			//miner.BuyMiner(strongminer.New(), ctx)
// 			miner.BuyMiner(baseminer.StrongMinerClass, ctx)
// 		case 5:
// 			factory_pack.GetInfoAboutMiners()
// 		case 6:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Кирка", 3000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 7:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Вентиляция", 15000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 8:
// 			if err := factory_pack.BuyItem(factory_pack.NewItem("Вагонетка", 50000)); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		case 9:
// 			fmt.Println("Предприятие решило завершить работу всех шахтёров...")
// 			stopMiners()
// 			return
// 		case 10:
// 			showMenu()
// 		case 0:
// 			fmt.Println("Выход из игры...")
// 			return
// 		default:
// 			fmt.Println("Неизвестная команда, попробуйте снова.")
// 		}
// 	}
// }

func checkWorkers() int {
	reader := bufio.NewReader(os.Stdin)
	var workers int
	var err error
	fmt.Print("Введите количество шахтёров для покупки: ")

	for {
		workersStr, _ := reader.ReadString('\n')
		workersStr = strings.TrimSpace(workersStr)
		workers, err = strconv.Atoi(workersStr)
		if err != nil || workers <= 0 {
			fmt.Println("❌ Введено неверное количество, попробуйте снова.")
			continue
		}
		break
	}
	return workers
}
