import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BairrosService {
  private apiUrl = 'http://localhost:3000/cidades';

  constructor(private http: HttpClient) { }

  getBairros(cidadeId: string): Observable<any> {
    const url = `${this.apiUrl}/${cidadeId}/bairros`;
    return this.http.get<any>(url);
  }
}
