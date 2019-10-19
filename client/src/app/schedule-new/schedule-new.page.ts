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
    public memo: string;
    public days: any = [];
    public generate_button_flag: boolean = false;
    public key: string;
    private auth_token: string;
    private user_name: string;

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
        this.days.push("");
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.user_name = params.get('user_name');
        });
    }

    incrementDay() {
        this.days.push("");
        console.log(this.days);
    }

    decrementDay() {
        this.days.pop("");
        console.log(this.days);
    }

    trackBy(index: number, obj: any): any {
        return index;
    }

    generateLink() {
        let result = {'plan_name': this.plan_name, 'memo': this.memo, 'day': []};
        //dayの整形
        for (let day of this.days) result.day.push(day.split('T')[0]);
        console.log(this.days);

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
