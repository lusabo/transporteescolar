import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EstadosService {
  private apiUrl = 'http://localhost:3000/estados';

  constructor(private http: HttpClient) { }

  getEstados(): Observable<any> {
    return this.http.get<any>(this.apiUrl);
  }
}
