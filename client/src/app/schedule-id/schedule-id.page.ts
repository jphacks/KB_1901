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
            console.log(data);
        }, error => {
            console.log(error);
            alert("何らかのエラーが発生しました。")
        });
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store');
    }

}
