import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {HttpClient} from "@angular/common/http";
import config from "../../../config";

@Component({
    selector: 'app-schedule-new',
    templateUrl: './schedule-new.page.html',
    styleUrls: ['./schedule-new.page.scss'],
})

export class ScheduleNewPage implements OnInit {

    public plan_name: string;
    public input_day: string;
    public input_area: string;
    public input_genre: string;

    public memo: string;
    public days: any = [];
    public areas: any = [];
    public genres: any = [];

    public generate_button_flag: boolean = false;
    public key: string;
    private auth_token: string;
    private user_name: string;

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.user_name = params.get('user_name');
        });
    }

    addArea() {
        this.areas.push(this.input_area);
    }

    addDay() {
        this.days.push(this.input_day);
    }

    addGenre() {
        this.genres.push(this.input_genre);
    }

    trackBy(index: number, obj: any): any {
        return index;
    }


    //TODO:複数回答強制

    generateLink() {
        if (!this.plan_name) {
            alert("件名が空欄です。");
            return;
        }
        let result = {'plan_name': this.plan_name, 'memo': this.memo, 'day': [], 'area': [], 'genre': []};


        if (this.days.length == 0) {
            alert("入力されていない日にちがあります。");
            return;
        } else {
            for (let day of this.days) result.day.push(day.split('T')[0]);
        }

        if (this.areas.length == 0) {
            alert("入力されていないエリアがあります。");
            return;
        } else {
            for (let area of this.areas) result.area.push(area);
        }

        if (this.genres.length == 0) {
            alert("入力されていないジャンルがあります。");
            return;
        } else {
            for (let genre of this.genres) result.genre.push(genre);
        }

        //dayの整形
        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_generate";
        const formData =
            "account=" + encodeURI(this.user_name) +
            "&plan_data=" + encodeURI(JSON.stringify(result));
        const headers = {
            "headers": {
                "Authorization": this.auth_token,
                "Content-Type": "application/x-www-form-urlencoded"
            }
        };
        console.log(headers);
        this.http.post(url, formData, headers).subscribe(data => {
            console.log(data);
            this.generate_button_flag = true;
            this.key = data["data"].key;
        }, error => {
            alert("何らかのエラーが発生しました。")
        });
    }

    goHome() {
        this.router.navigateByUrl('/home/' + this.auth_token + '/' + this.user_name);
    }
}
