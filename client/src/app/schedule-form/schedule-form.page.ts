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

    public title: string;
    public candidates = [];
    public areas = [];
    public genres = [];

    public input_area;
    public input_genre;
    public input_candidate = [];
    public input_free;

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

        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_check";
        const formData = "plan_key=" + encodeURI(this.key);
        const headers = {"headers": {"Content-Type": "application/x-www-form-urlencoded"}};
        this.http.post(url, formData, headers).subscribe(data => {
            let json_data = JSON.parse(data["data"].json);
            this.title = json_data.plan_name;
            for (let day of json_data.day) {
                this.candidates.push(day);
                this.input_candidate.push("");
            }
            for (let area of json_data.area) this.areas.push(area);
            for (let genre of json_data.genre) this.genres.push(genre);

        }, error => {
            console.log(error);
        });
    }


    trackBy(index: number, obj: any): any {
        return index;
    }

    submitData() {
        console.log(this.input_free);
        console.log(this.input_candidate);
        console.log(this.input_genre);
        console.log(this.input_area);

        let result = {
            'area': this.areas[this.input_area],
            'genre': this.genres[this.input_genre],
            'free': this.input_free,
            'select_day': []
        };

        for (let i = 0; i < this.candidates.length; i++) result['select_day'].push({
            'day': this.candidates[i],
            'Check': Number(this.input_candidate[i])
        });

        console.log(result);

        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_form";
        const formData =
            "plan_key=" + encodeURI(this.key) +
            "&form_data=" + encodeURI(JSON.stringify(result));
        const headers = {
            "headers": {
                "Content-Type": "application/x-www-form-urlencoded"
            }
        };
        this.http.post(url, formData, headers).subscribe(data => {
            console.log(data);
            alert("送信が完了しました。");
        }, error => {
            alert("何らかのエラーが発生しました。")
        });


    }
}
