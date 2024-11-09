package service

import (
	"fmt"
	"net/http"
	"williams/internal/data"
)

type WilliamsService struct {
	http.Server

	Circuits  []*CircuitResponse
	Races     []data.RaceJSON
	Drivers   []Driver
	LapsJSON  []data.LapJSON
	Standings []data.StandingJSON
	Countries []data.CountryJSON

	RacesMap   map[int64]*Race
	DriversMap map[int64]*Driver
	Laps       []Lap
}

func NewService(port int) *WilliamsService {

	svc := &WilliamsService{
		Server: http.Server{
			Addr: fmt.Sprintf("0.0.0.0:%d", port),
		},
		Circuits:  []*CircuitResponse{},
		Races:     []data.RaceJSON{},
		Drivers:   []Driver{},
		LapsJSON:  []data.LapJSON{},
		Standings: []data.StandingJSON{},

		RacesMap:   make(map[int64]*Race),
		DriversMap: make(map[int64]*Driver),
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", svc.RootHandler)
	handler.HandleFunc("/drivers", svc.GetDrivers)
	handler.HandleFunc("/drivers/{driverID}", svc.GetStandingsForDriver)
	handler.HandleFunc("/circuits", svc.GetCircuits)
	handler.HandleFunc("/circuits/{circuitID}", svc.GetCircuitDetails)

	svc.Handler = handler
	return svc
}

func (w *WilliamsService) InitDatasets() {
	var err error
	w.Circuits, err = data.ReadJSON[[]*CircuitResponse]("../dataset/circuits.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Standings, err = data.ReadJSON[[]data.StandingJSON]("../dataset/driver_standings.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Drivers, err = data.ReadJSON[[]Driver]("../dataset/drivers.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.LapsJSON, err = data.ReadJSON[[]data.LapJSON]("../dataset/lap_times.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Races, err = data.ReadJSON[[]data.RaceJSON]("../dataset/races.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.Countries, err = data.ReadJSON[[]data.CountryJSON]("internal/data/country-codes.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	w.ProcessInputData()
}

func (w *WilliamsService) Serve() error {
	return w.ListenAndServe()
}
