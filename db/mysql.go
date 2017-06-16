package db

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"

	"frest_echo/config"
)

func Init() *dbr.Session {

	session := getSession()

	return session
}

func getSession() *dbr.Session {

	db, err := dbr.Open("mysql",
		config.USER+":"+config.PASSWORD+"@tcp("+config.HOST+":"+config.PORT+")/"+config.DB+"?parseTime=true",
		nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
