import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Race, Circuit, Driver } from '../shared/types';

@Injectable({
  providedIn: 'root'
})
export class RaceServiceService {

  constructor(private http: HttpClient) { }
  
  getCircuits() {
    return this.http.get<Circuit[]>("/api/circuits")
  }

  getCircuit(id: string) {
    return this.http.get<Circuit>("/api/circuits/"+id)
  }
  
  getDrivers() {
    return this.http.get<Driver[]>("/api/drivers")
  }

}
