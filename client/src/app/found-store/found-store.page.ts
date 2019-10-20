import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {HttpClient} from "@angular/common/http";
import config from "../../../config";

@Component({
    selector: 'app-found-store',
    templateUrl: './found-store.page.html',
    styleUrls: ['./found-store.page.scss'],
})
export class FoundStorePage implements OnInit {
    private area: string;
    public freeword: string;
    public login_flag: boolean = true;
    private json_data;

    public api_param = [
        '禁煙席あり',
        'クレジットカード対応',
        '飲み放題あり',
        '食べ放題あり',
        '個室あり',
        '深夜も営業',
        '駐車場あり',
        'Wifiあり',
        'プロジェクター・スクリーンあり',
        'web予約可',
    ];

    public send_data_list = [];

    constructor(
        private router: Router,
        private http: HttpClient,
        private route: ActivatedRoute,
    ) {
        for (let param of this.api_param) this.send_data_list.push(false);
    }

    ngOnInit() {
        this.route.paramMap.subscribe((params: ParamMap) => {
            this.area = params.get('area');
            this.freeword = params.get('genre');
        });
        if (this.area === 'null' || this.freeword === 'null') {
            this.area = "";
            this.freeword = "";
            this.login_flag = false;
        }
    }

    goResult() {
        if (!this.area || !this.freeword) {
            alert("エリア、もしくはフリーワードが空白です。");
            return;
        }
        const url: string = config.urlScheme + config.host + config.port + "/app/v0/store_search";
        const formData =
            'no_smorking=' + this.send_data_list[0] +
            '&card=' + this.send_data_list[1] +
            '&bottomless_cup=' + this.send_data_list[2] +
            '&buffet=' + this.send_data_list[3] +
            '&private_room=' + this.send_data_list[4] +
            '&midnight=' + this.send_data_list[5] +
            '&parking=' + this.send_data_list[6] +
            '&wifi=' + this.send_data_list[7] +
            '&projecter_screen=' + this.send_data_list[8] +
            '&web_reserve=' + this.send_data_list[9] +
            '&freeword=' + this.freeword +
            '&area=' + this.area;
        const headers = {"headers": {"Content-Type": "application/x-www-form-urlencoded"}};
        this.http.post(url, formData, headers).subscribe(data => {
            this.json_data = JSON.parse(data['data'].json);
            if (this.json_data === null) alert("見つかりませんでした。");
        }, error => {
            console.log(error);
        });
    }

}
