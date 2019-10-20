import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";

@Component({
    selector: 'app-result',
    templateUrl: './result.page.html',
    styleUrls: ['./result.page.scss'],
})
export class ResultPage implements OnInit {

    private json_data;

    constructor(
        private router: Router,
        private route: ActivatedRoute,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.json_data = JSON.parse(params.get('json_data'));
        });
        console.log(this.json_data);
    }

}
