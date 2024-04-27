import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CidadesService {
  private apiUrl = 'http://localhost:3000/estados';

  constructor(private http: HttpClient) { }

  getCidades(estadoId: string): Observable<any> {
    const url = `${this.apiUrl}/${estadoId}/cidades`;
    return this.http.get<any>(url);
  }
}
