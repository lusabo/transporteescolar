import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { PesquisaMotoristaComponent } from './pesquisa-motorista/pesquisa-motorista.component';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, PesquisaMotoristaComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'frontend';
}
