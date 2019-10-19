import {Component} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-home',
    templateUrl: 'home.page.html',
    styleUrls: ['home.page.scss'],
})
export class HomePage {

    constructor(
        private router: Router,
    ) {
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
