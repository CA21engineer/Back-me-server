package infrastructure

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"os"
)

type SqlHandler struct {
	DbMap *gorp.DbMap
}

func NewSqlHandler() db.SqlHandler {
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")

	database, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(db:3306)/"+dbName+"?parseTime=true")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: database, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	dbmap.AddTableWithName(entity.Video{}, "videos").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
