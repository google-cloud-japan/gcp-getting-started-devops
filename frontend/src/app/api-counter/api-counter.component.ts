/**
 * Copyright 2021 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
