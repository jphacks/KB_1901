package util

import (
	"fmt"
	"net/http"
	"strconv"
)

// Respond 適切なヘッダーを追加し、レスポンスを行う。
func Respond(res []byte,  w http.ResponseWriter) {

	// ヘッダーセット
	SetCORS(&w)
	// setContentLength(&w, &res)
	// for key, v := range headers {
	// 	w.Header().Set(key, v)
	// }
	SetAllowHeader(&w)
	SetAllowMehod(&w)

	response := string(res)

	// null配列か判定し出力
	if response == "null" {
		fmt.Fprint(w, "[]")
	} else {
		fmt.Fprint(w, response)
	}
}

func SetCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func SetAllowHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func SetAllowMehod(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

func setContentLength(w *http.ResponseWriter, res *[]byte) {
	(*w).Header().Set("Content-Length", strconv.Itoa(len(*res)))
}

func CORSforOptions(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	(*w).WriteHeader(204)
}
