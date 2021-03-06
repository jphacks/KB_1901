# らくだ幹事長
![logo](imgs/logos/1_rakuda_logo_orange.png)
[![video](Screenshot.png)](https://youtu.be/cooiSmNAkRI)

## 製品概要


### 飲み会 x Tech
**今度の打ち上げ全部任せましたよ！幹事さん！**

### 背景（製品開発のきっかけ、課題等）

「じゃあ飲み会の幹事は君ね」「今度のご飯だれが企画する？え、私？」
そういっていきなり幹事を任されても面倒なことばっかり。
誰か、スケジュール調整からお店探しまで幹事のお仕事をサポートしてほしい……
それがこのサービス開発の原点でした。

幹事に任された時に役に立つサービスはたくさんあります。
例えばスケジュール調整するサービス。お店を探すサイト。
たくさんありますが、多すぎて何を使えばいいのか、何がなんだかわかりません。
また、全てが一つのサービスに集約されているものはありませんでした。

また面倒なことが多すぎて幹事をしてくれる人も積極的には出てこなくて結局ご飯の予定は立ち消え……なんてことも。
そんな問題を解決するのがこのサービスです！

### 製品説明（具体的な製品の説明）
幹事のお仕事が「楽だ！！」と思えるアプリケーション。それが **らくだ幹事長** です。
スケジュール調整からアンケート集約、お店探しまで幹事さんのやるべきことをお手伝いします！

### 特長
#### 1. すべての幹事のお仕事をサポート！
必要最低限の情報を入力するだけでアンケートフォームが完成！
簡単な操作ですぐに企画することができます。

ご飯の予定を企画したら、あなたのすることはメンバーにリンクを渡すだけです。

#### 2. 最適なお店を提案いたします
参加メンバーの回答から最適な日程やエリア、お店をいくつか提案します。
メンバーの好みから提案するので後からお店選びに文句も言われません。

### 解決出来ること
スケジュール調整やお店探しなど幹事さんの煩わしい作業からあなたを解放します。
お店を探してインターネットの世界を右往左往することもありません。

### 今後の展望
*今回は実現できなかったが、今後改善すること、どのように展開していくことが可能かについて記載をしてください。*

- アカウント作成フォーム
- マルチプラットフォーム対応
- デザインの充実

## 開発内容・開発技術
### 活用した技術
#### API・データ
今回スポンサーから提供されたAPI、製品などの外部技術があれば記述をして下さい。

* GURUNAVI レストラン検索API 
> https://api.gnavi.co.jp/api/

#### フレームワーク・ライブラリ・モジュール
* Ionic Framework (Web UI)
* aws:ec2 (仮想サーバー)
* aws:RDS (データベースサーバー)
* aws:S3 (ファイルサーバー)
* aws-sdk-go
* PostgreSQL (データベース)

#### デバイス
* *特になし*

### 独自開発技術（Hack Dayで開発したもの）

#### 2日間に開発した独自の機能・技術
- 予定作成・管理API
- フォーム回答API
- 店検索API
- APIserver
- アカウント情報DB
- 予定情報DB
- クライアント側UI

## 実行

1. `cd client`  
2. `npm install`  
3. `npm install -g ionic@4.12.0`  
4. `ionic serve`  

