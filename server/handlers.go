package server

import (
	coal_package "coal_mine/coal"
	factory_pack "coal_mine/factory"
	miner "coal_mine/miners"
	baseminer "coal_mine/miners/base_miner"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	factory   *factory_pack.Factory
	ctx       context.Context
	cancelCtx func()
}

func NewHTTPHandlers(factory *factory_pack.Factory) *HTTPHandlers {
	factoryCtx, factoryCancel := context.WithCancel(context.Background())
	return &HTTPHandlers{factory: factory, ctx: factoryCtx, cancelCtx: factoryCancel}
}

func (h *HTTPHandlers) GetMinerSalaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	minerClass := strings.ToLower(mux.Vars(r)["miner"])
	var salary int

	switch minerClass {
	case strings.ToLower(string(baseminer.LittleMinerClass)):
		salary = int(baseminer.LittleMinerSalary)
	case strings.ToLower(string(baseminer.NormalMinerClass)):
		salary = int(baseminer.NormalMinerSalary)
	case strings.ToLower(string(baseminer.StrongMinerClass)):
		salary = int(baseminer.StrongMinerSalary)
	default:
		http.Error(w, "Unknown miner", http.StatusBadRequest)
		return
	}

	response := PriceResponseDTO{
		Class: minerClass,
		Price: salary,
	}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) BuyMiner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request BuyDTORequest
	var response BuyDTOResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Check your request", http.StatusBadRequest)
		return
	}

	if err := miner.BuyAndStartMiner(baseminer.MinerClass(request.Class), request.Quantity, h.ctx); err != nil {
		response = BuyDTOResponse{
			Class:    strings.ToLower(request.Class),
			Quantity: request.Quantity,
			IsOk:     false,
			Message:  err.Error(),
		}

		b, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(b); err != nil {
			http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	response = BuyDTOResponse{
		Class:    request.Class,
		Quantity: request.Quantity,
		IsOk:     true,
		Message:  "Покупка успешна совершена!",
	}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Must be GET method", http.StatusMethodNotAllowed)
		return
	}

	response := BalanceResponseDTO{Balance: coal_package.GetCurrentBalance()}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) GetAllWorkingMiners(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Must be GET method", http.StatusMethodNotAllowed)
		return
	}

	minerClassParam := r.URL.Query().Get("class")
	isWorkingParamStr := r.URL.Query().Get("working")
	var isWorkingParam bool
	var err error

	if isWorkingParamStr != "" {
		isWorkingParam, err = strconv.ParseBool(isWorkingParamStr)
		if err != nil {
			http.Error(w, "Invalid value for working param", http.StatusBadRequest)
			return
		}
	}

	var filtered []baseminer.BaseMiner
	workingMiners := h.factory.GetAllWorkingMiners()

	for _, miner := range workingMiners {
		if minerClassParam != "" && miner.Class != baseminer.MinerClass(minerClassParam) {
			continue
		}

		if isWorkingParamStr != "" && miner.IsWorking != isWorkingParam {
			continue
		}
		filtered = append(filtered, miner)
	}

	response := AllMinersInfoDTO{Miners: filtered}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) GetItemsCost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Must be GET method", http.StatusMethodNotAllowed)
		return
	}

	itemParam := strings.ToLower(mux.Vars(r)["item"])

	response := PriceResponseDTO{
		Class: itemParam,
		Price: factory_pack.GetItemCost(factory_pack.ItemName(itemParam)),
	}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) BuyItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Must be POST method", http.StatusMethodNotAllowed)
		return
	}

	var request BuyDTORequest
	var response BuyDTOResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Check your request", http.StatusBadRequest)
		return
	}

	if err := factory_pack.BuyItem(factory_pack.ItemName(strings.ToLower(request.Class))); err != nil {
		response = BuyDTOResponse{
			Class:    request.Class,
			Quantity: request.Quantity,
			IsOk:     false,
			Message:  "Не хватает средств для покупки предмета",
		}

		b, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
			return
		}
		return
	}

	response = BuyDTOResponse{
		Class:    request.Class,
		Quantity: request.Quantity,
		IsOk:     true,
		Message:  "Предмет успешно куплен",
	}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandlers) CheckItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Must be GET method", http.StatusMethodNotAllowed)
		return
	}

	boughtItems := factory_pack.GetBoughtItems()

	response := BoughtItemsDTO{Items: boughtItems}

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		http.Error(w, "Failed to write HTTP-response", http.StatusInternalServerError)
		return
	}
}

// func (h *HTTPHandlers) ControlGame(w http.ResponseWriter, r *http.Request) {
// 	commandStr := mux.Vars(r)["command"]
// 	command, err := strconv.Atoi(commandStr)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	menu.ControlGame(command, h.ctx, h.cancelCtx)
// }
