package main

import (
	coal_package "coal_mine/coal"
	factory_pack "coal_mine/factory"
	"coal_mine/menu"
	"coal_mine/server"
	"fmt"
	"time"
)

func main() {

	go func() {
		factory := factory_pack.New()
		httpHanders := server.NewHTTPHandlers(factory)
		httpServer := server.NewHTTPServer(httpHanders)
		if err := httpServer.StartServer(); err != nil {
			fmt.Println("Failed to start HTTP server: ", err)
			return
		}
	}()

	//factoryCtx, factoryStop := context.WithCancel(context.Background())
	startTime := time.Now()
	go menu.ShowMenu()

	go coal_package.IncreaseBalancePerSecond()

	//go menu.StartGame(factoryCtx, factoryStop)

	finishGameChan := factory_pack.IsFinishedGame()
	<-finishGameChan
	//factoryStop()
	fmt.Println("\n🎉 Вы купили все предметы на предприятие и тем самым прошли игру!")
	fmt.Println("⏱ Прохождение заняло:", time.Since(startTime))
	fmt.Println("📊 Статистика за всю игру:")
}
