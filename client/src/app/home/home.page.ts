import {Component} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";

@Component({
    selector: 'app-home',
    templateUrl: 'home.page.html',
    styleUrls: ['home.page.scss'],
})
export class HomePage {

    private user_name: string;
    private auth_token: string;

    constructor(
        private router: Router,
        private route: ActivatedRoute,
    ) {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('authToken');
            this.user_name = params.get('user_name');
        });
    }

    goLogin() {
        console.log(this.user_name);
        if (this.user_name !== 'null') {
            alert("ログアウトしますか？");
        } else {
            this.router.navigateByUrl('/login');
        }
    }

    goScheduleNew() {
        this.router.navigateByUrl('/schedule-new');
    }

    goScheduleList() {
        this.router.navigateByUrl('/schedule-list');
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store');
    }

}