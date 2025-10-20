package server

import baseminer "coal_mine/miners/base_miner"

type MinerRequestDTO struct {
	Class string `json:"type"`
}

type MinerResponseDTO struct {
	Class string `json:"type"`
	Cost  int    `json:"cost"`
}

type BuyMinerDTORequest struct {
	Class string `json:"type"`
	Amount int `json:"amount"`
}

type BuyMinerDTOResponse struct {
	Class string `json:"type"`
	Amount int `json:"amount"`
	IsOk bool `json:"isOk"`
	Message string `json:"message"`
}

type ControlGameDTO struct {
	Command int `json:"command"`
}

type BalanceResponseDTO struct {
	Balance int `json:"balance"`
}

type AllMinersInfoDTO struct {
	Miners []baseminer.BaseMiner `json:"miners"`
}