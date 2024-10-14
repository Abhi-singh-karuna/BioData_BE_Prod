package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	config "myapp/config"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	pvtconfig "github.com/Abhi-singh-karuna/my_Liberary/config"
	"github.com/Abhi-singh-karuna/my_Liberary/gateway"
	sqlhandler "github.com/Abhi-singh-karuna/my_Liberary/sqlhandler"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Write = "write"
)

func ConnectDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	fmt.Println("------------dsn--------", dsn)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("issues with open the database....")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("issues with ping the database....")
		return nil, err
	}

	if err := runInitScript(db, "././migration/init.sql", cfg); err != nil {
		fmt.Println("issues with ./app/migration/init.sql the database....")

		return nil, err
	}
	fmt.Println("here connected successfully.....")

	return db, nil
}

func runInitScript(db *sql.DB, scriptPath string, cfg config.DatabaseConfig) error {
	sqlBytes, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to read SQL script file: %v", err)
	}

	// Initialize the sqlHandler map
	sqlHandler := make(map[string]gateway.SqlHandler)

	// Assuming you have a logger and config package
	logger := baselogger.NewBaseLogger()

	// Replace with your actual logger
	config := pvtconfig.SQL{
		Host:     cfg.Host,
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	}
	// Initialize the SQL handler
	sqlHandler[Write] = sqlhandler.NewSqlHandler(logger, config)

	if err := sqlHandler[Write].MultiExec(string(sqlBytes)); err != nil {
		return err
	}
	return nil
}
