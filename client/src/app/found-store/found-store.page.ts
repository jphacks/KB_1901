import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {HttpClient} from "@angular/common/http";

@Component({
    selector: 'app-found-store',
    templateUrl: './found-store.page.html',
    styleUrls: ['./found-store.page.scss'],
})
export class FoundStorePage implements OnInit {
    private auth_token: string;
    private components: string;

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.auth_token = params.get('auth_token');
            this.components = params.get('components');
        });
    }

    goResult() {
        this.router.navigateByUrl('/result');
    }

}
