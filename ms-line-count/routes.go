package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//fetchBrowserCounts
func handleLineCount(w http.ResponseWriter, r *http.Request) {
	browserData := LogStore.fetchBrowserData()

	browserStats := make(map[string]int)

	for _, v := range browserData {
		if _, ok := browserStats[v.Browser]; !ok {
			browserStats[v.Browser] = 0
		}
		browserStats[v.Browser] += v.Count

	}

	jOut, err := json.Marshal(browserStats)
	if err != nil {
		fmt.Println("Error Unmarshalling data", err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jOut))
}
