package service

import (
	"fmt"
	"net/http"
)

type WilliamsService struct {
	http.Server

	Circuits []*CircuitResponse
	Races    []RaceJSON
	Drivers  []*Driver
	LapsJSON []*Lap

	// Maps used for faster processing
	RacesMap   map[int64]*RaceJSON
	DriversMap map[int64]*Driver
}

// NewService returns a new service, ready to serve on the given port
func NewService(port int) *WilliamsService {

	svc := &WilliamsService{
		Server: http.Server{
			Addr: fmt.Sprintf("0.0.0.0:%d", port),
		},
		Circuits: []*CircuitResponse{},
		Races:    []RaceJSON{},
		Drivers:  []*Driver{},
		LapsJSON: []*Lap{},

		RacesMap:   make(map[int64]*RaceJSON),
		DriversMap: make(map[int64]*Driver),
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", svc.RootHandler)
	handler.HandleFunc("/drivers", svc.GetDrivers)
	handler.HandleFunc("/circuits", svc.GetCircuits)
	handler.HandleFunc("/circuits/{circuitID}", svc.GetCircuitDetails)

	svc.Handler = handler
	return svc
}

// InitDatasets reads in the datasets
func (w *WilliamsService) InitDatasets() {
	var err error
	w.Circuits, err = ReadJSON[[]*CircuitResponse]("../dataset/circuits.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Drivers, err = ReadJSON[[]*Driver]("../dataset/drivers.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.LapsJSON, err = ReadJSON[[]*Lap]("../dataset/lap_times.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Races, err = ReadJSON[[]RaceJSON]("../dataset/races.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.ProcessInputData()
}

// Serve starts the HTTP server
func (w *WilliamsService) Serve() error {
	return w.ListenAndServe()
}
