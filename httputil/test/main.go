package main

import (
	"fmt"
	"net/http"

	"github.com/ntons/tongo/httputil"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(httputil.GetRemoteIpFromRequest(r))
	})
	http.ListenAndServe(":8800", mux)
}
