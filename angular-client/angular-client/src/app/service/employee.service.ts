import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { catchError, Observable, throwError } from 'rxjs';
import { Employee } from '../models/employee';
import { ReturnStatement } from '@angular/compiler';

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

  addEmployee(employee: Employee): Observable<Employee> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-type': 'application/json',
      }),
    };
    return this.http.post<Employee>(`${this.baseUrl}/employees/add`, employee, httpOptions);
  }

  private handleError(error: HttpErrorResponse) {
    console.error(error);
    return throwError(() => error.message || 'server error');
  }
}
