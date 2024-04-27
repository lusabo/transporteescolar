import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EscolasService {
  private apiUrl = 'http://localhost:3000/cidades';

  constructor(private http: HttpClient) { }

  getEscolas(cidadeId: string): Observable<any> {
    const url = `${this.apiUrl}/${cidadeId}/escolas`;
    return this.http.get<any>(url);
  }
}
