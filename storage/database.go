package storage

import (
	"database/sql"
	"fmt"

	// Only used to query database
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderstream/twitch/logger"
)

// Database storage
type Database struct {
	DB     *sql.DB
	Logger logger.Logger
}

// DatabaseSettings info
type DatabaseSettings struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Query contains query and Parameters to process sql query
type Query struct {
	Query      string
	Parameters map[string]interface{}
}

// NewDatabase load the object with credential
func NewDatabase() *Database {
	return &Database{}
}

// Prepare does additional process and calls sql.Prepare()
func (s *Database) Prepare(q Query) *sql.Stmt {
	s.Logger.Log(fmt.Sprintf("[Storage][Prepare] %#v", q))
	stmt, err := s.DB.Prepare(q.Query)
	if err != nil {
		s.Logger.LogInterface(err)
		return nil
	}

	return stmt
}

// Run has to prepare and exec the query
func (s *Database) Run(q Query, args ...interface{}) bool {
	stmt := s.Prepare(q)
	if stmt == nil {
		return false
	}
	defer stmt.Close()

	_, err := s.DB.Exec(q.Query, args...)
	if err != nil {
		s.Logger.LogInterface(err)
		return false
	}

	return true
}

// Query has to execute the query and return Rows
func (s *Database) Query(query Query, args ...interface{}) *sql.Rows {
	rows, err := s.DB.Query(query.Query, args...)

	s.Logger.Log(fmt.Sprintf("[Storage][Query] %#v", query))
	if err != nil {
		s.Logger.LogInterface(err)
		return nil
	}

	return rows
}

// QueryRow has to execute the query and return Row
func (s *Database) QueryRow(query Query, args ...interface{}) *sql.Row {
	s.Logger.Log(fmt.Sprintf("[Storage][QueryRow] %#v", query))
	row := s.DB.QueryRow(query.Query, args...)

	return row
}

// ScanRows has to store result insides interfaces args and some more process
func (s *Database) ScanRows(rows *sql.Rows, dest ...interface{}) bool {

	s.Logger.LogInterface("Scan")
	err := rows.Scan(dest...)
	if err != nil {
		s.Logger.LogErrInterface(fmt.Sprintf("[Storage][ScanRows] %#v", err))
		return false
	}

	s.Logger.LogInterface(fmt.Sprintf("[Storage][Succeed] %+v", rows))

	return true
}

// ScanRow has to store result insides interfaces args and some more process
func (s *Database) ScanRow(row *sql.Row, dest ...interface{}) bool {
	err := row.Scan(dest...)
	if err != nil {
		s.Logger.LogErrInterface(fmt.Sprintf("[Storage][ScanRow] %#v", err))
		return false
	}
	s.Logger.LogInterface(fmt.Sprintf("[Storage][Succeed] %#v", dest))

	return true
}

// Connect to the server
func (s *Database) Connect(dbSettings *DatabaseSettings) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbSettings.User,
			dbSettings.Password,
			dbSettings.URL,
			dbSettings.Port,
			dbSettings.Name,
		),
	)
	if err != nil {
		s.Logger.Log(err.Error())
		panic(err.Error())
	}
	s.DB = db

	s.Logger.Log("Database storage connected to " + dbSettings.URL)
}

// IsConnected return current status
func (s *Database) IsConnected() bool {
	return s.DB != nil
}
