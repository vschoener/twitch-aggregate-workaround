package storage

import (
	"database/sql"
	"errors"
	"fmt"

	// Only used to query database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wonderstream/twitch/logger"
)

// DBName type to enum and restrict db name from settings
type DBName string

const (
	// DBAggregation name
	DBAggregation DBName = "aggregation"
	// DBActivity name
	DBActivity DBName = "activity"
)

// Database storage
type Database struct {
	DB     *sql.DB
	Logger logger.Logger
	Gorm   *gorm.DB
}

// DatabaseManager manage databases
type DatabaseManager struct {
	DBs map[DBName]*Database
}

// GetDM constructor
func GetDM() *DatabaseManager {
	return &DatabaseManager{
		DBs: make(map[DBName]*Database),
	}
}

// ConnectNewDatabase connects to a dB
func (dm *DatabaseManager) ConnectNewDatabase(name DBName, settings DatabaseSettings, l logger.Logger) *DatabaseManager {
	database := NewDatabase()
	database.Logger = l
	database.Logger.SetPrefix("STORAGE")
	dm.DBs[name] = database.Connect(&settings)

	return dm
}

// GetNames of the database
func (dm DatabaseManager) GetNames() []DBName {
	return []DBName{DBAggregation, DBActivity}
}

// Get return a DB instance of the requested DB
func (dm *DatabaseManager) Get(name DBName) *Database {
	if db, ok := dm.DBs[name]; ok {
		return db
	}

	return nil
}

// DatabaseSettings info
type DatabaseSettings struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Check settings integrity
func (ds DatabaseSettings) Check() error {
	var err error
	if len(ds.User) == 0 {
		err = errors.New("User is required")
	} else if len(ds.Password) == 0 {
		err = errors.New("Password is required")
	} else if len(ds.URL) == 0 {
		err = errors.New("URL is required")
	} else if len(ds.Port) == 0 {
		err = errors.New("Port is required")
	} else if len(ds.Name) == 0 {
		err = errors.New("Database Name is required")
	}

	return err
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
	s.Logger.Log(fmt.Sprintf("[Prepare] %#v", q))
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

	s.Logger.Log(fmt.Sprintf("[Query] %#v", query))
	if err != nil {
		s.Logger.LogInterface(err)
		return nil
	}

	return rows
}

// QueryRow has to execute the query and return Row
func (s *Database) QueryRow(query Query, args ...interface{}) *sql.Row {
	s.Logger.Log(fmt.Sprintf("[QueryRow] %#v", query))
	row := s.DB.QueryRow(query.Query, args...)

	return row
}

// ScanRows has to store result insides interfaces args and some more process
func (s *Database) ScanRows(rows *sql.Rows, dest ...interface{}) bool {

	s.Logger.LogInterface("Scan")
	err := rows.Scan(dest...)
	if err != nil {
		s.Logger.LogErrInterface(fmt.Sprintf("[ScanRows] %#v", err))
		return false
	}

	s.Logger.LogInterface(fmt.Sprintf("[Succeed] %+v", rows))

	return true
}

// ScanRow has to store result insides interfaces args and some more process
func (s *Database) ScanRow(row *sql.Row, dest ...interface{}) bool {
	err := row.Scan(dest...)
	if err != nil {
		s.Logger.LogErrInterface(fmt.Sprintf("[ScanRow] %#v", err))
		return false
	}
	s.Logger.LogInterface(fmt.Sprintf("[Succeed] %#v", dest))

	return true
}

// Connect to the server
func (s *Database) Connect(dbSettings *DatabaseSettings) *Database {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbSettings.User,
		dbSettings.Password,
		dbSettings.URL,
		dbSettings.Port,
		dbSettings.Name,
	)
	s.Logger.Log(uri)
	db, err := gorm.Open(
		"mysql",
		uri,
	)
	if err != nil {
		s.Logger.Log(err.Error())
		panic(err.Error())
	}

	s.Gorm = db
	s.DB = db.DB()

	s.Logger.Log("Database storage connected to " + dbSettings.URL)

	return s
}

// IsConnected return current status
func (s *Database) IsConnected() bool {
	return s.DB != nil
}

// Close database connection
func (s *Database) Close() *Database {
	s.DB.Close()
	return s
}
