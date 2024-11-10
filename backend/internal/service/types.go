package service

type CircuitResponse struct {
	CircuitID int64   `json:"circuitId"`
	Ref       string  `json:"circuitRef"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	Country   string  `json:"country"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lng"`
	Alt       int     `json:"alt"`
	URL       string  `json:"url"`
	Code      string  `json:"countryCode"`
	Laps      []Lap   `json:"laps"`
	Races     int64   `json:"races"`
}

type Lap struct {
	RaceID       int64   `json:"raceId"`
	DriverID     int64   `json:"driverId"`
	Lap          int64   `json:"lap"`
	Position     int64   `json:"position"`
	Time         string  `json:"time"`
	Milliseconds float64 `json:"milliseconds"`
	Driver       string  `json:"driver"`
	Year         int64   `json:"year"`
}

type Race struct {
	RaceID  int64
	Circuit *CircuitResponse
	Year    int64
	Round   int64
	Name    string
	Date    string
	Time    string
	Url     string
	Laps    []Lap
	MaxLap  int64
}

type Driver struct {
	DriverID    int64  `json:"driverId"`
	DriverRef   string `json:"driverRef"`
	Number      string `json:"number"`
	Code        string `json:"code"`
	Forename    string `json:"forename"`
	Surname     string `json:"surname"`
	DOB         string `json:"dob"`
	Nationality string `json:"nationality"`
	URL         string `json:"url"`
	Wins        int64  `json:"wins"`
	Podiums     int64  `json:"podiums"`
}

// ------------ JSON Input types ----------

type CircuitJSON struct {
	CircuitID int64   `json:"circuitId"`
	Ref       string  `json:"circuitRef"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	Country   string  `json:"country"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lng"`
	Alt       int     `json:"alt"`
	URL       string  `json:"url"`
	Code      string  `json:"countryCode"`
}

type DriverJSON struct {
	DriverID    int64  `json:"driverId"`
	DriverRef   string `json:"driverRef"`
	Number      string `json:"number"`
	Code        string `json:"code"`
	Forename    string `json:"forename"`
	Surname     string `json:"surname"`
	DOB         string `json:"dob"`
	Nationality string `json:"nationality"`
	URL         string `json:"url"`
}

type LapJSON struct {
	RaceID       int64   `json:"raceId"`
	DriverID     int64   `json:"driverId"`
	Lap          int64   `json:"lap"`
	Position     int64   `json:"position"`
	Time         string  `json:"time"`
	Milliseconds float64 `json:"milliseconds"`
}

type RaceJSON struct {
	RaceID     int64  `json:"raceId"`
	Year       int64  `json:"year"`
	Round      int64  `json:"round"`
	CircuitID  int64  `json:"circuitId"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	Url        string `json:"url"`
	Fp1Date    string `json:"fp1_date"`
	Fp1Time    string `json:"fp1_time"`
	Fp2Date    string `json:"fp2_date"`
	Fp2Time    string `json:"fp2_time"`
	Fp3Date    string `json:"fp3_date"`
	Fp3Time    string `json:"fp3_time"`
	QualiDate  string `json:"quali_date"`
	QualiTime  string `json:"quali_time"`
	SprintDate string `json:"sprint_date"`
	SprintTime string `json:"sprint_time"`
	Laps       []Lap
	MaxLap     int64
}
