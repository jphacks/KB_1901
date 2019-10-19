import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-login',
    templateUrl: './login.page.html',
    styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {

    constructor(
        private router: Router
    ) {
    }

    ngOnInit() {
    }

    goHome() {
        this.router.navigateByUrl('/home');
    }

    goNewAccount() {
        alert("実装中です。");
    }
}
