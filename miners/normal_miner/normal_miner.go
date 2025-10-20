package normalminer

// import (
// 	coal_package "coal_mine/coal"
// 	baseminer "coal_mine/miners/base_miner"
// 	minerinfo "coal_mine/miners/miner_info"
// 	"context"
// 	"fmt"
// 	"time"
// )

// type NormalMiner struct {
// 	baseminer.BaseMinerInfo
// }

// func (n *NormalMiner) Run(ctx context.Context, worker int) <-chan coal_package.Coal {
// 	coalChan := make(chan coal_package.Coal, n.Info().GetPower())
// 	ticker := time.NewTicker(time.Duration(n.Info().GetBreakTime()) * time.Second)

// 	fmt.Printf("[%s] Средний шахтёр начал добывать уголь...\n", n.GetStartTime())
// 	power := n.GetPower()
// 	go func() {
// 		defer ticker.Stop()
// 		defer close(coalChan)

// 		for i := 1; i <= power; i++ {
// 			select {
// 			case <-ctx.Done():
// 				continue
// 			case <-ticker.C:
// 				coalChan <- coal_package.Coal{CraftedCoals: n.Info().GetProfit()}
// 				coal_package.AddToBalance(n.Info().GetProfit())
// 				n.LessPower()
// 			}
// 		}
// 	}()
// 	return coalChan
// }

// func (n *NormalMiner) Info() minerinfo.MinerInfo {
// 	return n
// }

// func New() *NormalMiner {
// 	return &NormalMiner{
// 		BaseMinerInfo: baseminer.NewBaseMinerInfo("Средний шахтёр", 50, 45, 3, 2),
// 	}
// }
