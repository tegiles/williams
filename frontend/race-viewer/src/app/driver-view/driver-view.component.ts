import { Component, OnInit, inject, ViewChild, AfterViewInit } from '@angular/core';
import { RaceServiceService } from '../service/race-service.service';
import { Driver } from '../shared/types';
import { MatCardModule } from '@angular/material/card';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { formatDate, NgClass, NgFor, NgIf, NgStyle } from '@angular/common';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';

@Component({
  selector: 'app-driver-view',
  standalone: true,
  imports: [MatCardModule, NgFor, NgClass, NgIf, NgStyle, MatGridListModule, MatFormFieldModule, MatInputModule, MatTableModule, MatPaginatorModule],
  templateUrl: './driver-view.component.html',
  styleUrl: './driver-view.component.css'
})
export class DriverViewComponent implements OnInit, AfterViewInit {
  constructor(private raceService: RaceServiceService) { }

  @ViewChild(MatPaginator) paginator: MatPaginator;

  showDetail: boolean = false;

  drivers: Driver[] = []
  selectedDriver: Driver = {} as Driver;

  driverDisplayedColumns: string[] = ['name', 'nationality'];
  driverDataSource = new MatTableDataSource(this.drivers);

  ngOnInit() {
    this.getDrivers()
  }

  ngAfterViewInit() {
    this.driverDataSource.paginator = this.paginator;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.driverDataSource.filter = filterValue.trim().toLowerCase();

    if (this.driverDataSource.paginator) {
      this.driverDataSource.paginator.firstPage();
    }
  }

  getDrivers() {
    this.raceService.getDrivers().subscribe(
      data => {
        this.drivers = data.sort(driverSorter)
        this.driverDataSource.data = this.drivers
      }
    )
  }

  tableRowSelected(driver: Driver) {
    if (this.selectedDriver.driverID != undefined && this.selectedDriver.driverID === driver.driverID) {
      this.showDetail = !this.showDetail
      return
    }
    this.selectedDriver = driver
    this.showDetail = true
  }

  formattedDate(date: string): string {
    const format = 'dd/MM/yyyy';
    const locale = 'en-US';

    return formatDate(date, format, locale)
  }
}

function driverSorter(a: Driver, b : Driver): number {
  return a.surname.localeCompare(b.surname)
}