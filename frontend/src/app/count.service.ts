import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { ApiService } from './api.service';
import { catchError, concatMap, delay, map, repeat } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class CountService {
  count: number = 0;

  constructor(private apiService: ApiService) { }

  reset() {
    this.count = 0;
  }

  beginNormalCountUp(): Observable<number> {
    return of(this.count).pipe(
      concatMap(() => this.apiService.accessNormalEndpoint()),
      map(res => {
        if (res.status === 200) this.count++;
        return this.count;
      }),
      catchError((err) => {
        return of(this.count);
      }
      ),
      delay(1000),
      repeat(),
    );
  }

  beginBenchCountUp(): Observable<number> {
    return of(this.count).pipe(
      concatMap(() => this.apiService.accessBenchEndpoint()),
      map(res => {
        if (res.status === 200) this.count++;
        return this.count;
      }),
      catchError((err) => {
        return of(this.count);
      }
      ),
      delay(2000),
      repeat(),
    );
  }
}
