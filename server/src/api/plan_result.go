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
	"sort"
	"fmt"
	//"os"
)

type Result struct {
	Answer_Count int `json:"answer_count"`
	Plan_Name string `json:"plan_name"`
	Day []string `json:"day"`
	Genre []string `json:"genre"`
	Area []string `json:"area"`
	Free []string `json:"free"`
}

type Result_Sort struct {
	Data string
	Age int
}

type ByAge []Result_Sort

func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age >= a[j].Age }

func Plan_Result( conf config.Connect_data, keys *jwt.JWTKeys) http.HandlerFunc {
	return func( w http.ResponseWriter, req *http.Request ) {
		logger.Write_log( "plan result start " + req.RemoteAddr, 1 )
		
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

		plan_key := req.FormValue( "plan_key" )

		if len( plan_key ) == 0 {
			logger.Write_log( "not set plan_key " + req.RemoteAddr, 1 )
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

		answer_count, err := database.Plan_Answer( db.Sess, plan_key )

		if err != nil {
			logger.Write_log( "fail get count", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return			
		}

		if answer_count == 0 {
			responseResult := ResponseResult{
				Status:    "OK",
				Data:      map[string]interface{}{ "count": 0 },
				ErrorText: "",
			}
			
			res, _ := json.Marshal( responseResult )
			
			logger.Write_log( "plan result form 0 " + req.RemoteAddr, 1 )
			util.Respond( res, w )
			return
		}

		plan_file_name := plan_key + ".json"
		json_byte, err := util.FileDownload( plan_file_name )

		if err != nil {
			logger.Write_log( "fail s3download " + req.RemoteAddr, 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return
		}
		
		plan_data := Plan{}
		err = json.Unmarshal( json_byte, &plan_data )

		if err != nil {
			logger.Write_log( "fail json change", 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		plan_name := plan_data.Plan_Name

		answer_file_name := plan_key + "_form.json"
		json_byte, err = util.FileDownload( answer_file_name )

		if err != nil {
			logger.Write_log( "fail s3download " + req.RemoteAddr, 1 )
			logger.Write_log( err.Error(), 1 )
			fmt.Fprintf( w, "false" )
			return
		}

		var answer_stoarge []Form
		err = json.Unmarshal( json_byte, &answer_stoarge )

		day_score := map[string]int{}
		genre_score := map[string]int{}
		area_score := map[string]int{}
		
		for i := 0; i < len( plan_data.Area ); i++ {
			area_score[ plan_data.Area[i] ] = 0
		}

		for i := 0; i < len( plan_data.Genre ); i++ {
			genre_score[ plan_data.Area[i] ] = 0
		}

		for i := 0; i < len( plan_data.Day ); i++ {
			day_score[ plan_data.Area[i] ] = 0
		}

		res_result := Result{}//最終

		for i := 0; i < len( answer_stoarge ); i++ {
			area_score[ answer_stoarge[i].Area ] += 1
			genre_score[ answer_stoarge[i].Genre ] += 1

			for r := 0; r < len( answer_stoarge[i].Day ); r++ {
				day_score[ answer_stoarge[i].Day[r].Day ] += answer_stoarge[i].Day[r].Check
			}

			res_result.Free = append( res_result.Free, answer_stoarge[i].Free )
		}

		day := ByAge{}
		for k, v := range day_score {
			day = append( day, Result_Sort{ k, v } )
		}

		sort.Sort( day )

		area := ByAge{}
		for k, v := range area_score {
			area = append( area, Result_Sort{ k, v } )
		}

		sort.Sort( area )

		genre := ByAge{}
		for k, v := range genre_score {
			genre = append( genre, Result_Sort{ k, v } )
		}

		sort.Sort( genre )

		res_result.Area = append( res_result.Area, area[0].Data )
		res_result.Area = append( res_result.Area, area[1].Data )

		res_result.Genre = append( res_result.Genre, genre[0].Data )
		res_result.Genre = append( res_result.Genre, genre[1].Data )

		res_result.Day = append( res_result.Day, day[0].Data )
		res_result.Day = append( res_result.Day, day[0].Data )

		res_result.Answer_Count = answer_count
		res_result.Plan_Name = plan_name

		json_byte, err = json.Marshal( res_result )

		if err != nil {
			logger.Write_log( "fail json_storage change", 4 )
			//logger.Write_log( string( json_byte ), 4 )
			logger.Write_log( err.Error(), 4 )
			fmt.Fprintf( w, "false" )
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

		logger.Write_log( "plan result success " + req.RemoteAddr, 1 )
		util.Respond( res, w )

	}
}
