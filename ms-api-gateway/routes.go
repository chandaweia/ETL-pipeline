package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//processLogFile takes in an uploaded logfile, stores the data, processes stats.
func processLogFile(rawLogFile []byte, fname string) bool {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(rawLogFile)))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	logFile := parseFile(lines, fname)

	//Store parsed logs
	LogStore.StoreLogLine(logFile)

	return true
}

//handleServeUploadPage serves the static html file
func handleServeUploadPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/upload.html")
}

//handleUploadLog handles uploading of log file and triggers etl pipeline
func handleUploadLog(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Received Uploaded File: %+v\n", handler.Filename)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// return that we have successfully uploaded our file!
	processLogFile(fileBytes, handler.Filename)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Pipeline Status</h1>")
	fmt.Fprintf(w, "Log File Uploaded Successfully<br>")
	fmt.Fprintf(w, `<strong>Count Browsers</strong>: <font size="3" color="green">Completed</font><br>`)
	fmt.Fprintf(w, `<strong>Count Visitors</strong>: <font size="3" color="green">Completed</font><br>`)

}
