import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MotoristasService {
  private apiUrl = 'http://localhost:3000/motoristas';

  constructor(private http: HttpClient) { }

  pesquisarMotoristas(bairroId: string, escolaId: string): Observable<any> {
    const url = `${this.apiUrl}?bairro=${bairroId}&escola=${escolaId}`;
    return this.http.get<any>(url);
  }
}
