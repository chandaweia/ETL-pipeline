package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//LogFile represents a logfile with multiple lines
type LogFile struct {
	Logs []LogLine
}

//LogLine represents fields in a given log line
type LogLine struct {
	Name          string
	RawLog        string
	RemoteAddr    string
	TimeLocal     string
	RequestType   string
	RequestPath   string
	Status        int
	BodyBytesSent int
	HTTPReferer   string
	HTTPUserAgent string
	Created       time.Time
}

//Response HTTP RESPONSE for messages
type Response struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       map[string]int `json:"data"`
}

//ErrorResponse, response to send when erroring out
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	json       string
}

//NewError generates a new http error
func NewError(code int, e string) ErrorResponse {
	err := ErrorResponse{code, e, ""}
	temp, _ := err.JSON()
	err.json = temp
	return err
}

//JSON returns json version of type
func (r *Response) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}

//JSON returns json version of type
func (r *ErrorResponse) JSON() (string, error) {
	jOut, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error Unmarshalling data", err)
		return "", err
	}

	return string(jOut), nil
}
