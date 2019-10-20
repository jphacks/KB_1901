package main

import (
	"./api"
	"./jwt"
	"./config"
	"./logger"
	"fmt"
	"net/http"
)

func main() {
	config_data := config.Config_data()

	keys, err := jwt.ReadKeys()

	if err != nil {
		logger.Write_log( "fail key generate", 4 )
		return
	}
	
	http.HandleFunc( "/test/v0", api.Test_http )
	http.HandleFunc( "/app/v0/login", api.Login( config_data, keys ) )
	http.HandleFunc( "/app/v0/plan_generate", api.Plan_Generate( config_data, keys ) )
	http.HandleFunc( "/app/v0/plan_form", api.Plan_Form( config_data, keys ) )
	http.HandleFunc( "/app/v0/plan_check", api.Plan_Check( config_data, keys ) )
	http.HandleFunc( "/app/v0/plan_list", api.Plan_List( config_data, keys ) )
	http.HandleFunc( "/app/v0/plan_result", api.Plan_Result( config_data, keys ) )

	http.HandleFunc( "/app/v0/store_search", api.Store_Search( config_data, keys ) )
	
	err = http.ListenAndServe( ":80", nil )

	if err != nil {
		fmt.Println( err.Error() )
		logger.Write_log( err.Error(), 1 )
	}
}
