import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {HttpClient} from "@angular/common/http";
import config from "../../../config";

@Component({
    selector: 'app-login',
    templateUrl: './login.page.html',
    styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {

    public user_name: string;
    public password: string;

    constructor(
        private router: Router,
        private http: HttpClient,
    ) {
    }

    ngOnInit() {
    }

    goHome() {
        const url: string = config.urlScheme + config.host + config.port + "/app/v0/login";
        const formData =
            "account=" + encodeURI(this.user_name) +
            "&password=" + encodeURI(this.password);
        const headers = {"headers": {"Content-Type": "application/x-www-form-urlencoded"}};
        this.http.post(url, formData, headers).subscribe(data => {
            console.log(data);
            this.router.navigateByUrl('/home/' + data["data"].token + '/' + this.user_name);
        }, error => {
            alert("ユーザー名かパスワードが間違っています。再度入力してください。")
        });
    }

    goNewAccount() {
        alert("実装中です。");
    }
}
