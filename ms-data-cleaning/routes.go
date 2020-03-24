package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//CountLinesReq is a body for requesting counting of lines.
type CountLinesReq struct {
	FName string `json:"fname"`
}

//fetchBrowserCounts
func handleCountLines(w http.ResponseWriter, r *http.Request) {

	//LOGIC for this endpoint

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, numLines)
}

//fetchBrowserCounts
func handleLineCount(w http.ResponseWriter, r *http.Request) {

	//fetch parameters from url
	params := mux.Vars(r)

	fname := params["fname"]

	//create response and get json
	res := Response{201, "Success", lineCountMap}
	jOut, _ := res.JSON() 


	


	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jOut))
}
