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
            this.auth_token = params.get('auth_token');
            this.user_name = params.get('user_name');
        });
    }

    goLogin() {
        if (this.user_name !== 'null' && this.auth_token !== 'null') {
            alert("ログアウトしますか？");
        } else {
            this.router.navigateByUrl('/login');
        }
    }

    goScheduleNew() {
        if (this.user_name !== 'null' && this.auth_token !== 'null') {
            this.router.navigateByUrl('/schedule-new/' + this.auth_token + '/' + this.user_name);
        } else {
            alert("ログインしていません。");
        }
    }

    goScheduleList() {
        if (this.user_name !== 'null' && this.auth_token !== 'null') {
            this.router.navigateByUrl('/schedule-list/' + this.auth_token + '/' + this.user_name);
        } else {
            alert("ログインしていません。");
        }
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store/' + this.auth_token + '/' + this.user_name);
    }

}
