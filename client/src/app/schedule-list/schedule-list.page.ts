import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import config from "../../../config";
import {HttpClient} from "@angular/common/http";

@Component({
    selector: 'app-schedule-list',
    templateUrl: './schedule-list.page.html',
    styleUrls: ['./schedule-list.page.scss'],
})
export class ScheduleListPage implements OnInit {

    private user_name: string;
    private auth_token: string;
    public plans = [];

    constructor(
        private router: Router,
        private route: ActivatedRoute,
        private http: HttpClient,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.user_name = params.get('user_name');
        });

        const url: string = config.urlScheme + config.host + config.port + "/app/v0/plan_list";
        const formData =
            "account=" + encodeURI(this.user_name);
        const headers = {
            "headers": {
                "Authorization": this.auth_token,
                "Content-Type": "application/x-www-form-urlencoded"
            }
        };
        this.http.post(url, formData, headers).subscribe(data => {
            this.plans = JSON.parse(data["data"].json);
        }, error => {
            console.log(error);
            alert("何らかのエラーが発生しました。")
        });
    }

    goScheduleId(id) {
        this.router.navigateByUrl('/schedule-id/' + this.auth_token + '/' + this.plans[id].key);
    }

}
