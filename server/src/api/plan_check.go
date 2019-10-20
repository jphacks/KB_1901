package api

import (
	"../util"
	"../logger"
	//"../database"
	"../config"
	"../jwt"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	//"os"
)

func Plan_Check( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "plan check start " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}

		plan_key := req.FormValue( "plan_key" )

		if len( plan_key ) == 0 {
			logger.Write_log( "not set plan_key " + req.RemoteAddr, 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		file_name := plan_key + ".json"
		json_bytes, err := util.FileDownload( file_name )

		if err != nil {
			logger.Write_log( "fail s3download", 1 )
			logger.Write_log( err.Error(), 1 )
			return
		}

		var buf bytes.Buffer
		
		_ = json.Indent( &buf, json_bytes, "", "  " )

		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "json": buf.String() },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )
		
		logger.Write_log( "plan check success" + req.RemoteAddr, 1 )
		util.Respond( res, w )

	}
}
