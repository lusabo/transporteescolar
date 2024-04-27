import { Component, OnInit } from '@angular/core';
import { EstadosService } from '../services/estados.service';
import { CidadesService } from '../services/cidades.service';
import { BairrosService } from '../services/bairros.service';
import { EscolasService } from '../services/escolas.service';
import { MotoristasService } from '../services/motoristas.service';
import { NgFor } from '@angular/common';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-pesquisa-motorista',
  standalone: true,
  imports: [NgIf, NgFor],
  templateUrl: './pesquisa-motorista.component.html',
  styleUrls: ['./pesquisa-motorista.component.css']
})
export class PesquisaMotoristaComponent implements OnInit {
  estados: any[] = [];
  cidades: any[] = [];
  bairros: any[] = [];
  escolas: any[] = [];
  motoristas: any[] = [];
  selectedEstado: string = "";
  selectedCidade: string = "";
  selectedBairro: string = "";
  selectedEscola: string = "";
  

  constructor(
    private estadosService: EstadosService,
    private cidadesService: CidadesService,
    private bairrosService: BairrosService,
    private escolasService: EscolasService,
    private motoristasService: MotoristasService
  ) { }

  ngOnInit(): void {
    this.getEstados();
  }

  getEstados(): void {
    this.estadosService.getEstados()
      .subscribe(estados => this.estados = estados);
  }

  onSelectEstado(event: Event): void {
    this.selectedEstado = (event.target as HTMLInputElement).value;
    this.getCidades(this.selectedEstado);
  }

  onSelectCidade(event: Event): void {
    this.selectedCidade = (event.target as HTMLInputElement).value;
    this.getBairros(this.selectedCidade);
    this.getEscolas(this.selectedCidade);
  }

  onSelectBairro(event: Event): void {
    this.selectedBairro = (event.target as HTMLInputElement).value;
  }

  onSelectEscola(event: Event): void {
    this.selectedEscola = (event.target as HTMLInputElement).value;
  }

  getCidades(estadoId: string): void {
    this.cidadesService.getCidades(estadoId)
      .subscribe(cidades => this.cidades = cidades);
  }

  getBairros(cidadeId: string): void {
    this.bairrosService.getBairros(cidadeId)
      .subscribe(bairros => this.bairros = bairros);
  }

  getEscolas(cidadeId: string): void {
    this.escolasService.getEscolas(cidadeId)
      .subscribe(escolas => this.escolas = escolas);
  }

  pesquisarMotoristas(): void {
    this.motoristasService.pesquisarMotoristas(this.selectedBairro, this.selectedEscola)
      .subscribe(motoristas => this.motoristas = motoristas);
  }
}
