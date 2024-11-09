package service

import (
	"fmt"
)

func (w *WilliamsService) ProcessInputData() {
	for _, lapJSON := range w.LapsJSON {
		w.Laps = append(w.Laps, Lap{
			RaceID:       lapJSON.RaceID,
			DriverID:     lapJSON.DriverID,
			Lap:          lapJSON.Lap,
			Position:     lapJSON.Position,
			Time:         lapJSON.Time,
			Milliseconds: lapJSON.Milliseconds,
			Driver:       w.getDriverName(lapJSON.DriverID),
		})

	}

	for _, raceJSON := range w.Races {
		laps, maxLap := w.getLapsForRace(raceJSON.RaceID, raceJSON.Year)
		w.RacesMap[raceJSON.RaceID] = &Race{
			RaceID:  raceJSON.RaceID,
			Circuit: w.getCircuit(raceJSON.CircuitID),
			Year:    raceJSON.Year,
			Round:   raceJSON.Round,
			Name:    raceJSON.Name,
			Date:    raceJSON.Date,
			Time:    raceJSON.Time,
			Url:     raceJSON.Url,
			Laps:    laps,
			MaxLap:  maxLap,
		}
	}

	for _, driverJSON := range w.Drivers {
		if driverJSON.Number == `\N` {
			driverJSON.Number = "No number assigned"
		}
		w.DriversMap[driverJSON.DriverID] = &Driver{
			DriverID:    driverJSON.DriverID,
			DriverRef:   driverJSON.DriverRef,
			Number:      driverJSON.Number,
			Code:        driverJSON.Code,
			Forename:    driverJSON.Forename,
			Surname:     driverJSON.Surname,
			DOB:         driverJSON.DOB,
			Nationality: driverJSON.Nationality,
			URL:         driverJSON.URL,
		}
	}

	for _, race := range w.RacesMap {
		for _, lap := range race.Laps {
			if lap.Lap == race.MaxLap && lap.Position == 1 {
				w.DriversMap[lap.DriverID].Wins++
				w.DriversMap[lap.DriverID].Podiums++
			} else if lap.Lap == race.MaxLap && lap.Position <= 3 {
				w.DriversMap[lap.DriverID].Podiums++
			}
		}
	}

	for _, circuit := range w.Circuits {
		for _, race := range w.RacesMap {
			if race.Circuit.CircuitID == circuit.CircuitID {
				circuit.Laps = append(circuit.Laps, race.Laps...)
				circuit.Races++
			}
		}
		for _, country := range w.Countries {
			if country.Name == circuit.Country {
				circuit.Code = country.Code
			}
		}
	}
}

func (w *WilliamsService) getCircuit(id int64) *CircuitResponse {
	for _, c := range w.Circuits {
		if c.CircuitID == id {
			return c
		}
	}
	return &CircuitResponse{}
}

func (w *WilliamsService) getLapsForRace(raceID int64, year int64) (laps []Lap, maxLap int64) {
	laps = []Lap{}
	for _, lap := range w.Laps {
		if lap.RaceID == raceID {
			if lap.Lap > maxLap {
				maxLap = lap.Lap
			}
			lap.Year = year
			laps = append(laps, lap)
		}
	}
	return
}

// func getRaceLapsForDriver(raceLaps []Lap, driverID int64) (driverLaps []Lap) {
// 	for _, lap := range raceLaps {
// 		if lap.DriverID == driverID {
// 			driverLaps = append(driverLaps, lap)
// 		}
// 	}
// 	return
// }

func (w *WilliamsService) getDriverName(id int64) string {
	for _, driver := range w.Drivers {
		if driver.DriverID == id {
			return fmt.Sprintf("%s %s", driver.Forename, driver.Surname)
		}
	}
	return ""
}
