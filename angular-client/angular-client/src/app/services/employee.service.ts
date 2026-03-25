import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError, Observable, throwError } from 'rxjs';
import { Employee } from '../models/employee';
import { error } from 'console';

@Injectable({
  providedIn: 'root',
})
export class EmployeeService {
  private baseUrl = 'http://localhost:8080';
  constructor(private http: HttpClient) {}
  getEmployees(): Observable<Employee[]> {
    return this.http
      .get<Employee[]>(`${this.baseUrl}/employees`)
      .pipe(catchError(this.handleError));
  }

  private handleError(error: HttpErrorResponse) {
    console.error(error);
    return throwError(() => error.message || 'server error');
  }
}
