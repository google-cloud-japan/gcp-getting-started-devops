import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiCounterBenchComponent } from './api-counter-bench.component';

describe('ApiCounterBenchComponent', () => {
  let component: ApiCounterBenchComponent;
  let fixture: ComponentFixture<ApiCounterBenchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApiCounterBenchComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiCounterBenchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
