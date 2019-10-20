package api

import (
	"../util"
	"../logger"
	"../database"
	"../config"
	"../jwt"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	//"os"
)

type Plan_Data struct {
	Name string `json:"name"`
	Memo string `json:"memo"`
	Key string `json:"key"`
	Count int `json:"count"`
}

func Plan_List( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {		
		logger.Write_log( "plan list start " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}

		//アクセストークンによる認証
		err := database.HttpRequestAuth( req, w, keys, conf )
		
		if err != nil {
			logger.Write_log( "access_token check fail", 1 )
			fmt.Fprintf( w, "ログインに失敗しました\n" )
			return
		}

		
		account := req.FormValue( "account" )

		if len( account ) == 0 {
			logger.Write_log( "not set account " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		db, err := database.Connect( conf.DB )
		defer database.Disconnect( db )
		
		if err != nil {
			logger.Write_log( "database not connect", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
			return
		}

		user_id, err := database.Account_ID( db.Sess, account )

		if err != nil {
			logger.Write_log( "fail get ID " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		key_list, err := database.Plan_Key( db.Sess, user_id )
		//fmt.Println( key_list )
		var name_key_list []Plan_Data
		check := Plan_Data{}
		
		for i := 0; i < len( key_list ); i++ {
			file_name := key_list[i] + ".json"
			bytes, err := util.FileDownload( file_name )

			if err != nil {
				logger.Write_log( "fail s3download", 1 )
				logger.Write_log( err.Error(), 1 )
				return
			}
			
			instance := Plan{}

			err = json.Unmarshal( bytes, &instance )
			
			if err != nil {
				logger.Write_log( "fail json change", 4 )
				logger.Write_log( err.Error(), 4 )
				fmt.Fprintf( w, "false" )
				return
			}

			answer_count, err := database.Plan_Answer( db.Sess, key_list[i] )

			if err != nil {
				logger.Write_log( "fail get answer", 4 )
				logger.Write_log( err.Error(), 4 )
				fmt.Fprintf( w, "false" )
				return				
			}
			
			//fmt.Println( instance )
			check.Key = key_list[i]
			check.Name = instance.Plan_Name
			check.Memo = instance.Memo
			check.Count = answer_count
			
			name_key_list = append( name_key_list, check )
		}

		//fmt.Println( name_key_list )

		res_bytes, err := json.Marshal( name_key_list )

		if err != nil {
			logger.Write_log( "fail json change", 1 )
			logger.Write_log( err.Error(), 1 )
			return
		}
		
		var buf bytes.Buffer
	
		_ = json.Indent( &buf, res_bytes, "", "  " )
		
		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "json": buf.String() },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "plan list success" + req.RemoteAddr, 1 )
		util.Respond( res, w )
	}
}
