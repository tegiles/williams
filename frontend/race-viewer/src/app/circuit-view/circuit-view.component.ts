import { Component, OnInit, inject, ViewChild, AfterViewInit } from '@angular/core';
import { RaceServiceService } from '../service/race-service.service';
import { Circuit, Lap } from '../shared/types';
import { MatCardModule } from '@angular/material/card';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { NgClass, NgFor, NgIf, NgStyle } from '@angular/common';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';

@Component({
  selector: 'app-circuit-view',
  standalone: true,
  imports: [MatCardModule, NgFor, NgClass, NgIf, NgStyle, MatGridListModule, MatFormFieldModule, MatInputModule, MatTableModule, MatPaginatorModule],
  templateUrl: './circuit-view.component.html',
  styleUrl: './circuit-view.component.css'
})
export class CircuitViewComponent implements OnInit, AfterViewInit {
  constructor(private raceService: RaceServiceService) { }

  @ViewChild(MatPaginator) paginator: MatPaginator;

  circuits: Circuit[] = []
  selectedCircuit: Circuit = {} as Circuit;
  showDetail: boolean = false;

  circuitDisplayedColumns: string[] = ['name', 'country', 'location'];
  circuitDataSource = new MatTableDataSource(this.circuits);


  lapsDisplayedColumns: string[] = ['driver', 'time', 'year'];
  lapsDataSource = new MatTableDataSource(this.selectedCircuit.laps);

  ngOnInit() {
    this.getCircuits()
  }

  ngAfterViewInit() {
    this.circuitDataSource.paginator = this.paginator;
  }


  getCircuits() {
    this.raceService.getCircuits().subscribe(
      data => {
        this.circuits = data.sort(circuitSorter)
        this.circuitDataSource.data = this.circuits
      }
    )
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.circuitDataSource.filter = filterValue.trim().toLowerCase();

    if (this.circuitDataSource.paginator) {
      this.circuitDataSource.paginator.firstPage();
    }
  }

  tableRowSelected(circuit: Circuit) {
    if (this.selectedCircuit.circuitId === circuit.circuitId) {
      this.showDetail = !this.showDetail
      return
    }
    this.raceService.getCircuit(String(circuit.circuitId)).subscribe(
      data => {
        this.selectedCircuit = data
        if (this.selectedCircuit.laps == null) {
          this.lapsDataSource.data = []
        } else {
          this.lapsDataSource.data = this.selectedCircuit.laps.sort(lapSorter).slice(0,5)
        }
        this.showDetail = true
      }
    )
  }

}

function circuitSorter(a: Circuit, b: Circuit): number {
  return a.country.localeCompare(b.country)
}

function lapSorter(a: Lap, b: Lap): number {
  return a.Milliseconds - b.Milliseconds
}