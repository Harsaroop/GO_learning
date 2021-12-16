package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "azr-ist-0120-wrk-cc-dri-0000-sql-mu08s.database.windows.net"
var protocol = "tcp"
var port = 1433
var user = "itsazadmin"
var password = "vJh7#LkTrbZ!3gP6mmzx&Z7U5%2mN#WJgkrzJYJi59kHmw"
var database = "azr-ist-0120-wrk-cc-dri-cpslrcdbs01"

func main() {
	//Connection string
	connString := fmt.Sprintf("server=%s;user=%s;password=%s;port=%d;database=%s;protocol=%s",
		server, user, password, port, database, protocol)

	var err error

	// connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating the connection pool: ", err.Error())

	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	fmt.Printf("Connected\n")

	// Running a query
	query := fmt.Sprintf("SELECT * FROM collective_agreement")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Print("No data")
	}

	defer rows.Close()

	fmt.Print(reflect.TypeOf(rows))

	// printing data into tabular form

	//for rows.Next() {
	//if err :=
	//}

	//router_handle := gin.Default()

	//router_handle.GET("/ping", PingGet)

	//router_handle.Run() // it will listen and serve on 0.0.0.0:8080
}
