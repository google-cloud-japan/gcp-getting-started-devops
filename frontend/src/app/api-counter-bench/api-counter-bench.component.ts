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
