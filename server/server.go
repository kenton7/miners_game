package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{httpHandlers: httpHandlers}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/miner/salary/{miner}").Methods("GET").HandlerFunc(s.httpHandlers.GetMinerSalaryHandler)
	router.Path("/buy").Methods("POST").HandlerFunc(s.httpHandlers.BuyMiner)
	//router.Path("/command/{command}").HandlerFunc(s.httpHandlers.ControlGame)
	router.Path("/balance").Methods("GET").HandlerFunc(s.httpHandlers.GetBalance)
	router.Path("/miner").Methods("GET").HandlerFunc(s.httpHandlers.GetAllWorkingMiners)
	router.Path("/item/{item}").Methods("GET").HandlerFunc(s.httpHandlers.GetItemsCost)
	router.Path("/item").Methods("GET").HandlerFunc(s.httpHandlers.CheckItems)
	router.Path("/item").Methods("POST").HandlerFunc(s.httpHandlers.BuyItem)

	if err := http.ListenAndServe(":9092", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
