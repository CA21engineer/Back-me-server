package infrastructure

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"os"
)

type SqlHandler struct {
	DbMap *gorp.DbMap
}

func NewSqlHandler() db.SqlHandler {
	mode := os.Getenv("GIN_MODE")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	instanceConnectionName := os.Getenv("CLOUDSQL_CONNECTION_NAME")

	var dbURI string

	// GAEデプロイ時には接続先をCloudSQLにする
	if mode == "release" {
		log.Println("Connecting to CloudSQL instance...")
		dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", dbUser, dbPass, instanceConnectionName, dbName)
	} else {
		log.Println("Connecting to local db instance...")
		dbURI = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", dbUser, dbPass, dbName)
	}

	database, err := sql.Open("mysql", dbURI)
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: database, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	dbmap.AddTableWithName(entity.Tag{}, "tags").SetKeys(true, "Id")
	dbmap.AddTableWithName(entity.Image{}, "images").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
