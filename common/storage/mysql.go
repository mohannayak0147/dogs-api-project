package storage

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// connect open a mysql connection and pings it
func connect(c mysql.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("connection failed with to %s: %s", c.FormatDSN(), err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database ping failed: %w , address: %s , user: %s , network: %s, dbname: %s", err, c.Addr, c.User, c.Net, c.DBName)
	}
	log.Debug("connected to DB", c.Addr)
	return db, nil
}

// GetDB created DB object
func GetDB(user, password, host, dbname string, port int) (*sqlx.DB, error) {
	return connect(mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", host, port),
		DBName:               dbname,
		AllowNativePasswords: true,
	})
}
