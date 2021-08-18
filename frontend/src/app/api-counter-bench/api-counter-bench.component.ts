import { Component, OnInit } from '@angular/core';
import { ApiCounterComponent } from '../api-counter/api-counter.component';
import { CountService } from '../count.service';

@Component({
  selector: 'app-api-counter-bench',
  templateUrl: '../api-counter/api-counter.component.html',
  styleUrls: ['../api-counter/api-counter.component.css']
})
export class ApiCounterBenchComponent extends ApiCounterComponent {

  constructor(public countService: CountService) {
    super(countService);
  }

  startCountUp() {
    this.resetCountUp();
    this.accessing = true;
    this.countSubscription = this.countService.beginBenchCountUp()
      .subscribe(value => this.count = value);
  }
}
