import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
    selector: 'app-schedule-new',
    templateUrl: './schedule-new.page.html',
    styleUrls: ['./schedule-new.page.scss'],
})
export class ScheduleNewPage implements OnInit {

    constructor(
        private router: Router,
    ) {
    }

    ngOnInit() {
    }

    generateLink() {
        alert("リンクできたよー:http://.......");
        this.router.navigateByUrl('/home');
    }

}
