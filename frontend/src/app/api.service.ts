import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Elapsed } from './elapsed';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {}

  accessNormalEndpoint(): Observable<HttpResponse<Elapsed>> {
    return this.http.get<Elapsed>('normal', {observe: 'response'});
  }

  accessBenchEndpoint(): Observable<HttpResponse<Elapsed>> {
    return this.http.get<Elapsed>('bench', {observe: 'response'});
  }
}
