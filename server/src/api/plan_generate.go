package api

import (
	"../util"
	"../logger"
	"../database"
	"../config"
	"../jwt"
	"math/rand"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"os"
)

type Plan struct {
	Plan_Name string `json:"plan_name"`
	Memo string `json:"memo"`
	Day [] string `json:"day"`
	Genre [] string `json:"genre"`
	Area [] string `json:"area"`
}

func Plan_Generate( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "plan generate start " + req.RemoteAddr, 1 )

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

		plan_data := req.FormValue( "plan_data" )
		account := req.FormValue( "account" )

		if len( plan_data ) == 0 {
			logger.Write_log( "not set plan " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		if len( account ) == 0 {
			logger.Write_log( "not set account " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}
		/*
		new_plan := new( Plan )
		err = json.Unmarshal( []byte( plan_data ), new_plan )

		if err != nil {
			logger.Write_log( "fail change json " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return			
		}
*/

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

		plan_key := key_generate()

		file_name := plan_key + ".json"

		file, _ := os.OpenFile("send.json", os.O_WRONLY|os.O_CREATE, 0666 )
		fmt.Fprintln( file, plan_data )
		defer file.Close()
		err = util.FileUpload( file_name )

		if err != nil {
			logger.Write_log( "fail s3upload " + req.RemoteAddr, 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return			
		}
		
		err = database.Plan_Insert( user_id, plan_key, db.Conn )

		if err != nil {
			logger.Write_log( "fail database insert " + req.RemoteAddr, 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "key": plan_key },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "plan generate success " + req.RemoteAddr, 1 )
		util.Respond( res, w )
	}
}

func key_generate() string {
	rand.Seed( time.Now().UnixNano() )
	i := 0
	cha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	paw := ""

	for i < 20 {
		a := rand.Intn( 52 )
		paw += string( cha[a] )
		i += 1
	}

	return paw
}

