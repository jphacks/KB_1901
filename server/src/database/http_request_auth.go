package database

import (
	"../config"
	"../jwt"
	"net/http"
)


// HttpRequestAuth HTTPリクエストに必ず必要な認証
func HttpRequestAuth(req *http.Request, w http.ResponseWriter, keys *jwt.JWTKeys, conf config.Connect_data ) error {
	_, err := jwt.CheckToken(*req, keys)
	
	if err != nil {
		//http.Error(w, "ログインに失敗しました。", http.StatusUnauthorized)
		return err
	}

	db, err := Connect(conf.DB)
	if err != nil {
		//http.Error(w, "アカウント情報を取得できませんでした。", http.StatusServiceUnavailable)
		return err
	}

	defer Disconnect(db)

	return nil
}

