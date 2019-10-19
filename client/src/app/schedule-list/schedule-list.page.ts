import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-schedule-list',
    templateUrl: './schedule-list.page.html',
    styleUrls: ['./schedule-list.page.scss'],
})
export class ScheduleListPage implements OnInit {

    constructor(
        private router: Router,
    ) {
    }

    ngOnInit() {
    }

    goScheduleId(id) {
      this.router.navigateByUrl('/schedule-id');
    }

}
