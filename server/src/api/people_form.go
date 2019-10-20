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
	"os"
)

type Form_Day struct {
	Day string `json:"day"`
	Check int `json:check`
}

type Form struct {
	Area string `json:"area"`
	Genre string `json:"genre"`
	Free string `json:"free"`
	Day []Form_Day `json:"select_day"`
}


func Plan_Form( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "form add start " + req.RemoteAddr, 1 )

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

		form_data := req.FormValue( "form_data" )

		if len( form_data ) == 0 {
			logger.Write_log( "not set form_data " + req.RemoteAddr, 1 )
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

		count, err := database.Plan_Answer( db.Sess, plan_key )

		if err != nil {
			logger.Write_log( "count not get", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
			return
		}

		current_form := Form{}
		err = json.Unmarshal( []byte( form_data ), &current_form )

		if err != nil {
			logger.Write_log( "fail json change", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
			return
		}

		var form_storage []Form
		file_name := plan_key + "_form.json" 
		
		if count == 0 {
			form_storage = append( form_storage, current_form )
			err = json_upload( form_storage, file_name )

			if err != nil {
				logger.Write_log( "fail s3upload " + req.RemoteAddr, 1 )
				logger.Write_log( err.Error(), 1 )
				fmt.Fprintf( w, "false" )
				return
			}

		} else {
			json_byte, err := util.FileDownload( file_name )

			if err != nil {
				logger.Write_log( "fail s3download " + req.RemoteAddr, 1 )
				logger.Write_log( err.Error(), 1 )
				fmt.Fprintf( w, "false" )
				return			
			}

			err = json.Unmarshal( json_byte, &form_storage )
			
			if err != nil {
				logger.Write_log( "fail json_storage change", 4 )
				//logger.Write_log( string( json_byte ), 4 )
				logger.Write_log( err.Error(), 4 )
				fmt.Fprintf( w, "false" )
				return
			}

			form_storage = append( form_storage, current_form )
			err = json_upload( form_storage, file_name )

			if err != nil {
				logger.Write_log( "fail s3upload " + req.RemoteAddr, 1 )
				logger.Write_log( err.Error(), 1 )
				fmt.Fprintf( w, "false" )
				return
			}
		}

		err = database.Plan_Answer_Update( db.Conn, plan_key, count + 1 )

		if err != nil {
			logger.Write_log( "count not change", 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
			return
		}
		
		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{},
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "form add success" + req.RemoteAddr, 1 )
		util.Respond( res, w )
	}
}

func json_upload( data []Form,  file_name string ) error {
	json_form_byte, err := json.Marshal( data )
	
	if err != nil {
		return err
	}
	
	var buf bytes.Buffer
	
	_ = json.Indent( &buf, json_form_byte, "", "  " )
	
	file, _ := os.OpenFile("send.json", os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Fprintln( file, buf.String() )
	defer file.Close()
	
	err = util.FileUpload( file_name )
	
	if err != nil {
		return	err		
	}

	return nil
}
