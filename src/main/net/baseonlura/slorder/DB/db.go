package db

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

type env struct {
	env    string
	port   string
	dbUser string
	dbPass string
	db     string
}

var connection *DB

func (db *DB) Initializer() error {

	// get environment variable
	var envCnf env
	envCnf.env = os.Getenv("ENV")
	envCnf.port = os.Getenv("PORT")
	envCnf.dbUser = os.Getenv("DB_USER")
	envCnf.dbPass = os.Getenv("DB_PASS")
	envCnf.db = os.Getenv("DB")

	if envCnf.env == "" || envCnf.port == "" ||
		envCnf.dbUser == "" || envCnf.dbPass == "" ||
		envCnf.db == "" {
		return errors.New("env")
	}

	// connection DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		envCnf.dbUser, envCnf.dbPass, envCnf.env, envCnf.port, envCnf.db)
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Connection = con

	return err
}

func GetDBConnection() (*DB, error) {
	if connection == nil {
		return nil, errors.New("connection not find")
	}
	return connection, nil
}

func SetDBConnection(con *DB) {
	connection = con
}
