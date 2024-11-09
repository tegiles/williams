export interface Race {
    raceID: number;
    circuit: Circuit;
    year: number;
    round: number;
    name: string;
    date: string;
    time: string;
    url: string;
}

export interface Circuit {
    circuitId: number;
    ref: string;
    name: string;
    location: string;
    country: string;
    lat: number;
    lon: number;
    alt: number;
    url: string;
    laps: Lap[];
    races: number;
}

export interface Lap {
    RaceID: number;
    DriverID: number;
    Lap: number;
    Position: number;
    Time: string;
    Milliseconds: number;
    Driver: string;
    Year: number;
}

export interface Driver {
    driverID: number;
    DriverRef: string;
    number: string;
    code: string;
    forename: string;
    surname: string;
    dob: string;
    nationality: string;
    url: string;
    wins: number;
    podiums: number;
}