package api

import (
	"fmt"
	"net/http"
)

func Test_http( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf( w, "test success" )
}
