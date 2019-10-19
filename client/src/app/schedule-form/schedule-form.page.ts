import {Component, OnInit} from '@angular/core';
import config from "../../../config";
import {HttpClient} from "@angular/common/http";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";

@Component({
    selector: 'app-schedule-form',
    templateUrl: './schedule-form.page.html',
    styleUrls: ['./schedule-form.page.scss'],
})

export class ScheduleFormPage implements OnInit {

    private key: string;

    title: string = '予定';
    candidates: { day: string }[] = [
        {day: '10/20'},
        {day: '10/21'},
        {day: '10/22'},
    ];

    select_day: { day: string, check: string }[] = [
        {day: '10/20', check: ""},
        {day: '10/21', check: ""},
        {day: '10/22', check: ""},
    ];
    area: string = "";
    genre: string = "";
    free: string = "";

    constructor(
        private http: HttpClient,
        private router: Router,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.key = params.get('key');
        });
        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_form";
        const formData = "key=" + encodeURI(this.key);
        const headers = {"headers": {"Content-Type": "application/x-www-form-urlencoded"}};
        this.http.post(url, formData, headers).subscribe(data => {
            console.log(data);
        }, error => {
            console.log(error);
        });
    }

    segmentChanged(ev: any, day: string) {
        console.log('Segment changed', day, ev.target.value);
        this.select_day.find((sd) => sd.day == day).check = ev.target.value;
        console.log(this.select_day);
    }

    handleAreaSelect(ev: any) {
        this.area = ev.detail.value;
        console.log(this.area);
    }

    handleGenreSelect(ev: any) {
        this.genre = ev.detail.value;
        console.log(this.area);
    }

    handleFreeWrite(ev: any) {
        this.free = ev.detail.value;
        console.log(this.area);
    }

    submitData() {
        let data = {
            "area": this.area,
            "genre": this.genre,
            "free": this.free,
            "select_day": this.select_day,
        };

        console.log(data);
    }
}
