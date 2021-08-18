import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiCounterComponent } from './api-counter.component';

describe('ApiCounterComponent', () => {
  let component: ApiCounterComponent;
  let fixture: ComponentFixture<ApiCounterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApiCounterComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiCounterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
