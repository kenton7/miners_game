package factory_pack

import (
	coal_package "coal_mine/coal"
	baseminer "coal_mine/miners/base_miner"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Name     ItemName
	Price    ItemsPrice
	BoughtAt string
}

type ItemName string

const (
	Pickaxe     ItemName = "кирка"
	Ventilation ItemName = "вентиляция"
	Wagon       ItemName = "вагонетка"
)

type ItemsPrice int

const (
	PickaxeCost     ItemsPrice = 3000
	VentilationCost ItemsPrice = 15000
	WagonCost       ItemsPrice = 50000
)

var miners []*baseminer.BaseMiner
var items []Item
var mtx sync.Mutex

type Factory struct {
	miners []*baseminer.BaseMiner
	items  []Item
	Miner  baseminer.BaseMiner
	mtx    sync.RWMutex
}

func New() *Factory {
	return &Factory{
		miners: make([]*baseminer.BaseMiner, 0),
		items:  make([]Item, 0),
	}
}

func (f *Factory) GetAllWorkingMiners() []baseminer.BaseMiner {
	var workingMiners []baseminer.BaseMiner

	for _, miner := range miners {
		if miner.IsWorking {
			workingMiners = append(workingMiners, *miner)
		}
	}
	return workingMiners
}

func AddMinerToFactory(miner *baseminer.BaseMiner) {
	mtx.Lock()
	defer mtx.Unlock()
	miners = append(miners, miner)
}

func GetInfoAboutMiners() {
	fmt.Println("Информация о шахтёрах на предприятии: ")
	for _, minerInfo := range miners {
		fmt.Printf("Шахтёр: %s. У него осталось %d энергии\n", minerInfo.GetClass(), minerInfo.GetPower())
	}
	fmt.Printf("Всего предприятием было нанято шахтёров: %d.\n", len(miners))
}

func newItem(name ItemName) Item {

	var price ItemsPrice

	switch name {
	case Pickaxe:
		price = PickaxeCost
	case Ventilation:
		price = VentilationCost
	case Wagon:
		price = WagonCost
	}

	return Item{
		Name:     name,
		Price:    price,
		BoughtAt: time.Now().Format("15:04:05"),
	}
}

func BuyItem(name ItemName) error {
	item := newItem(name)
	if coal_package.GetCurrentBalance() >= int(item.Price) {
		mtx.Lock()
		coal_package.PayForWork(int(item.Price))
		items = append(items, item)
		fmt.Printf("✅ [%s] Предмет %s куплен!\n", item.BoughtAt, item.Name)
		mtx.Unlock()
		return nil
	}
	return errors.New("Не хвататет денег для покупки")
}

// func BuyItem(item Item) error {
// 	if coal_package.GetCurrentBalance() >= item.Amount {
// 		mtx.Lock()
// 		coal_package.PayForWork(item.Amount)
// 		items = append(items, item)
// 		mtx.Unlock()
// 		fmt.Printf("✅ [%s] Предмет %s куплен!\n", item.BoughtAt, item.Name)
// 		return nil
// 	}
// 	return errors.New("Не хвататет денег для покупки")
// }

func GetItemCost(item ItemName) int {
	switch item {
	case Pickaxe:
		return int(PickaxeCost)
	case Ventilation:
		return int(VentilationCost)
	case Wagon:
		return int(WagonCost)
	default:
		return 0
	}
}

func GetBoughtItems() []Item {
	var boughtItems []Item

	for _, item := range items {
		if item.BoughtAt != "" {
			boughtItems = append(boughtItems, item)
		}
	}
	
	return boughtItems
}

func IsFinishedGame() <-chan struct{} {
	finishGame := make(chan struct{})

	go func() {
		for {
			hasPickaxe := false
			hasVentilation := false
			hasWagon := false

			for _, item := range items {
				switch item.Name {
				case "Кирка":
					hasPickaxe = true
				case "Вентиляция":
					hasVentilation = true
				case "Вагонетка":
					hasWagon = true
				}
			}
			if hasPickaxe && hasVentilation && hasWagon {
				close(finishGame)
				return
			}
			time.Sleep(time.Second)
		}
	}()
	return finishGame
}

func GetStats() {
	GetInfoAboutMiners()
	for _, item := range items {
		fmt.Printf("Оборудование %s было куплено в: %s", item.Name, item.BoughtAt)
	}
}
