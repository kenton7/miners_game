 package strongminer

// import (
// 	coal_package "coal_mine/coal"
// 	baseminer "coal_mine/miners/base_miner"
// 	minerinfo "coal_mine/miners/miner_info"
// 	"context"
// 	"fmt"
// 	"time"
// )

// type StrongMiner struct {
// 	baseminer.BaseMinerInfo
// }

// func (s *StrongMiner) Run(ctx context.Context, worker int) <-chan coal_package.Coal {
// 	coalChan := make(chan coal_package.Coal, s.Info().GetPower())
// 	ticker := time.NewTicker(time.Duration(s.Info().GetBreakTime()) * time.Second)

// 	fmt.Printf("[%s] Сильный шахтёр начал добывать уголь...\n", s.GetStartTime())
// 	power := s.GetPower()
// 	go func() {
// 		defer ticker.Stop()
// 		defer close(coalChan)

// 		for i := 1; i <= power; i++ {
// 			select {
// 			case <-ctx.Done():
// 				continue
// 			case <-ticker.C:
// 				coalChan <- coal_package.Coal{CraftedCoals: s.Info().GetProfit()}
// 				s.Boost(3)
// 				s.LessPower()
// 			}
// 		}
// 	}()
// 	return coalChan
// }

// func (s *StrongMiner) Info() minerinfo.MinerInfo {
// 	return s
// }

// func New() *StrongMiner {
// 	return &StrongMiner{
// 		BaseMinerInfo: baseminer.NewBaseMinerInfo("Сильный шахтёр", 450, 60, 10, 1),
// 	}
// }
