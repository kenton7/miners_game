package littleminer

// func (l *LittleMiner) Run(ctx context.Context, worker int) <-chan coal_package.Coal {
// 	coalChan := make(chan coal_package.Coal, l.GetPower())
// 	ticker := time.NewTicker(time.Duration(l.GetBreakTime()) * time.Second)

// 	fmt.Printf("[%s] Маленький шахтёр начал добывать уголь...\n", l.GetStartTime())
// 	power := l.GetPower()
// 	go func() {
// 		defer ticker.Stop()
// 		defer close(coalChan)

// 		for i := 1; i <= power; i++ {
// 			select {
// 			case <-ctx.Done():
// 				continue
// 			case <-ticker.C:
// 				coalChan <- coal_package.Coal{CraftedCoals: l.GetProfit()}
// 				coal_package.AddToBalance(l.GetProfit())
// 				l.LessPower()
// 			}
// 		}
// 	}()
// 	return coalChan
// }

// func (l *LittleMiner) Info() minerinfo.MinerInfo {
// 	return l
// }

// func New() *LittleMiner {
// 	return &LittleMiner{
// 		BaseMinerInfo: baseminer.NewBaseMinerInfo("Маленький шахтёр", 5, 30, 1, 3),
// 	}
// }
