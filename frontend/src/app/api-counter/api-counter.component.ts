import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { CountService } from '../count.service';

@Component({
  selector: 'app-api-counter',
  templateUrl: './api-counter.component.html',
  styleUrls: ['./api-counter.component.css']
})
export class ApiCounterComponent implements OnDestroy {
  countSubscription!: Subscription;
  count = 0;
  accessing = false;

  constructor(public countService: CountService) {}

  ngOnDestroy(): void {
    console.log("Destroy is called!");
    this.countSubscription.unsubscribe()
  }

  startCountUp() {
    this.resetCountUp();
    this.accessing = true;
    this.countSubscription = this.countService.beginNormalCountUp()
      .subscribe(value => this.count = value);
  }

  resetCountUp() {
    this.accessing = false;
    this.count = 0;
    if (this.countSubscription) this.countSubscription.unsubscribe();
    this.countService.reset();
  }
}
