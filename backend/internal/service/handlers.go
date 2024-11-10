package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

func (w *WilliamsService) RootHandler(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("You've reaced the Williams Service"))
}

func (w *WilliamsService) GetDrivers(write http.ResponseWriter, req *http.Request) {
	response := []Driver{}
	for _, driver := range w.DriversMap {
		response = append(response, *driver)
	}

	data, err := json.Marshal(response)
	if err != nil {
		write.WriteHeader(500)
		return
	}
	write.Write(data)
}

func (w *WilliamsService) GetCircuits(write http.ResponseWriter, req *http.Request) {
	response := []CircuitResponse{}
	for _, c := range w.Circuits {
		newCircuit := *c
		newCircuit.Laps = []Lap{}
		response = append(response, newCircuit)
	}
	data, err := json.Marshal(response)
	if err != nil {
		write.WriteHeader(500)
		return
	}
	write.Write(data)
}

func (w *WilliamsService) GetCircuitDetails(write http.ResponseWriter, req *http.Request) {
	circuitIDString := req.PathValue("circuitID")
	circuitID, err := strconv.Atoi(circuitIDString)
	if err != nil {
		fmt.Printf("Error getting circuitID from request: %s", err.Error())
		write.WriteHeader(501)
		return
	}
	for _, circuit := range w.Circuits {
		if circuit.CircuitID == int64(circuitID) {
			response := *circuit
			slices.SortFunc(response.Laps, func(a, b Lap) int {
				return int(a.Milliseconds) - int(b.Milliseconds)
			})

			data, err := json.Marshal(response)
			if err != nil {
				write.WriteHeader(500)
				return
			}
			write.Write(data)
		}
	}
}
