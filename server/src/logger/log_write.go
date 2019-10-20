package logger

import (
	"os"
	"fmt"
	"time"
	//"strings"
	"strconv"
)

func Write_log( log string, level int ) {
	file := open_logfile()

	defer file.Close()

	now := time.Now()
	jst := time.FixedZone( "Asia/Tokyo", 9*60*60 )
	nowJST := now.In( jst )
	
	current_time := strconv.Itoa( nowJST.Year() ) + "-" + strconv.Itoa( int( nowJST.Month() ) ) + "-" +
		strconv.Itoa( nowJST.Day() ) + " " + strconv.Itoa( nowJST.Hour() ) + ":" +strconv.Itoa( nowJST.Minute() ) +
		":" + strconv.Itoa( nowJST.Second() )

	var level_name string
	var message string

	if level == 1 {
		level_name = "INFO"
	} else if level == 2 {
		level_name = "WARN"
	} else if level == 3 {
		level_name = "ERROR"
	} else if level == 4 {
		level_name = "FATAL"
	} else {
		return
	}

	message = current_time + " " + level_name + " " + log

	fmt.Fprintln( file, message )
}

func open_logfile() *os.File {
	file, err := os.OpenFile( "../log/api_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666 )
	
    if err != nil {
        //エラー処理
		fmt.Println( err )
        return nil
    }

	return file
}
