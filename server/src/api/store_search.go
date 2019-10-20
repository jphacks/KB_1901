package api

import (
	"../util"
	"../logger"
	//"../database"
	"../config"
	"../jwt"
	"encoding/json"
	"net/http"
	"strings"
	"bytes"
	"fmt"
	//"os"
)

type Store_Data struct {
	Store_Name string `json:"store_name"`
	Category string `json:"category"`
	Tel_Number string `json:"tel_number"`
	URL string `json:"url"`
	Store_Image string `json:"store_image"`
	Rest_Day string `json:"rest_day"`
	Area string `json:"area"`
	Average_Money int `json:"average_money"`
}

func Store_Search( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "store search start " + req.RemoteAddr, 1 )

		if req.Method == "OPTIONS" {
			logger.Write_log( "OPTIONS return", 1 )
			util.CORSforOptions( &w )
			return
		}

		api_par := map[string]string{}

		if len( req.FormValue( "freeword" ) ) != 0 {
			api_par["freeword"] = req.FormValue( "freeword" )
		} else {
			api_par["freeword"] = "0"
		}
		
		if len( req.FormValue( "no_smorking" ) ) != 0 { 
			api_par["no_smorking"] = value_go( req.FormValue( "no_smorking" ) )
		} else {
			api_par["no_smorking"] = "0"
		}

		if len( req.FormValue( "card" ) ) != 0  {
			api_par["card"] = value_go( req.FormValue( "card" ) )
		} else {
			api_par["card"] = "0"
		}

		if len( req.FormValue( "bottomless_cup" ) ) != 0 {
			api_par["bottomless_cup"] = value_go( req.FormValue( "bottomless_cup" ) )
		} else {
			api_par["bottomless_cup"] = "0"
		}

		if len( req.FormValue( "buffet" ) ) != 0 {
			api_par["buffet"] = value_go( req.FormValue( "buffet" ) )
		} else {
			api_par["buffet"] = "0"
		}

		if len( req.FormValue( "private_room" ) ) != 0 {
			api_par["private_room"] = value_go( req.FormValue( "private_room" ) )
		} else {
			api_par["private_room"] = "0"
		}

		if len( req.FormValue( "midnight") ) != 0 {
			api_par["midnight"] = value_go( req.FormValue( "midnight" ) )
		} else {
			api_par["midnight"] = "0"
		}

		if len( req.FormValue( "wifi") ) != 0 {
			api_par["wifi"] = value_go( req.FormValue( "wifi" ) )
		} else {
			api_par["wifi"] = "0"
		}

		if len( req.FormValue( "projecter_screen" ) ) != 0 {
			api_par["projecter_screen"] = value_go( req.FormValue( "projecter_screen" ) )
		} else {
			api_par["projecter_screen"] = "0"
		}

		if len( req.FormValue( "web_reserve" ) ) != 0 {
			api_par["web_reserve"] = value_go( req.FormValue( "web_reserve" ) )
		} else {
			api_par["web_reserve"] = "0"
		}

		url_option := ""

		for k, v := range api_par {
			if v != "0" {
				url_option += "&" + k + "=" + v
			}
		}

		search := false

		if len( req.FormValue( "area" ) ) != 0 {
			search_area := req.FormValue( "area" )
			search_area = strings.Replace( search_area, "駅", "", -1 )
			
			area_result, err := AreaReturnStruct()
			if err == nil {
				for i := 0; i < len( area_result.GareaSmall ); i++ {
					s_slice := strings.Split( area_result.GareaSmall[i].AreanameS, "・" )

					if !search { 
						for r := 0; r < len( s_slice ); r++ {
							if s_slice[r] == search_area {
								url_option += "&areacode_s=" + area_result.GareaSmall[i].AreacodeS
								search = true
								logger.Write_log( "area check ok", 1 )
								break
							}
						}
					}

					if !search {
						m_slice := area_result.GareaSmall[i].GareaMiddle.AreanameM

						if m_slice == search_area {
							url_option += "&areacode_m=" + area_result.GareaSmall[i].GareaMiddle.AreacodeM
							search = true
							logger.Write_log( "area check ok", 1 )
							break
						}
					}


					l_search_area := ""
					if search_area == "東京" {
						l_search_area = search_area + "都"
					} else if search_area == "大阪" {
						l_search_area = search_area  + "府"
					} else if search_area == "京都" {
						l_search_area = search_area + "府"
					} else if search_area == "北海道" {
						l_search_area = search_area
					} else {
						l_search_area = search_area + "県"
					}

					if !search {
						if area_result.GareaSmall[i].Pref.PrefName == l_search_area {
							url_option += "&pref=" + area_result.GareaSmall[i].Pref.PrefCode
							search = true
							logger.Write_log( "area check ok", 1 )
							break						
						}
						
						if search {
							break
						}
					}
					
				}
			} else {
				logger.Write_log( err.Error(), 1 )
			}
		}

		result, err := StoreReturnStruct( url_option )

		if err != nil {
			logger.Write_log( "fail tap api", 1 )
			logger.Write_log( err.Error(), 1 )
			
			responseResult := ResponseResult{
				Status:    "No",
				Data:      map[string]interface{}{},
				ErrorText: "",
			}

			res, _ := json.Marshal( responseResult )
			
			util.Respond( res, w )
			return
		}

		var res_store_data []Store_Data

		for i := 0; i < len( result.Rest ); i++ {
			instance := Store_Data{}
			instance.Store_Name = result.Rest[i].Name
			instance.Category = result.Rest[i].Category
			instance.URL = result.Rest[i].URL
			instance.Store_Image = result.Rest[i].ImageURL.ShopImage1
			instance.Rest_Day = result.Rest[i].Holiday
			instance.Area = result.Rest[i].Code.AreanameS
			instance.Average_Money = result.Rest[i].Budget
			instance.Tel_Number = result.Rest[i].Tel
			res_store_data = append( res_store_data, instance )
		}

		json_byte, err := json.Marshal( res_store_data )

		if err != nil {
			logger.Write_log( "fail change json", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false")
			return
		}

		var buf bytes.Buffer
	
		_ = json.Indent( &buf, json_byte, "", "  " )

		responseResult := ResponseResult{
			Status:    "OK",
			Data:      map[string]interface{}{ "json": buf.String() },
			ErrorText: "",
		}

		res, _ := json.Marshal( responseResult )

		logger.Write_log( "store search success " + req.RemoteAddr, 1 )

		util.Respond( res, w )
 	}
}

func value_go( bool_value string ) string {
	if bool_value == "false" {
		return "0"
	} else {
		return "1"
	}
}
