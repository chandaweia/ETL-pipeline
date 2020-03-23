package main

import (
	"database/sql"
	"fmt"
	"log"
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
	StatusCode int            `json:"StatusCode"`
	Message    string         `json:"message"`
	Data       map[string]int `json:"data"`
}

//LineCountRow represents a row in the database for line counts
type LineCountRow struct {
	Key   string
	Count int
}

//Database controls database functionality
type Database struct {
	db *sql.DB
}

//StoreCountLines stores a logfile in a database
func (d *Database) StoreCountLines(fname string, count int) error {

	sqlStmt := `
	INSERT INTO lineCount (key, count)
	VALUES (?,?)
	`
	statement, err := d.db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = statement.Exec(fname, count)
	if err != nil {
		log.Println("Failed to execute sql", err)
		return err
	}

	return nil

}

//fetchData allows you to fetch log data from db.
func (d *Database) fetchLineCount(fname string) ([]LineCountRow, error) {
	rows, _ := d.db.Query("SELECT * FROM lineCount where key='" + fname + "'")
	lc := []LineCountRow{}
	for rows.Next() {
		lcr := LineCountRow{}
		err := rows.Scan(&lcr.Key,
			&lcr.Count)
		if err != nil {
			log.Println("Failed to fetch data from db: ", err)
			return lc, err
		}
		lc = append(lc, lcr)
	}
	return lc, nil
}

//fetchData allows you to fetch log data from db.
func (d *Database) fetchData(fname string) (LogFile, error) {
	rows, _ := d.db.Query("SELECT * FROM logs where name='" + fname + "'")
	lf := LogFile{}
	for rows.Next() {
		logLine := LogLine{}
		err := rows.Scan(&logLine.Name,
			&logLine.RawLog,
			&logLine.RemoteAddr,
			&logLine.TimeLocal,
			&logLine.RequestType,
			&logLine.RequestPath,
			&logLine.Status,
			&logLine.BodyBytesSent,
			&logLine.HTTPReferer,
			&logLine.HTTPUserAgent,
			&logLine.Created)
		if err != nil {
			log.Println("Failed to fetch data from db: ", err)
			return lf, err
		}
		lf.Logs = append(lf.Logs, logLine)
	}
	return lf, nil
}

func (d *Database) dbinit() {

	//create browser table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS lineCount (
		key TEXT,
		count int
		)
	`

	statement, _ := d.db.Prepare(sqlStmt)
	statement.Exec()
}
