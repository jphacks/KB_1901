import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {HttpClient} from "@angular/common/http";

@Component({
    selector: 'app-schedule-id',
    templateUrl: './schedule-id.page.html',
    styleUrls: ['./schedule-id.page.scss'],
})
export class ScheduleIdPage implements OnInit {
    private auth_token: string;
    private key: string;

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.key = params.get('key');
        });
    }

    goFoundStore() {
        this.router.navigateByUrl('/found-store');
    }

}
