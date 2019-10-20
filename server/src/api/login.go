package api

import (
	"../jwt"
	"../util"
	"../config"
	"../database"
	"../logger"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
)

// ResponseResult レスポンス結果に関する構造体
type ResponseResult struct {
	Status    string                 `json:"status"`
	Data      map[string]interface{} `json:"data"`
	ErrorText string                 `json:"errorText"`
}

func Login( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		
		logger.Write_log( "login start" + " " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}

		account := req.FormValue("account")
		password := req.FormValue("password")

		if len( account ) == 0 ||
			30 < len( account ) {
			logger.Write_log( "not set account", 1 )
			fmt.Fprintf( w, "ログインに失敗しました" )
			return
		}

		if len( password ) == 0 ||
			10 < len( password ) {
			logger.Write_log( "not set password", 1 )
			fmt.Fprintf( w, "ログインに失敗しました" )
			return
		}

		db, err := database.Connect( conf.DB )
		defer database.Disconnect( db )
		
		if err != nil {
			logger.Write_log( "database not connect", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "ログインに失敗しました" )
			return
		}

		err = database.IsLogin( db.Sess, account, password )

		if err != nil {
			logger.Write_log( "password or account wrong", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "ログインに失敗しました" )
			return
		}

		tokenString, err := jwt.NewToken( keys, account, time.Now().UTC() )

		
		if err != nil {
			logger.Write_log( "generate access_token fail", 4 )
			fmt.Fprintf( w, "ログインに失敗しました\n" )
			return
		}

		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "token": tokenString },
			ErrorText: "",
		}

		res, err := json.Marshal( responseResult )

		if err != nil {
			logger.Write_log( "generate  json_file fail", 4 )
			fmt.Fprintf( w, "ログインに失敗しました\n" )
			return
		}

		logger.Write_log( "login success " + req.RemoteAddr, 1 )
		util.Respond( res, w )
	}
}
