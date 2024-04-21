import { Component, OnInit } from '@angular/core';
import { DataService } from '../data.service';

@Component({
  selector: 'app-helloworld',
  standalone: true,
  imports: [],
  templateUrl: './helloworld.component.html',
  styleUrls: ['./helloworld.component.css']
})
export class HelloWorldComponent implements OnInit {
  textFromBackend: string = '';

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
    this.dataService.getTextFromBackend().subscribe((data: any) => {
      this.textFromBackend = data;
    });
  }
}
