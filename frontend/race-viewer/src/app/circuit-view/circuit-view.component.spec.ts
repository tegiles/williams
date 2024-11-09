import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CircuitViewComponent } from './circuit-view.component';

describe('CircuitViewComponent', () => {
  let component: CircuitViewComponent;
  let fixture: ComponentFixture<CircuitViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CircuitViewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CircuitViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
