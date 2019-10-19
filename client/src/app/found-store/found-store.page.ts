import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-found-store',
    templateUrl: './found-store.page.html',
    styleUrls: ['./found-store.page.scss'],
})
export class FoundStorePage implements OnInit {

    constructor(
        private router: Router,
    ) {
    }

    ngOnInit() {
    }

    goResult() {
        this.router.navigateByUrl('/result');
    }

}
