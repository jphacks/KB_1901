import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-schedule-form',
  templateUrl: './schedule-form.page.html',
  styleUrls: ['./schedule-form.page.scss'],
})
export class ScheduleFormPage implements OnInit {
  title: string = '予定';
  candidates: { day: string }[] = [
    { day: '10/20' },
    { day: '10/21' },
    { day: '10/22' },
  ];

  select_day: {day: string, check: string}[] = [
    { day: '10/20' , check: ""},
    { day: '10/21' , check: ""},
    { day: '10/22' , check: ""},
  ];
  area: string = "";
  genre: string = "";
  free: string = "";

  constructor() { }

  ngOnInit() {
  }

  segmentChanged(ev: any, day: string) {
    console.log('Segment changed', day, ev.target.value);
    this.select_day.find((sd) => sd.day == day).check = ev.target.value;
    
    console.log(this.select_day);
  }

  handleAreaSelect(ev: any){
    this.area = ev.detail.value;
    console.log(this.area);
  }

  handleGenreSelect(ev: any){
    this.genre = ev.detail.value;
    console.log(this.area);
  }

  handleFreeWrite(ev: any){
    this.free = ev.detail.value;
    console.log(this.area);
  }

  submitData(){
    let data = {
      "area": this.area,
      "genre": this.genre,
      "free": this.free,
      "select_day": this.select_day,
    };

    console.log(data);
  }
}
