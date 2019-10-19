import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-schedule-id',
    templateUrl: './schedule-id.page.html',
    styleUrls: ['./schedule-id.page.scss'],
})
export class ScheduleIdPage implements OnInit {

    constructor(
        private router: Router,
    ) {
    }

    ngOnInit() {
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store');
    }

}
