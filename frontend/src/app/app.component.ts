import { Component, ViewChild } from '@angular/core';
import { MatTabChangeEvent } from '@angular/material/tabs';
import { ApiCounterBenchComponent } from './api-counter-bench/api-counter-bench.component';
import { ApiCounterComponent } from './api-counter/api-counter.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  @ViewChild(ApiCounterComponent, {static: true}) apiCounter?: ApiCounterComponent;
  @ViewChild(ApiCounterBenchComponent, {static: true}) apiCounterBench?: ApiCounterBenchComponent;

  onTabChanged(tabChangeEvent: MatTabChangeEvent): void {
    if (tabChangeEvent.index === 0) {
      this.apiCounterBench?.resetCountUp();
    } else {
      this.apiCounter?.resetCountUp();
    }
  }
}
