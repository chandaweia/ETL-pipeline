package main

import (
	"database/sql"
	"time"
)

//LogFile represents a logfile with multiple lines
type LogFile struct {
	Logs []LogLine
}

//LogLine represents fields in a given log line
type LogLine struct {
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

//LineCountRow represents a row in the database for line counts
type LineCountRow struct {
	Key     string
	Date    string
	Browser string
	Count   int
}

//Database controls database functionality
type Database struct {
	db *sql.DB
}

//fetchData allows you to fetch log data from db.
func (d *Database) fetchData() LogFile {
	rows, _ := d.db.Query("SELECT * FROM logs")
	lf := LogFile{}
	for rows.Next() {
		logLine := LogLine{}
		rows.Scan(&logLine.RawLog,
			&logLine.RemoteAddr,
			&logLine.TimeLocal,
			&logLine.RequestType,
			&logLine.RequestPath,
			&logLine.Status,
			&logLine.BodyBytesSent,
			&logLine.HTTPReferer,
			&logLine.HTTPUserAgent,
			&logLine.Created)
		lf.Logs = append(lf.Logs, logLine)
	}
	return lf
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
