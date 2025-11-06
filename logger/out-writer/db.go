package out_writer

import (
	"database/sql"
	"fmt"
	"logger/logger"
)

type DBSink struct {
	db *sql.DB
}

func NewDBSink(connString string) (*DBSink, error) {
	// Example: Replace with actual DB connection logic
	// return &DBSink{db: db}, nil
	return &DBSink{}, nil
}

func (d *DBSink) Write(level logger.LogLevel, msg string) error {
	// example
	fmt.Printf("DBSink: [%s] %s\n", level, msg)
	return nil
}

func (d *DBSink) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}
