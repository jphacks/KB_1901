import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {HttpClient} from "@angular/common/http";
import config from "../../../config";

@Component({
    selector: 'app-schedule-id',
    templateUrl: './schedule-id.page.html',
    styleUrls: ['./schedule-id.page.scss'],
})
export class ScheduleIdPage implements OnInit {
    private auth_token: string;
    private key: string;
    public plan_name: string;
    public answer_count;
    public answer_genre;
    public day;
    public area;
    public free;

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.key = params.get('key');
        });

        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_result";
        const formData =
            "plan_key=" + encodeURI(this.key);
        const headers = {
            "headers": {
                "Authorization": this.auth_token,
                "Content-Type": "application/x-www-form-urlencoded"
            }
        };
        this.http.post(url, formData, headers).subscribe(data => {
            let json_data = JSON.parse(data['data'].json);
            console.log(json_data);
            this.plan_name = json_data.plan_name;
            this.answer_count = json_data.answer_count;
            this.day = json_data.day;
            this.area = json_data.area;
            this.answer_genre = json_data.genre;
            this.free = json_data.free;
        }, error => {
            console.log(error);
            alert("何らかのエラーが発生しました。")
        });
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store');
    }

}
