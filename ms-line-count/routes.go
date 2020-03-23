package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CountLinesReq is a body for requesting counting of lines.
type CountLinesReq struct {
	FName string `json:"fname"`
}

//fetchBrowserCounts
func handleCountLines(w http.ResponseWriter, r *http.Request) {

	//object for body
	var CLR CountLinesReq

	//decode data in body
	err := json.NewDecoder(r.Body).Decode(&CLR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fetch data for file name
	lf, err := LogStore.fetchData(CLR.FName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Number of lines:", len(lf.Logs))

	//store num count
	err = LogStore.StoreCountLines(CLR.FName, len(lf.Logs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//convert number of lines to string
	numLines := strconv.Itoa(len(lf.Logs))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, numLines)
}

//fetchBrowserCounts
func handleLineCount(w http.ResponseWriter, r *http.Request) {

	//fetch parameters from url
	params := mux.Vars(r)

	fname := params["fname"]

	lc, err := LogStore.fetchLineCount(fname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lineCountMap := make(map[string]int)

	lineCountMap[fname] = lc[0].Count

	// dataOut, err := json.Marshal(lineCountMap)
	// if err != nil {
	// 	fmt.Println("Error Unmarshalling data", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	res := Response{201, "Success", lineCountMap}

	jOut, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error Unmarshalling data", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jOut))
}
