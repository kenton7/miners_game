package server

import (
	factory_pack "coal_mine/factory"
	baseminer "coal_mine/miners/base_miner"
)

type CostRequestDTO struct {
	Class string `json:"type"`
}

type PriceResponseDTO struct {
	Class string `json:"type"`
	Price int    `json:"price"`
}

type BuyDTORequest struct {
	Class    string `json:"type"`
	Quantity int    `json:"quantity"`
}

type BuyDTOResponse struct {
	Class    string `json:"type"`
	Quantity int    `json:"quantity"`
	IsOk     bool   `json:"isOk"`
	Message  string `json:"message"`
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

type TotalMinersDTO struct {
	TotalMiners int `json:"total_miners"`
}

type BoughtItemsDTO struct {
	Items []factory_pack.Item `json:"items"`
}
